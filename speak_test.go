package azuretts_test

import (
	"testing"

	"github.com/caiyunapp/azuretts"
)

func BenchmarkToXML(b *testing.B) {
	speak := azuretts.NewSpeak(
		azuretts.WithLanguage(azuretts.LanguageZhCN),
		azuretts.WithVoiceName(azuretts.VoiceNameZhCNYunxiNeural),
		azuretts.WithStyle(azuretts.StyleChat),
		azuretts.WithRate(1),
		azuretts.WithVoiceStyledegree(2),
		azuretts.WithSpeechText("你好，世界"),
		azuretts.WithVolume(100),
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := speak.ToXML()
		if err != nil {
			b.Fatal(err)
		}
	}
}
