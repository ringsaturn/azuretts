package azuretts

// https://learn.microsoft.com/en-us/azure/ai-services/speech-service/speech-synthesis-markup-voice
type Role string

const (
	RoleGirl             Role = "Girl"             // The voice imitates a girl.
	RoleBoy              Role = "Boy"              // The voice imitates a boy.
	RoleYoungAdultFemale Role = "YoungAdultFemale" // The voice imitates a young adult female.
	RoleYoungAdultMale   Role = "YoungAdultMale"   // The voice imitates a young adult male.
	RoleOlderAdultFemale Role = "OlderAdultFemale" // The voice imitates an older adult female.
	RoleOlderAdultMale   Role = "OlderAdultMale"   // The voice imitates an older adult male.
	RoleSeniorFemale     Role = "SeniorFemale"     // The voice imitates a senior female.
	RoleSeniorMale       Role = "SeniorMale"       // The voice imitates a senior male.
)

func (r Role) String() string {
	return string(r)
}
