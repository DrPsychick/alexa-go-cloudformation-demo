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
	want := &alexa.OutputSpeech{Type: "SSML", SSML: "<speak>test</speak>"}
	assert.Exactly(t, bdr, b)
	assert.Equal(t, want, resp.Response.OutputSpeech)
}

func TestResponseBuilder_WithSimpleCard(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithSimpleCard("test title", "test text")

	resp := bdr.Build()
	want := &alexa.Card{Type: "Simple", Title: "test title", Content: "test text"}
	assert.Exactly(t, bdr, b)
	assert.Equal(t, want, resp.Response.Card)
}

func TestResponseBuilder_WithStandardCard(t *testing.T) {
	img := &alexa.Image{}

	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithStandardCard("test title", "test text", img)

	resp := bdr.Build()
	want := &alexa.Card{Type: "Standard", Title: "test title", Text: "test text", Image: img}
	assert.Exactly(t, bdr, b)
	assert.Equal(t, want, resp.Response.Card)
}

func TestResponseBuilder_WithCard(t *testing.T) {
	card := &alexa.Card{}

	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithCard(card)

	resp := bdr.Build()
	assert.Exactly(t, bdr, b)
	assert.Equal(t, card, resp.Response.Card)
}

func TestResponseBuilder_WithShouldEndSession(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithShouldEndSession(true)

	resp := bdr.Build()
	assert.Exactly(t, bdr, b)
	assert.True(t, resp.Response.ShouldEndSession)
}

func TestResponseBuilder_WithSessionAttributes(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithSessionAttributes(map[string]interface{}{"foo": "bar"})

	resp := bdr.Build()
	assert.Exactly(t, bdr, b)
	assert.Equal(t, map[string]interface{}{"foo": "bar"}, resp.SessionAttributes)
}

func TestResponseBuilder_WithCanFulfillIntent(t *testing.T) {
	fulfill := &alexa.CanFulfillIntent{CanFulfill: "YES"}

	bdr := &alexa.ResponseBuilder{}

	b := bdr.WithCanFulfillIntent(fulfill)

	resp := bdr.Build()
	assert.Exactly(t, bdr, b)
	assert.Equal(t, fulfill, resp.Response.CanFulfillIntent)
}

func TestResponseBuilder_AddDirective(t *testing.T) {
	directive := &alexa.Directive{}

	bdr := &alexa.ResponseBuilder{}

	b := bdr.AddDirective(directive)

	resp := bdr.Build()
	assert.Exactly(t, bdr, b)
	assert.Equal(t, []*alexa.Directive{directive}, resp.Response.Directives)
}
