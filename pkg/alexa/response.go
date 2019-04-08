package alexa

// Payload of response
type Payload struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Text    string `json:"text,omitempty"`
	SSML    string `json:"ssml,omitempty"`
	Content string `json:"content,omitempty"`
	Image   Image  `json:"image,omitempty"`
}

// Image definition
type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// Reprompt
type Reprompt struct {
	OutputSpeech Payload `json:"outputSpeech,omitempty"`
}

// Response is the response back to the response speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResponseBody           `json:"response"`

	Error error `json:"-"`
}

// ResponseBody contains Speech Card etc.
type ResponseBody struct {
	OutputSpeech     *Payload     `json:"outputSpeech,omitempty"`
	Card             *Payload     `json:"card,omitempty"`
	Reprompt         *Reprompt    `json:"reprompt,omitempty"`
	Directives       []Directives `json:"directives,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
}

// NewEmptyResponse builds an empty response.
func NewEmptyResponse() Response {
	return Response{Version: "1.0"}
}

// NewTerminateResponse builds an empty response that terminates the session.
func NewTerminateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true

	return r
}

// NewSpeechResponse builds a response with the given speech.
func NewSpeechResponse(speech string) Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = true
	r.Body.OutputSpeech = &Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple response to advance to the next step.
func NewDialogDelegateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = false
	r.Body.Directives = append(r.Body.Directives, Directives{Type: DirectiveTypeDialogDelegate})

	return r
}

// NewSimpleResponse builds a response with the given title and text.
func NewSimpleResponse(title string, text string) Response {
	r := Response{
		Version: "1.0",
		Body: ResponseBody{
			OutputSpeech: &Payload{
				Type: "PlainText",
				Text: text,
			},
			Card: &Payload{
				Type:    "Simple",
				Title:   title,
				Content: text,
			},
			ShouldEndSession: true,
		},
	}

	return r
}

// NewErrorResponse creates an error response with the given error.
func NewErrorResponse(err error) Response {
	return Response{
		Version: "1.0",
		Error:   err,
	}
}
