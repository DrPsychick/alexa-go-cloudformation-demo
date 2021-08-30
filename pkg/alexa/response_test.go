package alexa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWith_Functions(t *testing.T) {
	b := &ResponseBuilder{}

	b.WithSpeech("speech")
	b.WithSimpleCard("title", "text")
	b.WithShouldEndSession(true)
	b.WithReprompt("reprompt")

	res := b.Build()
	assert.Equal(t, "title", res.Response.Card.Title)
	assert.Equal(t, "text", res.Response.Card.Content)
	assert.Equal(t, "speech", res.Response.OutputSpeech.Text)
	assert.Equal(t, "reprompt", res.Response.Reprompt.OutputSpeech.Text)

	b.WithSpeech(l10n.Speak("speech"))
	b.WithReprompt(l10n.Speak("reprompt"))
	res = b.Build()
	assert.Equal(t, l10n.Speak("speech"), res.Response.OutputSpeech.SSML)
	assert.Equal(t, l10n.Speak("reprompt"), res.Response.Reprompt.OutputSpeech.SSML)

	b.WithStandardCard("title", "text", &Image{})
	res = b.Build()
	assert.Equal(t, "title", res.Response.Card.Title)
	assert.Equal(t, "text", res.Response.Card.Text)
	assert.Equal(t, &Image{}, res.Response.Card.Image)
	assert.Empty(t, res.Response.Card.Content)
}

func TestCanFulfillIntent(t *testing.T) {
	b := ResponseBuilder{}
	b.WithCanFulfillIntent(&CanFulfillIntent{
		CanFulfill: string(TypeLaunchRequest),
		Slots:      map[string]CanFulfillSlot{},
	})

	res := b.Build()
	assert.Equal(t, string(TypeLaunchRequest), res.Response.CanFulfillIntent.CanFulfill)
}

func TestAddDirective(t *testing.T) {
	b := &ResponseBuilder{}

	b.AddDirective(&Directive{
		Type:          DirectiveTypeDialogDelegate,
		SlotToElicit:  "",
		UpdatedIntent: nil,
		PlayBehavior:  "",
		AudioItem:     nil,
	})
	res := b.Build()

	assert.Equal(t, DirectiveTypeDialogDelegate, res.Response.Directives[0].Type)
}

func TestSessionAttributes(t *testing.T) {
	b := &ResponseBuilder{}

	b.WithSessionAttributes(map[string]interface{}{
		"foo": "Bar",
	})
	res := b.Build()
	assert.Equal(t, "Bar", res.SessionAttributes["foo"])
}
