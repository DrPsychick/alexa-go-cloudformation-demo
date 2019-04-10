package alexa_test

import (
	"testing"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/stretchr/testify/assert"
)

func TestResponseBuilder_WithSpeechPlainText(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithSpeech("test")

	resp := bdr.Build()
	want := &alexa.OutputSpeech{Type: "PlainText", Text: "test"}
	assert.Exactly(t, bdr, b)
	assert.Equal(t, want, resp.Response.OutputSpeech)
}

func TestResponseBuilder_WithSpeechSSML(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithSpeech("<speak>test</speak>")

	resp := bdr.Build()
	want := &alexa.OutputSpeech{Type: "SSML", Text: "<speak>test</speak>"}
	assert.Exactly(t, bdr, b)
	assert.Equal(t, want, resp.Response.OutputSpeech)
}
