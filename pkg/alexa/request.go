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
	Slots              map[string]*Slot   `json:"slots"`
	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus"`
}

// Slot is an Alexa skill slot.
type Slot struct {
	Name        string       `json:"name"`
	Value       string       `json:"value"`
	Resolutions *Resolutions `json:"resolutions"`
	Source      string       `json:"source"`
	SlotValue   *SlotValue   `json:"slotValue"`
}

// SlotValue defines the value or values captured by the slot
type SlotValue struct {
	Type        string       `json:"type"`
	Value       string       `json:"value"`
	Resolutions *Resolutions `json:"resolutions"`
}

type AuthorityValueValue struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type AuthorityValue struct {
	Value *AuthorityValueValue `json:"value,omitempty"`
}

// ResolutionStatus represents the status code of a slot resolution
type StatusCode string

const (
	// ResolutionStatusMatch is the status code for match
	ResolutionStatusMatch StatusCode = "ER_SUCCESS_MATCH"
	// ResolutionStatusNoMatch is the status code for no match
	ResolutionStatusNoMatch StatusCode = "ER_SUCCESS_NO_MATCH"
	// ResolutionStatusTimeout is the status code for an error due to timeout
	ResolutionStatusTimeout StatusCode = "ER_ERROR_TIMEOUT"
	// ResolutionStatusException is the status code for an error in processing
	ResolutionStatusException StatusCode = "ER_ERROR_EXCEPTION"
)

type ResolutionStatus struct {
	Code StatusCode `json:"code"`
}

type PerAuthority struct {
	Authority string            `json:"authority"`
	Status    *ResolutionStatus `json:"status,omitempty"`
	Values    []*AuthorityValue `json:"values,omitempty"`
}

// Resolutions is an Alexa skill resolution.
type Resolutions struct {
	PerAuthority []*PerAuthority `json:"resolutionsPerAuthority"`
}

// UpdatedIntent is to update the Intent.
// **same** as Intent, just with a different json key
//type UpdatedIntent struct {
//	Name               string             `json:"name,omitempty"`
//	ConfirmationStatus ConfirmationStatus `json:"confirmationStatus,omitempty"`
//	Slots              map[string]Slot    `json:"slots,omitempty"`
//}

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

// DialogStateType represents JSON request `request.dialogState`, see https://developer.amazon.com/docs/custom-skills/delegate-dialog-to-alexa.html
type DialogStateType string

const (
	DialogStateStarted    DialogStateType = "STARTED"
	DialogStateInProgress DialogStateType = "IN_PROGRESS"
	DialogStateCompleted  DialogStateType = "COMPLETED"
)

// Request represents the information about the request.
type Request struct {
	Type        RequestType     `json:"type"`
	RequestID   string          `json:"requestId"`
	Timestamp   string          `json:"timestamp"`
	Locale      string          `json:"locale"`
	Intent      Intent          `json:"intent,omitempty"`
	Reason      string          `json:"reason,omitempty"`
	DialogState DialogStateType `json:"dialogState,omitempty"`

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
