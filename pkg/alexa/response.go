package alexa

type Payload struct {
	Type    string `json:"type,omitempty"`
	Title   string `json:"title,omitempty"`
	Text    string `json:"text,omitempty"`
	SSML    string `json:"ssml,omitempty"`
	Content string `json:"content,omitempty"`
	Image   Image  `json:"image,omitempty"`
}

type Image struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

type Reprompt struct {
	OutputSpeech Payload `json:"outputSpeech,omitempty"`
}

// Response is the response back to the response speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              ResponseBody           `json:"response"`
}

type ResponseBody struct {
	OutputSpeech     *Payload     `json:"outputSpeech,omitempty"`
	Card             *Payload     `json:"card,omitempty"`
	Reprompt         *Reprompt    `json:"reprompt,omitempty"`
	Directives       []Directives `json:"directives,omitempty"`
	ShouldEndSession bool         `json:"shouldEndSession"`
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
	r.Body.OutputSpeech = &Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple response response to advance to the next step
func NewDialogDelegateResponse() Response {
	r := NewEmptyResponse()
	r.Body.ShouldEndSession = false
	r.Body.Directives = append(r.Body.Directives, Directives{Type: DirectiveTypeDialogDelegate})

	return r
}

// NewSimpleResponse builds a session response
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
