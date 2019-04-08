package alexa

// stolen from: https://github.com/soloworks/go-alexa-models

// RequestType represents JSON request `request.type`, see https://developer.amazon.com/docs/custom-skills/request-types-reference.html
type RequestType string

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

// RequestBody represents the information about the request.
type RequestBody struct {
	Type        RequestType `json:"type"`
	RequestID   string      `json:"requestId"`
	Timestamp   string      `json:"timestamp"`
	Locale      Locale      `json:"locale"`
	Intent      Intent      `json:"intent,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	DialogState string      `json:"dialogState,omitempty"`
}

// Request represents the alexa request.
type Request struct {
	Version string      `json:"version"`
	Session Session     `json:"session"`
	Context Context     `json:"context"`
	Body    RequestBody `json:"request"`
}
