package alexa

// built in intents
const (
	//HelpIntent is the Alexa built-in Help Intent
	HelpIntent = "AMAZON.HelpIntent"

	//CancelIntent is the Alexa built-in Cancel Intent
	CancelIntent = "AMAZON.CancelIntent"

	//StopIntent is the Alexa built-in Stop Intent
	StopIntent = "AMAZON.StopIntent"
)

// Intent is the Alexa skill intent
type Intent struct {
	Name               string             `json:"name"`
	Slots              map[string]Slot    `json:"slots"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
}
