package lambda_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/drpsychick/go-alexa-lambda/ssml"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
	"github.com/stretchr/testify/assert"
	"testing"
)

func initLocaleRegistry(t *testing.T) {
	loca.Registry = l10n.NewRegistry()
	err := loca.Registry.Register(&l10n.Locale{Name: "en-US", TextSnippets: l10n.Snippets{}})
	assert.NoError(t, err)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	loc.Set(l10n.KeyErrorTitle, []string{"error"})
	loc.Set(l10n.KeyErrorText, []string{"An error occurred: %s"})
	loc.Set(l10n.KeyErrorSSML, []string{ssml.Speak("An error occurred.")})
	loc.Set(l10n.KeyErrorLocaleNotFoundTitle, []string{"error"})
	loc.Set(l10n.KeyErrorLocaleNotFoundText, []string{"Locale '%s' not found!"})
	loc.Set(l10n.KeyErrorLocaleNotFoundSSML, []string{"<speak>Locale '%s' not found!<speak>"})
	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"error"})
	loc.Set(l10n.KeyErrorNoTranslationText, []string{"Key '%s' not found!"})
	loc.Set(l10n.KeyErrorNoTranslationSSML, []string{"<speak>Key '%s' not found!<speak>"})
	loc.Set(l10n.KeyErrorTranslationTitle, []string{"Translation error"})
	loc.Set(l10n.KeyErrorTranslationText, []string{"An error occurred in translation. The developer is informed."})
	loc.Set(l10n.KeyErrorTranslationSSML, []string{"<speak>An error occurred during translation. The developer is informed.<speak>"})
}

func TestLambda_HandleLaunch(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	assert.NotEmpty(t, loc)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeLaunchRequest,
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	m.Serve(b, r)
	resp := b.Build()

	// locale not found
	assert.NotEmpty(t, resp.Response.Card)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// now with locale
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyLaunchTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// now with loca
	loc.Set(l10n.KeyLaunchTitle, []string{"Start"})
	loc.Set(l10n.KeyLaunchText, []string{"Und los..."})
	loc.Set(l10n.KeyLaunchSSML, []string{ssml.Speak("Und los.")})
	assert.NotEmpty(t, loc.Get(l10n.KeyLaunchTitle))
	assert.NotEmpty(t, loc.GetAny(l10n.KeyLaunchText))
	assert.NotEmpty(t, loc.GetAny(l10n.KeyLaunchSSML))

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyLaunchTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyLaunchText), resp.Response.Card.Content)
}

func TestLambda_HandleEnd(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeSessionEndedRequest,
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// with existing locale, but missing text
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyStopTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with translations
	loc.Set(l10n.KeyStopTitle, []string{"Stop"})
	loc.Set(l10n.KeyStopText, []string{"Alright, it's over now."})
	loc.Set(l10n.KeyStopSSML, []string{ssml.Speak("Alright, it's over now.")})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyStopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyStopText), resp.Response.Card.Content)
}

func TestLambda_HandleHelp(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: alexa.HelpIntent,
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// with existing locale, but missing text
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyHelpTitle)
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	loc.ResetErrors()

	// with translations
	loc.Set(l10n.KeyHelpTitle, []string{"Help"})
	loc.Set(l10n.KeyHelpText, []string{"I'd love to help you"})
	loc.Set(l10n.KeyHelpSSML, []string{ssml.Speak("I'd love to help you")})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyHelpTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyHelpText), resp.Response.Card.Content)
}

func TestLambda_HandleStop(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: alexa.StopIntent,
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// with existing locale, but missing text
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	assert.Contains(t, loc.GetErrors()[0].Error(), l10n.KeyStopTitle)
	loc.ResetErrors()

	// with translations
	loc.Set(l10n.KeyStopTitle, []string{"Stop"})
	loc.Set(l10n.KeyStopText, []string{"Alright, it's over now."})
	loc.Set(l10n.KeyStopSSML, []string{ssml.Speak("Alright, it's over now.")})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(l10n.KeyStopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(l10n.KeyStopText), resp.Response.Card.Content)
}

func TestLambda_HandleSaySomething2(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)
	personId := "John"

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Context: &alexa.Context{
			System: &alexa.ContextSystem{
				Person: &alexa.ContextSystemPerson{
					PersonID: personId,
				},
			},
		},
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: loca.SaySomething,
			},
		},
	}
	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, loca.SaySomethingUserText)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.SaySomethingUserTitle, []string{"Hi %s!"})
	loc.Set(loca.SaySomethingUserText, []string{"Sadly, I have nothing to tell you %s."})
	loc.Set(loca.SaySomethingUserSSML, []string{ssml.Speak(ssml.UseVoiceLang("Kendra", "en-US", "%s do you like the Autobahn?"))})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.SaySomethingUserTitle, personId), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.SaySomethingUserText, personId), resp.Response.Card.Content)
}

func TestLambda_HandleAWSStatus(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: loca.AWSStatus,
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "Error", resp.Response.Card.Title)
	assert.Equal(t, resp.Response.Card.Content, "Locale not found!")

	// with existing locale, but missing text
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	assert.Contains(t, loc.GetErrors()[0].Error(), loca.AWSStatusTitle)
	loc.ResetErrors()

	// with translations
	loc.Set(loca.AWSStatusTitle, []string{"Status"})
	loc.Set(loca.AWSStatusAreaElicitText, []string{"Elicit Area"})
	loc.Set(loca.AWSStatusAreaElicitSSML, []string{"<speak>Elicit Area<speak>"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.AWSStatusTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.AWSStatusAreaElicitText), resp.Response.Card.Content)
}

func TestLambda_HandleAWSStatus_WithSlots(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "en-US",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: loca.AWSStatus,
				Slots: map[string]*alexa.Slot{
					loca.TypeAreaName: {
						Name:  loca.TypeAreaName,
						Value: "Europe",
						// required!
						Resolutions: &alexa.Resolutions{
							ResolutionsPerAuthority: []*alexa.PerAuthority{
								{
									Authority: "",
									Status: &alexa.ResolutionStatus{
										Code: "ER_SUCCESS_MATCH",
									},
									Values: []*alexa.AuthorityValue{
										{
											Value: &alexa.AuthorityValueValue{
												Name: "Europe",
												ID:   "4312d5c8cdda027420c474e2221abc34",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	sb := skill.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.AWSStatusTitle, []string{"Status"})
	loc.Set(loca.AWSStatusRegionElicitText, []string{"Elicit Region"})
	loc.Set(loca.AWSStatusRegionElicitSSML, []string{"<speak>Elicit Region<speak>"})

	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.AWSStatusTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.AWSStatusRegionElicitText), resp.Response.Card.Content)

	// with Region and missing loca
	r.Request.Intent.Slots[loca.TypeRegionName] = &alexa.Slot{
		Name:  loca.TypeRegionName,
		Value: "Frankfurt",
		// required!
		Resolutions: &alexa.Resolutions{
			ResolutionsPerAuthority: []*alexa.PerAuthority{
				{
					Authority: "",
					Status: &alexa.ResolutionStatus{
						Code: "ER_SUCCESS_MATCH",
					},
					Values: []*alexa.AuthorityValue{
						{
							Value: &alexa.AuthorityValueValue{
								Name: "Frankfurt",
								ID:   "asdf",
							},
						},
					},
				},
			},
		},
	}

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, loc.GetErrors())
	assert.Equal(t, "Translation error", resp.Response.Card.Title)
	assert.Contains(t, loc.GetErrors()[0].Error(), loca.AWSStatusText)
	loc.ResetErrors()

	// with Region and loca
	loc.Set(loca.AWSStatusText, []string{"Everything alright in %s %s"})
	loc.Set(loca.AWSStatusSSML, []string{"<speak>All good</speak>"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Empty(t, resp.Response.Card.Text)
	assert.Equal(t, loc.Get(loca.AWSStatusTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.AWSStatusText, "Europe", "Frankfurt"), resp.Response.Card.Content)
}
