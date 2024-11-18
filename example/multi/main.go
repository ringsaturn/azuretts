package main

import (
	"context"
	"os"

	"github.com/caiyunapp/azuretts"
)

type MultiSpeak struct {
	XML string
}

func (ms *MultiSpeak) ToXML() ([]byte, error) {
	return []byte(ms.XML), nil
}

var mess = `<speak version="1.0" xmlns="http://www.w3.org/2001/10/synthesis" xmlns:mstts="https://www.w3.org/2001/mstts" xml:lang="zh-CN">
<voice name="zh-CN-XiaomoNeural">
	女儿看见父亲走了进来，问道：
	<mstts:express-as role="YoungAdultFemale" style="calm">
		“您来的挺快的，怎么过来的？”
	</mstts:express-as>
	父亲放下手提包，说：
	<mstts:express-as role="OlderAdultMale" style="calm">
		“刚打车过来的，路上还挺顺畅。”
	</mstts:express-as>
</voice>
</speak>`

func main() {
	c := azuretts.NewClient(
		os.Getenv("SPEECH_KEY"),
		azuretts.Region(os.Getenv("SPEECH_REGION")),
	)

	b, err := c.GetSynthesize(context.Background(), &azuretts.SynthesisRequest{
		Speak:  &MultiSpeak{XML: mess},
		Output: azuretts.AudioOutputFormat_Streaming_Audio16Khz32KbitrateMonoMp3,
	})
	if err != nil {
		panic(err)
	}
	if err := b.Error(); err != nil {
		panic(err)
	}
	err = os.WriteFile("multi.mp3", b.Body, 0644)
	if err != nil {
		panic(err)
	}
}
