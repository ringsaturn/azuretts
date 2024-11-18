package main

import (
	"context"
	"flag"
	"os"

	"github.com/caiyunapp/azuretts"
)

var (
	languageFlag    = flag.String("language", "zh-CN", "Language Flags")
	voiceFlag       = flag.String("voice", azuretts.VoiceNameZhCNYunxiNeural.String(), "Voice Flags")
	styleFlag       = flag.String("style", azuretts.StyleChat.String(), "Style Flags")
	rateFlag        = flag.Int("rate", 1, "Rate")
	styleDegreeFlag = flag.Int("styledegree", 2, "Style Degree")
	volumeFlag      = flag.Int("volume", 100, "Volume")
	textFlag        = flag.String("text", "你好，世界", "Text")
	outputFormat    = flag.String("format", azuretts.AudioOutputFormat_Streaming_Audio16Khz32KbitrateMonoMp3.String(), "Output Format")
	outputFileName  = flag.String("output", "audio.mp3", "Output File Name")
)

func main() {
	flag.Parse()

	c := azuretts.NewClient(
		os.Getenv("SPEECH_KEY"),
		azuretts.Region(os.Getenv("SPEECH_REGION")),
	)
	speak := azuretts.NewSpeak(
		azuretts.WithLanguage(azuretts.Language(*languageFlag)),
		azuretts.WithVoiceName(azuretts.VoiceName(*voiceFlag)),
		azuretts.WithStyle(azuretts.Style(*styleFlag)),
		azuretts.WithRate(float64(*rateFlag)),
		azuretts.WithVoiceStyledegree(float64(*styleDegreeFlag)),
		azuretts.WithSpeechText(*textFlag),
		azuretts.WithVolume(*volumeFlag),
	)
	b, err := c.GetSynthesize(context.Background(), &azuretts.SynthesisRequest{
		Speak:  speak,
		Output: azuretts.AudioOutputFormat(*outputFormat),
	})
	if err != nil {
		panic(err)
	}
	if err := b.Error(); err != nil {
		panic(err)
	}
	err = os.WriteFile(*outputFileName, b.Body, 0644)
	if err != nil {
		panic(err)
	}
}
