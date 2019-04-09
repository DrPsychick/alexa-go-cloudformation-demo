package lambda

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

const (
	SSMLDemoIntent     = "SSMLDemoIntent"
	SaySomethingIntent = "SaySomething"
	DemoIntent         = "DemoIntent"
)

type Application interface {
	Help() (string, string)
	Stop(l *l10n.Locale) (string, string, string)
	SSMLDemo(l *l10n.Locale) (string, string, string)
	SaySomething(l *l10n.Locale) (string, string, string)
}

func NewMux(app Application) alexa.Handler {
	mux := alexa.NewServerMux()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch)
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)

	mux.HandleIntent(alexa.HelpIntent, handleHelp(app))
	mux.HandleIntent(alexa.CancelIntent, handleStop(app))
	mux.HandleIntent(alexa.StopIntent, handleStop(app))

	mux.HandleIntent(SSMLDemoIntent, handleSSMLResponse(app))
	mux.HandleIntent(SaySomethingIntent, handleSaySomethingResponse(app))
	mux.HandleIntentFunc(DemoIntent, handleDemo)

	return mux
}

func handleLaunch(b *alexa.ResponseBuilder, r *alexa.Request) {
	title := "Launch"
	//text := "Willkommen beim Karlsruhe Golang Meetup #3!"
	text := "Ja?"

	b.WithSpeech(text).
		WithSimpleCard(title, text).
		WithShouldEndSession(false)
}

func handleCanFulfillIntent(b *alexa.ResponseBuilder, r *alexa.Request) {
	intent := r.Body.Intent.Name
	if intent == SSMLDemoIntent || intent == SaySomethingIntent || intent == DemoIntent {
		b.WithCanFulfillIntent(&alexa.CanFulfillIntent{
			CanFulfill: "YES",
		})
		return
	}

	b.WithCanFulfillIntent(&alexa.CanFulfillIntent{
		CanFulfill: "NO",
	})
}

func handleHelp(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		title, text := app.Help()

		b.WithSpeech(text).
			WithSimpleCard(title, text)
	})
}

func handleStop(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(string(r.Body.Locale))
		if err != nil {
			// TODO: maybe say something here
			return
		}

		title, text, _ := app.Stop(l)

		b.WithSpeech(text).
			WithSimpleCard(title, text)
	})
}

func handleSSMLResponse(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(string(r.Body.Locale))
		if err != nil {
			// TODO: maybe say something here
			return
		}

		title, text, ssmlText := app.SSMLDemo(l)

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

func handleSaySomethingResponse(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(string(r.Body.Locale))
		if err != nil {
			return
		}

		title, text, ssmlText := app.SaySomething(l)

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

func handleDemo(b *alexa.ResponseBuilder, r *alexa.Request) {
	title := "Test"
	text := "Pace ist geil!"
	ssml := `<speak>
		<voice name="Kendra"><lang xml:lang="en-US"><emphasis level="strong">pace</emphasis></lang></voice>
		<voice name="Marlene">iss <emphasis level="strong">geil!</emphasis></voice>
	</speak>`

	b.WithSpeech(ssml).
		WithSimpleCard(title, text)
}
