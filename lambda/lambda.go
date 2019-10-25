package lambda

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

const (
	SSMLDemoIntent     = "SSMLDemoIntent"
	SaySomethingIntent = "SaySomething"
	DemoIntent         = "DemoIntent"
)

type Application interface {
	log.Loggable
	stats.Statable

	Launch(l l10n.LocaleInstance) (string, string)
	Help(l l10n.LocaleInstance) (string, string, string)
	Stop(l l10n.LocaleInstance) (string, string, string)
	SSMLDemo(l l10n.LocaleInstance) (string, string, string)
	SaySomething(l l10n.LocaleInstance) (string, string, string)
	Demo(l l10n.LocaleInstance) (string, string, string)
}

func NewMux(app Application) alexa.Handler {
	mux := alexa.NewServerMux()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)

	mux.HandleIntent(alexa.HelpIntent, handleHelp(app))
	mux.HandleIntent(alexa.CancelIntent, handleStop(app))
	mux.HandleIntent(alexa.StopIntent, handleStop(app))

	mux.HandleIntent(SSMLDemoIntent, handleSSMLResponse(app))
	mux.HandleIntent(SaySomethingIntent, handleSaySomethingResponse(app))
	mux.HandleIntent(DemoIntent, handleDemo(app))

	return mux
}

func handleCanFulfillIntent(b *alexa.ResponseBuilder, r *alexa.Request) {
	intent := r.Intent.Name
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

func handleLaunch(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			return
		}
		title, text := app.Launch(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text)
	})
}

func handleHelp(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			// TODO: maybe say something here
			return
		}
		title, text, _ := app.Help(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

func handleStop(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			// TODO: maybe say something here
			return
		}

		title, text, _ := app.Stop(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

func handleSSMLResponse(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			// TODO: maybe say something here
			return
		}

		title, text, ssmlText := app.SSMLDemo(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

func handleSaySomethingResponse(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			return
		}

		title, text, ssmlText := app.SaySomething(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

func handleDemo(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			return
		}

		title, text, ssmlText := app.Demo(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

func handleLocaleErrors(b *alexa.ResponseBuilder, errs []error) {
	b.WithSimpleCard("error", fmt.Sprintf("last error: %s", errs[len(errs)-1].Error())).
		WithSpeech(l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "An error occurred")))
}
