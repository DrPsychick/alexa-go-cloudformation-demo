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

// Intent is the Alexa skill intent.
type Intent struct {
	Name               string             `json:"name"`
	Slots              map[string]Slot    `json:"slots"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
}

// Slot is an Alexa skill slot
type Slot struct {
	Name        string       `json:"name"`
	Value       string       `json:"value"`
	Resolutions *Resolutions `json:"resolutions"`
}

// Resolutions is an Alexa skill resolution.
type Resolutions struct {
	PerAuthority []*struct {
		Authority string `json:"authority"`
		Status    struct {
			Code string `json:"code"`
		} `json:"status"`
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}

// UpdatedIntent is to update the Intent
type UpdatedIntent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus ConfirmationStatus     `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}
