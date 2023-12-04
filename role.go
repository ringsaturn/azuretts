package azuretts

// https://learn.microsoft.com/en-us/azure/ai-services/speech-service/speech-synthesis-markup-voice
type Role string

const (
	RoleGirl             Role = "Girl"
	RoleBoy              Role = "Boy"
	RoleYoungAdultFemale Role = "YoungAdultFemale"
	RoleYoungAdultMale   Role = "YoungAdultMale"
	RoleOlderAdultFemale Role = "OlderAdultFemale"
	RoleOlderAdultMale   Role = "OlderAdultMale"
	RoleSeniorFemale     Role = "SeniorFemale"
	RoleSeniorMale       Role = "SeniorMale"
)

func (r Role) String() string {
	return string(r)
}
