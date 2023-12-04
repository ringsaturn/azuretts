package azuretts

// https://github.com/MicrosoftDocs/azure-docs/blob/main/articles/ai-services/speech-service/includes/language-support/voice-styles-and-roles.md?plain=1
type Style string

const (
	StyleCheerful                Style = "cheerful"
	StyleChat                    Style = "chat"
	StyleSad                     Style = "sad"
	StyleAngry                   Style = "angry"
	StyleCustomerservice         Style = "customerservice"
	StyleEmpathetic              Style = "empathetic"
	StyleExcited                 Style = "excited"
	StyleFriendly                Style = "friendly"
	StyleHopeful                 Style = "hopeful"
	StyleNarrationprofessional   Style = "narration-professional"
	StyleNewscastcasual          Style = "newscast-casual"
	StyleNewscastformal          Style = "newscast-formal"
	StyleShouting                Style = "shouting"
	StyleTerrified               Style = "terrified"
	StyleUnfriendly              Style = "unfriendly"
	StyleWhispering              Style = "whispering"
	StyleNewscast                Style = "newscast"
	StyleAssistant               Style = "assistant"
	StyleCalm                    Style = "calm"
	StyleAffectionate            Style = "affectionate"
	StyleDisgruntled             Style = "disgruntled"
	StyleEmbarrassed             Style = "embarrassed"
	StyleFearful                 Style = "fearful"
	StyleGentle                  Style = "gentle"
	StyleSerious                 Style = "serious"
	StyleDepressed               Style = "depressed"
	StyleEnvious                 Style = "envious"
	StyleLyrical                 Style = "lyrical"
	StylePoetryreading           Style = "poetry-reading"
	StyleAdvertisementupbeat     Style = "advertisement-upbeat"
	StyleDocumentarynarration    Style = "documentary-narration"
	StyleNarrationrelaxed        Style = "narration-relaxed"
	StyleSportscommentary        Style = "sports-commentary"
	StyleSportscommentaryexcited Style = "sports-commentary-excited"
)

func (s Style) String() string {
	return string(s)
}
