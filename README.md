# Azure TTS (Text to Speech) API wrapper for Go.

## Quick Start

```bash
go get github.com/ringsaturn/azuretts
```

```go
package main

import (
	"context"
	"os"

	"github.com/ringsaturn/azuretts"
)

func main() {
	c := azuretts.NewClient(
		os.Getenv("SPEECH_KEY"),
		azuretts.Region(os.Getenv("SPEECH_REGION")),
	)
	speak := azuretts.NewSpeak(
		azuretts.WithLanguage(azuretts.LanguageZhCN),
		azuretts.WithVoiceName(azuretts.VoiceNameZhCNYunxiNeural),
		azuretts.WithStyle(azuretts.StyleChat),
		azuretts.WithRate(1),
		azuretts.WithVoiceStyledegree(2),
		azuretts.WithSpeechText("你好，世界"),
		azuretts.WithVolume(100),
	)
	b, err := c.GetSynthesize(context.Background(), &azuretts.SynthesisRequest{
		Speak:  speak,
		Output: azuretts.AudioOutputFormat_Streaming_Audio16Khz32KbitrateMonoMp3,
	})
	if err != nil {
		panic(err)
	}
	if err := b.Error(); err != nil {
		panic(err)
	}
	err = os.WriteFile("sample.mp3", b.Body, 0644)
	if err != nil {
		panic(err)
	}
}
```

## Thanks

I learned a lot from <https://github.com/jesseward/azuretexttospeech>.
