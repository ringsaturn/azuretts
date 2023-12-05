package azuretts

import (
	"encoding/xml"
	"fmt"
)

type Payload interface {
	ToXML() ([]byte, error)
}

type Prosody struct {
	SpeechText string `xml:",chardata"`
	Rate       string `xml:"rate,attr"`
	Volume     string `xml:"volume,attr"`
}

type ExpressAs struct {
	// Text        string  `xml:",chardata"`
	Style       Style   `xml:"style,attr"`
	Styledegree string  `xml:"styledegree,attr"`
	Prosody     Prosody `xml:"prosody"`
	Role        Role    `xml:"role,attr,omitempty"`
}

type Voice struct {
	// Text      string    `xml:",chardata"`
	Name      VoiceName `xml:"name,attr"`
	ExpressAs ExpressAs `xml:"mstts:express-as"`
	Effect    Effect    `xml:"effect,omitempty"`
}

// Speak is the root element of the SSML document.
//
// https://learn.microsoft.com/en-us/azure/ai-services/speech-service/speech-synthesis-markup-structure
type Speak struct {
	XMLName xml.Name `xml:"speak"`
	// Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Xmlns   string   `xml:"xmlns,attr"`
	Mstts   string   `xml:"xmlns:mstts,attr"`
	Lang    Language `xml:"xml:lang,attr"`
	Voice   Voice    `xml:"voice"`
}

func (s *Speak) ToXML() ([]byte, error) {
	return xml.MarshalIndent(s, "", "  ")
}

type SpeakOption func(*Speak)

func WithStyle(style Style) SpeakOption {
	return func(s *Speak) {
		s.Voice.ExpressAs.Style = style
	}
}

func WithVoiceName(name VoiceName) SpeakOption {
	return func(s *Speak) {
		s.Voice.Name = name
	}
}

func WithSpeechText(text string) SpeakOption {
	return func(s *Speak) {
		s.Voice.ExpressAs.Prosody.SpeechText = text
	}
}

func WithLanguage(lang Language) SpeakOption {
	return func(s *Speak) {
		s.Lang = lang
	}
}

func WithRate(rate float64) SpeakOption {
	return func(s *Speak) {
		s.Voice.ExpressAs.Prosody.Rate = fmt.Sprintf("%.2f", rate)
	}
}

func WithVoiceStyledegree(degree float64) SpeakOption {
	return func(s *Speak) {
		s.Voice.ExpressAs.Styledegree = fmt.Sprintf("%.2f", degree)
	}
}

func WithVolume(volume int) SpeakOption {
	return func(s *Speak) {
		s.Voice.ExpressAs.Prosody.Volume = fmt.Sprintf("%d", volume)
	}
}

func NewDefaultSpeak() *Speak {
	return &Speak{
		Version: "1.0",
		Xmlns:   "https://www.w3.org/2001/10/synthesis",
		Mstts:   "https://www.w3.org/2001/mstts",
		Voice: Voice{
			Name: VoiceNameEnUSAIGenerate1Neural,
			ExpressAs: ExpressAs{
				Style:       "chat",
				Styledegree: "1",
				Prosody: Prosody{
					Rate:       "1.0",
					Volume:     "100",
					SpeechText: "",
				},
			},
		},
	}
}

func NewSpeak(opts ...SpeakOption) *Speak {
	s := NewDefaultSpeak()
	for _, opt := range opts {
		opt(s)
	}
	return s
}
