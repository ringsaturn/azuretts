package azuretts

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	textToSpeechAPI = "https://%s.tts.speech.microsoft.com/cognitiveservices/v1"
	tokenRefreshAPI = "https://%s.api.cognitive.microsoft.com/sts/v1.0/issueToken"
	userAgent       = "github.com/ringsaturn/azuretts@v0"

	// https://learn.microsoft.com/en-us/azure/api-management/api-management-subscriptions
	subscriptionKeyHeader = "Ocp-Apim-Subscription-Key"

	outputFormatHeader = "X-Microsoft-OutputFormat"
)

var (
	accessTokenTTL = 8 * time.Minute // normaly 10 minutes
)

// AccessTokenSaver is an interface to save and get access token.
//
// Please implement this interface if you want to save access token in Redis or other places.
type AccessTokenSaver interface {
	GetAccessToken(context.Context) (string, int64, error)
	SetAccessToken(context.Context, string, int64) error
}

type memorySaver struct {
	mapAccessToken     string
	mapAccessTokenExp  int64
	mapAccessTokenLock sync.RWMutex
}

func (s *memorySaver) GetAccessToken(ctx context.Context) (string, int64, error) {
	s.mapAccessTokenLock.Lock()
	defer s.mapAccessTokenLock.Unlock()
	return s.mapAccessToken, s.mapAccessTokenExp, nil
}

func (s *memorySaver) SetAccessToken(ctx context.Context, accessToken string, exp int64) error {
	s.mapAccessTokenLock.RLock()
	s.mapAccessToken = accessToken
	s.mapAccessTokenExp = exp
	s.mapAccessTokenLock.RUnlock()
	return nil
}

// AutoRefresh is a function to refresh access token based on token expire time.
type AutoRefresh func(ctx context.Context, client Client) (string, int64, error)

func newAutoRefresh() AutoRefresh {
	mutex := &sync.Mutex{}
	return func(ctx context.Context, client Client) (string, int64, error) {
		token, exp, err := client.GetAccessToken(ctx)
		if err != nil {
			return "", 0, err
		}
		// If lock failed, it means another goroutine is refreshing token.
		// So we just return current token.
		locked := mutex.TryLock()
		if !locked {
			return token, exp, nil
		}
		defer mutex.Unlock()
		now := time.Now().Unix()
		if exp-now > 60 {
			return token, exp, nil
		}
		resp, err := client.GetNewAccessToken(ctx)
		if err != nil {
			return "", 0, err
		}
		err = client.SetAccessToken(ctx, resp.AccessToken, resp.ExpiresInSeconds)
		if err != nil {
			return "", 0, err
		}
		return resp.AccessToken, resp.ExpiresInSeconds, nil
	}
}

type AccessTokenResponse struct {
	AccessToken      string `json:"token"`
	ExpiresInSeconds int64  `json:"exp"`
}

type Client interface {
	AccessTokenSaver

	GetNewAccessToken(context.Context) (*AccessTokenResponse, error)
	GetSynthesize(ctx context.Context, req *SynthesisRequest) (*SynthesisResponse, error)
}

type Option func(*baseClient)

// Will use `memorySaver`(an internal function) default.
//
// Please implement your own `AccessTokenSaver` if you want to save access token
// in Redis or other places.
func WithTokenSaver(saver AccessTokenSaver) Option {
	return func(c *baseClient) {
		c.tokenSaver = saver
	}
}

// Will use `defaultAutoRefresh`(an internal function) by default.
// If you want to disable auto refresh, please set this option to nil.
//
// If you want to implement your own auto refresh function, please make sure the
// function is thread safe. Because the function could be called by multiple
// goroutines.
func WithAutoTokenRefresh(fn AutoRefresh) Option {
	return func(c *baseClient) {
		c.autoRefreshFn = fn
	}
}

// Will use [http.DefaultClient] by default.
func WithHTTPClient(client *http.Client) Option {
	return func(c *baseClient) {
		c.client = client
	}
}

type baseClient struct {
	subscriptionKey string
	region          Region

	authURL string
	ttsURL  string

	tokenSaver    AccessTokenSaver
	client        *http.Client
	autoRefreshFn AutoRefresh
}

func NewClient(subscriptionKey string, region Region, opts ...Option) Client {
	c := &baseClient{
		subscriptionKey: subscriptionKey,
		region:          region,
		tokenSaver:      &memorySaver{},
		client:          http.DefaultClient,
		autoRefreshFn:   newAutoRefresh(),
	}
	c.authURL = fmt.Sprintf(tokenRefreshAPI, c.region.String())
	c.ttsURL = fmt.Sprintf(textToSpeechAPI, c.region.String())
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *baseClient) GetAccessToken(ctx context.Context) (string, int64, error) {
	return c.tokenSaver.GetAccessToken(ctx)
}

func (c *baseClient) SetAccessToken(ctx context.Context, accessToken string, exp int64) error {
	return c.tokenSaver.SetAccessToken(ctx, accessToken, exp)
}

func (c *baseClient) GetNewAccessToken(ctx context.Context) (*AccessTokenResponse, error) {
	request, _ := http.NewRequestWithContext(ctx, http.MethodPost, c.authURL, nil)
	request.Header.Set(subscriptionKeyHeader, c.subscriptionKey)

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get access token, status code: %d", response.StatusCode)
	}

	body, _ := io.ReadAll(response.Body)
	accessToken := string(body)
	now := time.Now()
	return &AccessTokenResponse{
		AccessToken:      accessToken,
		ExpiresInSeconds: now.Add(accessTokenTTL).Unix(),
	}, nil
}

type SynthesisRequest struct {
	Speak  Payload
	Output AudioOutputFormat
}

type SynthesisResponse struct {
	Status     int
	Body       []byte
	Resp       *http.Response
	RequestXML []byte
}

func (resp *SynthesisResponse) Error() error {
	if resp.Status != http.StatusOK {
		return fmt.Errorf("failed to synthesis, status code: %d", resp.Status)
	}
	return nil
}

func (c *baseClient) GetSynthesize(ctx context.Context, req *SynthesisRequest) (*SynthesisResponse, error) {
	var (
		accessToken string
		err         error
	)
	if c.autoRefreshFn != nil {
		accessToken, _, err = c.autoRefreshFn(ctx, c)
	} else {
		accessToken, _, err = c.GetAccessToken(ctx)
	}
	if err != nil {
		return nil, err
	}

	xmlBytes, err := req.Speak.ToXML()
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(
		ctx, http.MethodPost,
		c.ttsURL,
		bytes.NewBuffer(xmlBytes),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set(outputFormatHeader, req.Output.String())
	request.Header.Set("Content-Type", "application/ssml+xml")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("User-Agent", userAgent)

	response, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &SynthesisResponse{
		Status:     response.StatusCode,
		Body:       body,
		Resp:       response,
		RequestXML: xmlBytes,
	}, nil
}
