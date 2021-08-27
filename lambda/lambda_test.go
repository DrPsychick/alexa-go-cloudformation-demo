package lambda_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
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

	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"error"})
	loc.Set(l10n.KeyErrorNoTranslationText, []string{"Key '%s' not found!"})

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

	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// now with locale
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Contains(t, resp.Response.Card.Content, loca.LaunchText)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// now with loca
	loc.Set(loca.LaunchTitle, []string{"Start"})
	loc.Set(loca.LaunchText, []string{"Und los..."})
	assert.NotEmpty(t, loc.Get(loca.LaunchTitle))
	assert.NotEmpty(t, loc.GetAny(loca.LaunchText))

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.LaunchTitle), resp.Response.Card.Title)
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

	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Contains(t, resp.Response.Card.Content, loca.Stop)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.StopTitle, []string{"Stop"})
	loc.Set(loca.Stop, []string{"Alright, it's over now."})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.StopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.Stop), resp.Response.Card.Content)
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

	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Contains(t, resp.Response.Card.Content, loca.Help)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.HelpTitle, []string{"Help"})
	loc.Set(loca.Help, []string{"I'd love to help you"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.HelpTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.Help), resp.Response.Card.Content)
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

	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, loca.Stop)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.StopTitle, []string{"Stop"})
	loc.Set(loca.Stop, []string{"Alright, it's over now."})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.StopTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.Stop), resp.Response.Card.Content)
}

func TestLambda_HandleSaySomething2(t *testing.T) {
	initLocaleRegistry(t)

	app := alfalfa.NewApplication(log.Null, stats.Null)

	r := &alexa.RequestEnvelope{
		Version: "1.0",
		Request: &alexa.Request{
			Locale: "de-DE",
			Type:   alexa.TypeIntentRequest,
			Intent: alexa.Intent{
				Name: loca.SaySomething,
			},
		},
	}
	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Equal(t, "Locale 'de-DE' is not supported!", resp.Response.Card.Content)

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, "Error")

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.SaySomethingTitle, []string{"?"})
	loc.Set(loca.SaySomethingText, []string{"Sadly, I have nothing to say."})
	loc.Set(loca.SaySomethingSSML, []string{l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "I like the Autobahn, it's so geil"))})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, loc.Get(loca.SaySomethingTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.SaySomethingText), resp.Response.Card.Content)
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

	sb := gen.NewSkillBuilder()
	b := &alexa.ResponseBuilder{}
	m := lambda.NewMux(app, sb)
	m = middleware.WithRequestStats(m, app)

	// missing locale
	m.Serve(b, r)
	resp := b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)

	// with existing locale, but missing text
	r.Request.Locale = "en-US"
	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, loca.AWSStatusAreaElicitSSML)

	// with translations
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)
	loc.Set(loca.AWSStatusTitle, []string{"Status"})
	//loc.Set(loca.AWSStatusText, []string{"Everything alright!"})
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
							PerAuthority: []*alexa.PerAuthority{
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

	sb := gen.NewSkillBuilder()
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
			PerAuthority: []*alexa.PerAuthority{
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
	assert.Equal(t, "error", resp.Response.Card.Title)
	assert.Contains(t, resp.Response.Card.Content, loca.AWSStatusSSML)

	// with Region and loca
	loc.Set(loca.AWSStatusText, []string{"Everything alright in %s %s"})
	loc.Set(loca.AWSStatusSSML, []string{"<speak>All good</speak>"})

	m.Serve(b, r)
	resp = b.Build()

	assert.NotEmpty(t, resp)
	assert.Empty(t, resp.Response.Card.Content)
	assert.Equal(t, loc.Get(loca.AWSStatusTitle), resp.Response.Card.Title)
	assert.Equal(t, loc.Get(loca.AWSStatusText, "Europe", "Frankfurt"), resp.Response.Card.Text)
}
