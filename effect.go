package azuretts

// The audio effect processor that's used to optimize the quality of the
// synthesized speech output for specific scenarios on devices.
//
// For some scenarios in production environments, the auditory experience might
// be degraded due to the playback distortion on certain devices.
// For example, the synthesized speech from a car speaker might sound dull and
// muffled due to environmental factors such as speaker response, room
// reverberation, and background noise. The passenger might have to turn up the
// volume to hear more clearly. To avoid manual operations in such a scenario,
// the audio effect processor can make the sound clearer by compensating the
// distortion of playback.
//
// Source: https://learn.microsoft.com/en-us/azure/ai-services/speech-service/speech-synthesis-markup-voice
type Effect string

const (
	// Optimize the auditory experience when providing high-fidelity speech in
	// cars, buses, and other enclosed automobiles.
	EffectEqCar Effect = "eq_car"

	// Optimize the auditory experience for narrowband speech in telecom or
	// telephone scenarios. You should use a sampling rate of 8 kHz. If the
	// sample rate isn't 8 kHz, the auditory quality of the output speech isn't
	// optimized.
	EffectEqTelecomhp8k Effect = "eq_telecomhp8k"
)
