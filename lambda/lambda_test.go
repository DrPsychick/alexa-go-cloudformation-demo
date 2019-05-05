package lambda_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/logger"
	"github.com/hamba/statter/l2met"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLambda_HandleSaySomething2(t *testing.T) {
	// app with stdout logger and l2met stats
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	r := &alexa.Request{
		Locale: "de-DE",
		Type:   alexa.TypeIntentRequest,
		Intent: alexa.Intent{
			Name: loca.SaySomething,
		},
	}
	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, resp.Response.Card.Content)
	//assert.NotEmpty(t, resp.Response.OutputSpeech.SSML)
}

func TestLambda_HandleSaySomething2_ErrorNoLocale(t *testing.T) {
	loca.Registry = l10n.NewRegistry()
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	r := &alexa.Request{
		Locale: "en-US",
		Type:   alexa.TypeIntentRequest,
		Intent: alexa.Intent{
			Name: loca.SaySomething,
		},
	}
	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)

	m.Serve(b, r)
	resp := b.Build()

	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, "locale en-US not found")
}

func TestLambda_HandleSaySomething2_ErrorNoTranslation(t *testing.T) {
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"Translation error"})
	loc.Set(l10n.KeyErrorNoTranslationText, []string{"No translation for '%s'"})
	loc.Set(l10n.KeyErrorNoTranslationSSML, []string{l10n.Speak("I am missing translations.")})
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	r := &alexa.Request{
		Locale: "en-US",
		Type:   alexa.TypeIntentRequest,
		Intent: alexa.Intent{
			Name: loca.SaySomething,
			Slots: map[string]alexa.Slot{
				"AWSArea": {
					Name:  "Area",
					Value: "Europa",
				},
				"AWSRegion": {
					Name:  "Region",
					Value: "Frankfurt",
				},
			},
		},
	}
	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)

	m.Serve(b, r)
	resp := b.Build()

	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, loca.SaySomething)
	assert.Contains(t, resp.Response.OutputSpeech.SSML, "missing translations")
}
