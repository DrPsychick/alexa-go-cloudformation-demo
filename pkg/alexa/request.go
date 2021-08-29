package alexa

import "errors"

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

// Intent returns the intent if it exists
func (r *RequestEnvelope) Intent() (Intent, error) {
	i := r.Request.Intent
	if i.Name == "" {
		return Intent{}, errors.New("request has no intent")
	}

	return i, nil
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

func (r *RequestEnvelope) Slot(name string) (Slot, error) {
	i, err := r.Intent()
	if err != nil {
		return Slot{}, err
	}

	s, ok := i.Slots[name]
	if !ok {
		return Slot{}, errors.New("slot does not exist")
	}

	return *s, nil
}

// SlotValue returns the value of the slot if it exists
func (r *RequestEnvelope) SlotValue(name string) (string, error) {
	s, err := r.Slot(name)
	if err != nil {
		return "", err
	}

	return s.Value, nil
}

func (s *Slot) SlotResolutionsPerAuthorities() ([]*PerAuthority, error) {
	if s.Resolutions == nil {
		return []*PerAuthority{}, errors.New("slot has no resolutions")
	}

	return s.Resolutions.PerAuthority, nil
}

// SlotAuthorities returns
func (s *Slot) FirstAuthorityWithMatch(name string) (PerAuthority, error) {
	auths, err := s.SlotResolutionsPerAuthorities()
	if err != nil {
		return PerAuthority{}, err
	}

	for _, a := range auths {
		if a.Status != nil && a.Status.Code == ResolutionStatusMatch {
			return *a, nil
		}
	}

	return PerAuthority{}, errors.New("no resolution with match")
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

// ResolutionStatus indicates the results of attempting to resolve the user utterance against the defined slot types
type ResolutionStatus struct {
	Code StatusCode `json:"code"`
}

// PerAuthority encapsualtes an Authority which is the source of the data provided
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

// ContextUser a string that represents a unique identifier for the Amazon account for which the skill is enabled
type ContextUser struct {
	UserID      string `json:"userId"`
	AccessToken string `json:"accessToken,omitempty"`
}

// ContextApplication is used to verify that the request was intended for your service, the ID is the appliation ID for your skill
type ContextApplication struct {
	ApplicationID string `json:"applicationId"`
}

// Session represents the Alexa skill session.
type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application ContextApplication     `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        ContextUser            `json:"user"`
}

// ContextSystemPerson describes the person who is making the request to Alexa (user recognized by voice, not account)
type ContextSystemPerson struct {
	PersonID    string `json:"personId"`
	AccessToken string `json:"accessToken,omitempty"`
}

// ContextSystem provides information about the current state of the Alexa service and the device interacting with your skill
type ContextSystem struct {
	// APIAccessToken a string containing a token that can be used to access Alexa-specific APIs
	APIAccessToken string `json:"apiAccessToken,omitempty"`
	// APIEndpoint a string that references the correct base URI to refer to by region, for use with APIs
	APIEndpoint string       `json:"apiEndpoint,omitempty"`
	User        *ContextUser `json:"user,omitempty"`
	// Device provides information about the device used to send the request
	Device struct {
		DeviceID            string              `json:"deviceId,omitempty"`
		SupportedInterfaces map[string]struct{} `json:"supportedInterfaces,omitempty"`
	} `json:"device,omitempty"`
	Application ContextApplication `json:"application"`
	// Unit represents a logical construct organizing actors
	Unit struct {
		UnitId           string `json:"unitId"`
		PersistentUnitId string `json:"persistentUnitId"`
	} `json:"unit,omitempty"`
	// Person describes the person who is making the request to Alexa (user recognized by voice, not account)
	Person *ContextSystemPerson `json:"person,omitempty"`
}

type AudioPlayerActivity string

const (
	// AudioPlayerActivityIDLE Nothing was playing, no enqueued items.
	AudioPlayerActivityIDLE AudioPlayerActivity = "IDLE"
	// AudioPlayerActivityPAUSE Stream was paused.
	AudioPlayerActivityPAUSED AudioPlayerActivity = "PAUSED"
	// AudioPlayerActivityPLAYING Stream was playing.
	AudioPlayerActivityPLAYING AudioPlayerActivity = "PLAYING"

	// AudioPlayerActivityBufferUnderrun Buffer underrun
	AudioPlayerActivityBufferUnderrun AudioPlayerActivity = "BUFFER_UNDERRUN"
	// AudioPlayerActivityFINISHED Stream was finished playing.
	AudioPlayerActivityFINISHED AudioPlayerActivity = "FINISHED"
	// AudioPlayerActivitySTOPPED Stream was interrupted.
	AudioPlayerActivitySTOPPED AudioPlayerActivity = "STOPPED"
)

type ContextAudioPlayer struct {
	Token                string              `json:"token"`
	OffsetInMilliseconds int                 `json:"offsetInMilliseconds"`
	PlayerActivity       AudioPlayerActivity `json:"playerActivity"`
}

// ViewportExperience
type ViewportExperience struct {
	ArcMinuteWidth  int  `json:"arcMinuteWidth"`
	ArcMinuteHeight int  `json:"arcMinuteHeight"`
	CanRotate       bool `json:"canRotate"`
	CanResize       bool `json:"canResize"`
}

// ContextViewportMode is the mode for the device
type ContextViewportMode string

const (
	ContextViewportModeHUB    ContextViewportMode = "HUB"
	ContextViewportModeTV     ContextViewportMode = "TV"
	ContextViewportModePC     ContextViewportMode = "PC"
	ContextViewportModeMobile ContextViewportMode = "MOBILE"
	ContextViewportModeAuto   ContextViewportMode = "AUTO"
)

// ContextViewportShape is the shape of the device
type ContextViewportShape string

const (
	ContextViewportShapeRound     ContextViewportShape = "ROUND"
	ContextViewportShapeRectangle ContextViewportShape = "RECTANGLE"
)

// ContextViewport provides information about the viewport if the device has a screen
type ContextViewport struct {
	Experiences        []*ViewportExperience `json:"experiances,omitempty"`
	Mode               ContextViewportMode   `json:"mode"`
	Shape              ContextViewportShape  `json:"shape"`
	PixelWidth         int                   `json:"pixelWidth"`
	PixelHeight        int                   `json:"pixelHeight"`
	CurrentPixelWidth  int                   `json:"currentPixelWidth"`
	CurrentPixelHeight int                   `json:"currentPixelHeight"`
	DPI                int                   `json:"dpi"`
	Touch              []string              `json:"touch"`
	Keyboard           []string              `json:"keyboard"`
	Video              struct {
		Codecs []string `json:"codecs"`
	} `json:"video"`
}

type ViewportConfiguration struct {
	Video struct {
		Codecs []string `json:"codecs"`
	} `json:"video,omitempty"`
	Size struct {
		Type        string `json:"type"`
		PixelWidth  int    `json:"pixelWidth"`
		PixelHeight int    `json:"pixelHeight"`
	} `json:"size,omitempty"`
}
type ContextViewportType struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	Shape            string `json:"shape"`
	DPI              int    `json:"dpi"`
	PresentationType string `json:"presentationType"`
	CanRotate        bool   `json:"canRotate"`
	Configuration    struct {
		Current ViewportConfiguration `json:"current"`
	} `json:"configuration"`
}

// Context represents the Alexa skill request context.
type Context struct {
	System      *ContextSystem         `json:"System,omitempty"`
	AudioPlayer *ContextAudioPlayer    `json:"audioPlayer,omitempty"`
	Viewport    *ContextViewport       `json:"Viewport,omitempty"`
	Viewports   []*ContextViewportType `json:"Viewports,omitempty"`
}

// RequestEnvelope represents the alexa request envelope.
type RequestEnvelope struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Context *Context `json:"context"`
	Request *Request `json:"request"`
}
