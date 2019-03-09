package alexa

// stolen from: https://github.com/soloworks/go-alexa-models

type RequestType int

const (
	// RequestTypeUndefined means incoming value incorrect or not supported
	RequestTypeUndefined RequestType = iota
	// RequestTypeLaunch is constant `LaunchRequest`
	RequestTypeLaunch
	// RequestTypeIntent is constant `IntentRequest`
	RequestTypeIntent
	// RequestTypeSessionEnded is constant `SessionEndedRequest`
	RequestTypeSessionEnded
	// RequestTypeCanFulfillIntent is constant `CanFulfillIntentRequest`
	RequestTypeCanFulfillIntent
)

var requestTypeStrings = [...]string{
	"Undefined", // Placeholder - should never be this
	"LaunchRequest",
	"IntentRequest",
	"SessionEndedRequest",
	"CanFulfillIntentRequest",
}

type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Body    struct {
		Type        RequestType `json:"type"`
		RequestID   string      `json:"requestId"`
		Timestamp   string      `json:"timestamp"`
		Locale      string      `json:"locale"`
		Intent      Intent      `json:"intent,omitempty"`
		Reason      string      `json:"reason,omitempty"`
		DialogState string      `json:"dialogState,omitempty"`
	} `json:"request"`
}
