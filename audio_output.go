package azuretts

// The supported streaming and non-streaming audio formats are sent in each
// request as the `X-Microsoft-OutputFormat` header. Each format incorporates a
// bit rate and encoding type. The Speech service supports 48-kHz, 24-kHz,
// 16-kHz, and 8-kHz audio outputs. Each prebuilt neural voice model is
// available at 24kHz and high-fidelity 48kHz.
//
// If you select 48kHz output format, the high-fidelity voice model with 48kHz
// will be invoked accordingly. The sample rates other than 24kHz and 48kHz can
// be obtained through upsampling or downsampling when synthesizing, for
// example, 44.1kHz is downsampled from 48kHz.
// If your selected voice and output format have different bit rates, the audio
// is resampled as necessary. You can decode the ogg-24khz-16bit-mono-opus
// format by using the [Opus codec](https://opus-codec.org/downloads/).
//
// Source: https://learn.microsoft.com/en-us/azure/ai-services/speech-service/rest-text-to-speech?tabs=nonstreaming#audio-outputs
type AudioOutputFormat string

const (
	AudioOutputFormat_Streaming_AmrWb16000Hz                  AudioOutputFormat = "amr-wb-16000hz"
	AudioOutputFormat_Streaming_Audio16Khz16Bit32KbpsMonoOpus AudioOutputFormat = "audio-16khz-16bit-32kbps-mono-opus"
	AudioOutputFormat_Streaming_Audio16Khz32KbitrateMonoMp3   AudioOutputFormat = "audio-16khz-32kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio16Khz64KbitrateMonoMp3   AudioOutputFormat = "audio-16khz-64kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio16Khz128KbitrateMonoMp3  AudioOutputFormat = "audio-16khz-128kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio24Khz16Bit24KbpsMonoOpus AudioOutputFormat = "audio-24khz-16bit-24kbps-mono-opus"
	AudioOutputFormat_Streaming_Audio24Khz16Bit48KbpsMonoOpus AudioOutputFormat = "audio-24khz-16bit-48kbps-mono-opus"
	AudioOutputFormat_Streaming_Audio24Khz48KbitrateMonoMp3   AudioOutputFormat = "audio-24khz-48kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio24Khz96KbitrateMonoMp3   AudioOutputFormat = "audio-24khz-96kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio24Khz160KbitrateMonoMp3  AudioOutputFormat = "audio-24khz-160kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio48Khz96KbitrateMonoMp3   AudioOutputFormat = "audio-48khz-96kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Audio48Khz192KbitrateMonoMp3  AudioOutputFormat = "audio-48khz-192kbitrate-mono-mp3"
	AudioOutputFormat_Streaming_Ogg16Khz16BitMonoOpus         AudioOutputFormat = "ogg-16khz-16bit-mono-opus"
	AudioOutputFormat_Streaming_Ogg24Khz16BitMonoOpus         AudioOutputFormat = "ogg-24khz-16bit-mono-opus"
	AudioOutputFormat_Streaming_Ogg48Khz16BitMonoOpus         AudioOutputFormat = "ogg-48khz-16bit-mono-opus"
	AudioOutputFormat_Streaming_Raw8Khz8BitMonoAlaw           AudioOutputFormat = "raw-8khz-8bit-mono-alaw"
	AudioOutputFormat_Streaming_Raw8Khz8BitMonoMulaw          AudioOutputFormat = "raw-8khz-8bit-mono-mulaw"
	AudioOutputFormat_Streaming_Raw8Khz16BitMonoPcm           AudioOutputFormat = "raw-8khz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Raw16Khz16BitMonoPcm          AudioOutputFormat = "raw-16khz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Raw16Khz16BitMonoTruesilk     AudioOutputFormat = "raw-16khz-16bit-mono-truesilk"
	AudioOutputFormat_Streaming_Raw22050Hz16BitMonoPcm        AudioOutputFormat = "raw-22050hz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Raw24Khz16BitMonoPcm          AudioOutputFormat = "raw-24khz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Raw24Khz16BitMonoTruesilk     AudioOutputFormat = "raw-24khz-16bit-mono-truesilk"
	AudioOutputFormat_Streaming_Raw44100Hz16BitMonoPcm        AudioOutputFormat = "raw-44100hz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Raw48Khz16BitMonoPcm          AudioOutputFormat = "raw-48khz-16bit-mono-pcm"
	AudioOutputFormat_Streaming_Webm16Khz16BitMonoOpus        AudioOutputFormat = "webm-16khz-16bit-mono-opus"
	AudioOutputFormat_Streaming_Webm24Khz16Bit24KbpsMonoOpus  AudioOutputFormat = "webm-24khz-16bit-24kbps-mono-opus"
	AudioOutputFormat_Streaming_Webm24Khz16BitMonoOpus        AudioOutputFormat = "webm-24khz-16bit-mono-opus"

	AudioOutputFormat_NonStreaming_Riff8Khz8BitMonoAlaw    AudioOutputFormat = "riff-8khz-8bit-mono-alaw"
	AudioOutputFormat_NonStreaming_Riff8Khz8BitMonoMulaw   AudioOutputFormat = "riff-8khz-8bit-mono-mulaw"
	AudioOutputFormat_NonStreaming_Riff8Khz16BitMonoPcm    AudioOutputFormat = "riff-8khz-16bit-mono-pcm"
	AudioOutputFormat_NonStreaming_Riff22050Hz16BitMonoPcm AudioOutputFormat = "riff-22050hz-16bit-mono-pcm"
	AudioOutputFormat_NonStreaming_Riff24Khz16BitMonoPcm   AudioOutputFormat = "riff-24khz-16bit-mono-pcm"
	AudioOutputFormat_NonStreaming_Riff44100Hz16BitMonoPcm AudioOutputFormat = "riff-44100hz-16bit-mono-pcm"
	AudioOutputFormat_NonStreaming_Riff48Khz16BitMonoPcm   AudioOutputFormat = "riff-48khz-16bit-mono-pcm"
)

func (a AudioOutputFormat) String() string {
	return string(a)
}
