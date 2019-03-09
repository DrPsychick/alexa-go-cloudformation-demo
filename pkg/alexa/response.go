package alexa

import (
	"github.com/arienmalec/alexa-go"
)

// Response is the response back to the response speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              struct {
		OutputSpeech     *alexa.Payload     `json:"outputSpeech,omitempty"`
		Card             *alexa.Payload     `json:"card,omitempty"`
		Reprompt         *alexa.Reprompt    `json:"reprompt,omitempty"`
		Directives       []alexa.Directives `json:"directives,omitempty"`
		ShouldEndSession bool               `json:"shouldEndSession"`
	} `json:"response"`
}

// NewEmptyResponse builds an empty response
func NewEmptyResponse() Response {
	return Response{Version: "1.0"}
}

// NewSimpleTerminateResponse builds an empty response
func NewSimpleTerminateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true
	return r
}

// NewSpeechResponse builds a simple speech response
func NewSpeechResponse(speech string) Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true
	r.Body.OutputSpeech = &alexa.Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple response response to advance to the next step
func NewDialogDelegateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = false
	r.Body.Directives = append(r.Body.Directives, alexa.Directives{Type: DirectiveTypeDialogDelegate})

	return r
}

//NewSimpleResponse builds a session response
func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &alexa.Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}
