package alexa

import "errors"

// Error constants
var (
	ErrorUnknown                   = errors.New("unknown error")
	ErrorNoIntent                  = errors.New("request has no intent")
	ErrorNoSlot                    = errors.New("slot does not exist")
	ErrorSlotNoResolutions         = errors.New("slot has no resolutions")
	ErrorSlotNoResolutionWithMatch = errors.New("no resolution with match")
	ErrorNoSystemInContext         = errors.New("no system in context")
	ErrorNoUserInContext           = errors.New("no user in context")
	ErrorNoPersonInContext         = errors.New("no person in system context")
	ErrorNoApplicationID           = errors.New("no application ID in the request")
)

// RequestLocale represents the locale of the request
type RequestLocale string

// Locale constants.
const (
	// LocaleAmericanEnglish is the locale for American English.
	LocaleAmericanEnglish RequestLocale = "en-US"

	// LocaleAustralianEnglish is the locale for Australian English.
	LocaleAustralianEnglish RequestLocale = "en-AU"

	// LocaleBritishEnglish is the locale for UK English.
	LocaleBritishEnglish RequestLocale = "en-GB"

	// LocaleCanadianEnglish is the locale for Canadian English.
	LocaleCanadianEnglish RequestLocale = "en-CA"

	// LocaleCanadianFrench is the locale for Canadian French.
	LocaleCanadianFrench RequestLocale = "fr-CA"

	// LocaleFrench is the locale for French (France).
	LocaleFrench RequestLocale = "fr-FR"

	// LocaleGerman is the locale for standard dialect German (Germany).
	LocaleGerman RequestLocale = "de-DE"

	//LocaleIndianEnglish is the locale for Indian English.
	LocaleIndianEnglish RequestLocale = "en-IN"

	// LocaleItalian is the locale for Italian (Italy).
	LocaleItalian RequestLocale = "it-IT"

	// LocaleJapanese is the locale for Japanese (Japan).
	LocaleJapanese RequestLocale = "ja-JP"

	// LocaleMexicanSpanish is the locale for Mexican Spanish.
	LocaleMexicanSpanish RequestLocale = "es-MX"

	// LocaleSpanish is the  for Spanish (Spain).
	LocaleSpanish RequestLocale = "es-ES"
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

// Intent returns the intent or an empty intent
func (r *RequestEnvelope) Intent() (Intent, error) {
	i := r.Request.Intent
	if i.Name == "" {
		return Intent{}, ErrorNoIntent
	}

	return i, nil
}

// IntentName returns the name of the intent or "" if it's no intent request
func (r *RequestEnvelope) IntentName() string {
	i, err := r.Intent()
	if err != nil {
		return ""
	}
	return i.Name
}

// IsIntentConfirmed returns true if the confirmation status is CONFIRMED
func (r *RequestEnvelope) IsIntentConfirmed() bool {
	i, err := r.Intent()
	if err != nil {
		return false
	}

	return i.ConfirmationStatus == ConfirmationStatusConfirmed
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

// Slots returns the list of slots, an empty list if no intent was found
func (r *RequestEnvelope) Slots() map[string]*Slot {
	i, err := r.Intent()
	if err != nil {
		return map[string]*Slot{}
	}

	if i.Slots == nil {
		return map[string]*Slot{}
	}

	return i.Slots
}

// Slot returns the named slot or an error
func (r *RequestEnvelope) Slot(name string) (Slot, error) {
	i, err := r.Intent()
	if err != nil {
		return Slot{}, err
	}

	s, ok := i.Slots[name]
	if !ok {
		return Slot{}, ErrorNoSlot
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

// SlotResolutionsPerAuthority returns the list of ResolutionsPerAuthority
func (s *Slot) SlotResolutionsPerAuthority() ([]*PerAuthority, error) {
	if s.Resolutions == nil {
		return []*PerAuthority{}, ErrorSlotNoResolutions
	}

	return s.Resolutions.PerAuthority, nil
}

// FirstAuthorityWithMatch returns the first authority with ResolutionStatusMatch
func (s *Slot) FirstAuthorityWithMatch() (PerAuthority, error) {
	auths, err := s.SlotResolutionsPerAuthority()
	if err != nil {
		return PerAuthority{}, err
	}

	for _, a := range auths {
		if a.Status != nil && a.Status.Code == ResolutionStatusMatch {
			return *a, nil
		}
	}

	return PerAuthority{}, ErrorSlotNoResolutionWithMatch
}

// AuthorityValueValue points to the unique ID and value
type AuthorityValueValue struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

// AuthorityValue is an entry in the list of values
type AuthorityValue struct {
	Value *AuthorityValueValue `json:"value,omitempty"`
}

// StatusCode represents the status code of a slot resolution
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

// PerAuthority encapsulates an Authority which is the source of the data provided
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
	// TypeLaunchRequest string that represents a launch request
	TypeLaunchRequest RequestType = "LaunchRequest"
	// TypeIntentRequest string that represents an intent request
	TypeIntentRequest RequestType = "IntentRequest"
	// TypeSessionEndedRequest string that represents a session end request
	TypeSessionEndedRequest RequestType = "SessionEndedRequest"
	// TypeCanFulfillIntentRequest string that represents a can fulfill intent request
	TypeCanFulfillIntentRequest RequestType = "CanFulfillIntentRequest"
)

func (r *RequestEnvelope) RequestType() RequestType {
	if r.Request == nil {
		return ""
	}
	return r.Request.Type
}

func (r *RequestEnvelope) IsIntentRequest() bool {
	if r.Request == nil || r.Request.Type == "" {
		return false
	}
	return r.Request.Type == TypeIntentRequest
}

func (r *RequestEnvelope) RequestLocale() string {
	if r.Request == nil {
		return ""
	}
	return string(r.Request.Locale)
}

func (r *RequestEnvelope) RequestDialogState() DialogStateType {
	if r.Request == nil {
		return ""
	}

	return r.Request.DialogState
}

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
	Locale      RequestLocale   `json:"locale"`
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

// ContextApplication is used to verify that the request was intended for your service, the ID is the application ID for your skill
type ContextApplication struct {
	ApplicationID string `json:"applicationId"`
}

// ApplicationID returns the application ID from the session first, then system or throw an error. Use it to verify the request is meant for your Skill.
func (r *RequestEnvelope) ApplicationID() (string, error) {
	// Session or System
	if r.Session == nil {
		s, err := r.System()
		if err != nil || s.Application == nil {
			return "", ErrorNoApplicationID
		}

		return s.Application.ApplicationID, nil
	}

	if r.Session.Application == nil {
		return "", ErrorNoApplicationID
	}

	return r.Session.Application.ApplicationID, nil
}

// Session represents the Alexa skill session.
type Session struct {
	New         bool                   `json:"new"`
	SessionID   string                 `json:"sessionId"`
	Application *ContextApplication    `json:"application"`
	Attributes  map[string]interface{} `json:"attributes"`
	User        *ContextUser           `json:"user"`
}

// SessionID returns the sessionID or an empty string
func (r *RequestEnvelope) SessionID() string {
	if r.Session == nil {
		return ""
	}
	return r.Session.SessionID

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
	Application *ContextApplication `json:"application"`
	// Unit represents a logical construct organizing actors
	Unit struct {
		UnitId           string `json:"unitId"`
		PersistentUnitId string `json:"persistentUnitId"`
	} `json:"unit,omitempty"`
	// Person describes the person who is making the request to Alexa (user recognized by voice, not account)
	Person *ContextSystemPerson `json:"person,omitempty"`
}

// System returns the system object if it exists in the context
func (r *RequestEnvelope) System() (*ContextSystem, error) {
	if r.Context == nil || r.Context.System == nil {
		return &ContextSystem{}, ErrorNoSystemInContext
	}

	return r.Context.System, nil
}

// ContextPerson returns the person in the context or throw an error if no person exists
func (r *RequestEnvelope) ContextPerson() (*ContextSystemPerson, error) {
	s, err := r.System()
	if err != nil || s.Person == nil {
		return &ContextSystemPerson{}, ErrorNoPersonInContext
	}
	return r.Context.System.Person, nil
}

// ContextUser returns the user in the context or throw an error if no user exists
func (r *RequestEnvelope) ContextUser() (*ContextUser, error) {
	s, err := r.System()
	if err != nil || s.User == nil {
		return &ContextUser{}, ErrorNoUserInContext
	}
	return r.Context.System.User, nil
}

type AudioPlayerActivity string

const (
	// AudioPlayerActivityIDLE Nothing was playing, no enqueued items.
	AudioPlayerActivityIDLE AudioPlayerActivity = "IDLE"
	// AudioPlayerActivityPAUSED Stream was paused.
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

// ViewportExperience has info about the device
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
	Experiences        []*ViewportExperience `json:"experiences,omitempty"`
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
