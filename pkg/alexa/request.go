package alexa

// Locale constants.
const (
	// LocaleAmericanEnglish is the locale for American English.
	LocaleAmericanEnglish = "en-US"

	// LocaleAustralianEnglish is the locale for Australian English.
	LocaleAustralianEnglish = "en-AU"

	// LocaleBritishEnglish is the locale for UK English.
	LocaleBritishEnglish = "en-GB"

	// LocaleCanadianEnglish is the locale for Canadian English.
	LocaleCanadianEnglish = "en-CA"

	// LocaleCanadianFrench is the locale for Canadian French.
	LocaleCanadianFrench = "fr-CA"

	// LocaleFrenchFrench is the locale for French (France).
	LocaleFrench = "fr-FR"

	// LocaleGerman is the locale for standard dialect German (Germany).
	LocaleGerman = "de-DE"

	//LocaleIndianEnglish is the locale for Indian English.
	LocaleIndianEnglish = "en-IN"

	// LocaleItalian is the locale for Italian (Italy).
	LocaleItalian = "it-IT"

	// LocaleJapanese is the locale for Japanese (Japan).
	LocaleJapanese = "ja-JP"

	// LocaleMexicanSpanish is the locale for Mexican Spanish.
	LocaleMexicanSpanish = "es-MX"

	// LocaleSpanish is the  for Spanish (Spain).
	LocaleSpanish = "es-ES"
)

// ConfirmationStatus represents confirmationStatus in JSON
type ConfirmationStatus string

const (
	// ConfirmationStatusNone is constant `NONE`.
	ConfirmationStatusNone = "NONE"

	// ConfirmationStatusConfirmed is constant `CONFIRMED`.
	ConfirmationStatusConfirmed = "CONFIRMED"

	// ConfirmationStatusDenied is constant `DENIED`.
	ConfirmationStatusDenied = "DENIED"
)

// Built in intents.
const (
	//HelpIntent is the Alexa built-in Help Intent.
	HelpIntent = "AMAZON.HelpIntent"

	//CancelIntent is the Alexa built-in Cancel Intent.
	CancelIntent = "AMAZON.CancelIntent"

	//StopIntent is the Alexa built-in Stop Intent.
	StopIntent = "AMAZON.StopIntent"
)

// Intent is the Alexa skill intent.
type Intent struct {
	Name               string             `json:"name"`
	Slots              map[string]Slot    `json:"slots"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
}

// Slot is an Alexa skill slot.
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

// UpdatedIntent is to update the Intent.
type UpdatedIntent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus ConfirmationStatus     `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}

// RequestType represents JSON request `request.type`, see https://developer.amazon.com/docs/custom-skills/request-types-reference.html
type RequestType string

// Request type constants.
const (
	// TypeLaunchRequest
	TypeLaunchRequest RequestType = "LaunchRequest"
	// TypeIntentRequest
	TypeIntentRequest RequestType = "IntentRequest"
	// TypeSessionEndedRequest
	TypeSessionEndedRequest RequestType = "SessionEndedRequest"
	// TypeCanFulfillIntentRequest
	TypeCanFulfillIntentRequest RequestType = "CanFulfillIntentRequest"
)

// Request represents the information about the request.
type Request struct {
	Type        RequestType `json:"type"`
	RequestID   string      `json:"requestId"`
	Timestamp   string      `json:"timestamp"`
	Locale      string      `json:"locale"`
	Intent      Intent      `json:"intent,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	DialogState string      `json:"dialogState,omitempty"`

	Context *Context `json:"-"`
	Session *Session `json:"-"`
}

// Session represents the Alexa skill session.
type Session struct {
	New         bool   `json:"new"`
	SessionID   string `json:"sessionId"`
	Application struct {
		ApplicationID string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		UserID      string `json:"userId"`
		AccessToken string `json:"accessToken,omitempty"`
	} `json:"user"`
}

// Context represents the Alexa skill request context.
type Context struct {
	System struct {
		APIAccessToken string `json:"apiAccessToken"`
		Device         struct {
			DeviceID string `json:"deviceId,omitempty"`
		} `json:"device,omitempty"`
		Application struct {
			ApplicationID string `json:"applicationId,omitempty"`
		} `json:"application,omitempty"`
	} `json:"System,omitempty"`
}

// RequestEnvelope represents the alexa request envelope.
type RequestEnvelope struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Context *Context `json:"context"`
	Request *Request `json:"request"`
}
