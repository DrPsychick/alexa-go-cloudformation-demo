package alexa

// stolen from: https://github.com/soloworks/go-alexa-models

// RequestType enum
//type RequestType int
//
//const (
//	// RequestTypeUndefined means incoming value incorrect or not supported
//	RequestTypeUndefined RequestType = iota
//	// RequestTypeLaunch is constant `LaunchRequest`
//	RequestTypeLaunch
//	// RequestTypeIntent is constant `IntentRequest`
//	RequestTypeIntent
//	// RequestTypeSessionEnded is constant `SessionEndedRequest`
//	RequestTypeSessionEnded
//	// RequestTypeCanFulfillIntent is constant `CanFulfillIntentRequest`
//	RequestTypeCanFulfillIntent
//)
//
//var requestTypeStrings = [...]string{
//	"Undefined", // Placeholder - should never be this
//	"LaunchRequest",
//	"IntentRequest",
//	"SessionEndedRequest",
//	"CanFulfillIntentRequest",
//}

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

// Request we get from Alexa
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Body    struct {
		Type        RequestType `json:"type"`
		RequestID   string      `json:"requestId"`
		Timestamp   string      `json:"timestamp"`
		Locale      Locale      `json:"locale"`
		Intent      Intent      `json:"intent,omitempty"`
		Reason      string      `json:"reason,omitempty"`
		DialogState string      `json:"dialogState,omitempty"`
	} `json:"request"`
}

// TODO: make this work again (and use similar functions for skill, model, ...)
// MarshalJSON Function to handle JSON parsing out
//func (r RequestType) MarshalJSON() ([]byte, error) {
//	j := string(`"` + requestTypeStrings[r] + `"`)
//	return []byte(j), nil
//}

// UnmarshalJSON Function to handle JSON parsing out
//func (r *RequestType) UnmarshalJSON(data []byte) error {
//	rt := RequestTypeUndefined
//	// Convert to string whilst removing quotes
//	x := string(data)[1 : len(data)-1]
//	// Find the type in the range of values
//	for i, s := range requestTypeStrings {
//		if strings.ToLower(s) == strings.ToLower(x) {
//			rt = RequestType(i)
//		}
//	}
//	*r = rt
//	return nil
//}
