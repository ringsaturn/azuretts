# Azure TTS (Text to Speech) API wrapper for Go.

## Quick Start

```bash
go get github.com/caiyunapp/azuretts
```

```go
package main

import (
	"context"
	"os"

	"github.com/caiyunapp/azuretts"
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

## CLI

```bash
go install github.com/caiyunapp/azuretts/cmd/azuretts
```

NOTE: Please setup environment variables `SPEECH_KEY` and `SPEECH_REGION` before
running the CLI.

```console
azuretts --help
Usage of ./azuretts:
  -language string
        Language Flags (default "zh-CN")
  -output string
        Output File Name (default "audio.mp3")
  -rate int
        Rate (default 1)
  -style string
        Style Flags (default "chat")
  -styledegree int
        Style Degree (default 2)
  -text string
        Text (default "你好，世界")
  -voice string
        Voice Flags (default "zh-CN-YunxiNeural")
  -volume int
        Volume (default 100)
```

## Thanks

I learned a lot from <https://github.com/jesseward/azuretexttospeech>.
