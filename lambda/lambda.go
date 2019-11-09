package lambda

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
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
	AWSStatus(l l10n.LocaleInstance, r string) (string, string, string)
}

func NewMux(app Application) alexa.Handler {
	mux := alexa.NewServerMux()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)
	mux.HandleRequestTypeFunc(alexa.TypeSessionEndedRequest, handleEnd(app))

	mux.HandleIntent(alexa.HelpIntent, handleHelp(app))
	mux.HandleIntent(alexa.CancelIntent, handleStop(app))
	mux.HandleIntent(alexa.StopIntent, handleStop(app))

	mux.HandleIntent(SSMLDemoIntent, handleSSMLResponse(app))
	mux.HandleIntent(SaySomethingIntent, handleSaySomethingResponse(app))
	mux.HandleIntent(DemoIntent, handleDemo(app))

	mux.HandleIntent(loca.AWSStatus, handleAWSStatus(app))

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
			handleMissingLocale(b, r.Locale)
			return
		}
		title, text := app.Launch(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(false)
	})
}

func handleEnd(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
	})
}

// handleHelp calls the app help method, it does not close the session
func handleHelp(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			handleMissingLocale(b, r.Locale)
			return
		}
		title, text, _ := app.Help(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(false)
	})
}

func handleStop(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			handleMissingLocale(b, r.Locale)
			return
		}

		title, text, _ := app.Stop(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
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
			handleMissingLocale(b, r.Locale)
			return
		}

		title, text, ssmlText := app.SSMLDemo(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
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
			handleMissingLocale(b, r.Locale)
			return
		}

		title, text, ssmlText := app.SaySomething(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
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
			handleMissingLocale(b, r.Locale)
			return
		}

		title, text, ssmlText := app.Demo(l)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

// SlotAuthorities always returns a PerAuthority slice
func SlotAuthorities(r *alexa.Request, n string) []*alexa.PerAuthority {
	s, ok := r.Intent.Slots[n]
	if !ok {
		return []*alexa.PerAuthority{}
	}
	if s.Resolutions == nil || s.Resolutions.PerAuthority == nil {
		return []*alexa.PerAuthority{}
	}
	return s.Resolutions.PerAuthority
}

func SlotMatch(r *alexa.Request, n string) bool {
	// TODO: what about multiple Authorities?
	sa := SlotAuthorities(r, n)
	if len(sa) == 0 {
		return false
	}
	if sa[0].Status == nil {
		return false
	}
	return sa[0].Status.Code == alexa.ResolutionStatusMatch
}

func SlotNoMatch(r *alexa.Request, n string) bool {
	sa := SlotAuthorities(r, n)
	if len(sa) == 0 {
		return false
	}
	if sa[0].Status == nil {
		return false
	}
	return sa[0].Status.Code == alexa.ResolutionStatusNoMatch
}

func handleAWSStatus(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := l10n.Resolve(r.Locale)
		if err != nil {
			handleMissingLocale(b, r.Locale)
			return
		}

		// TODO: make request content directly accessible
		// SlotMatch(r, "Region") = true/false
		// SlotValue(r, "Region")
		if SlotNoMatch(r, "Region") {
			// failed validation -> elicit (but need to provide prompt!)
		}

		region := "unknown"
		// -> r.Intent.Slots["Region"].Resolutions.PerAuthority[0].Values[0].Value.Name
		if rs, ok := r.Intent.Slots["Region"]; ok {
			if rs.Resolutions != nil && rs.Resolutions.PerAuthority != nil {
				if rsa := rs.Resolutions.PerAuthority; len(rsa) > 0 {
					if rsav := rsa[0].Values; len(rsav) > 0 {
						region = rsav[0].Value.Name
					}
				}
			}
		}
		// TODO: if unknown, respond with Dialog:Delegate
		if region == "unknown" {
			if r.DialogState == alexa.DialogStateCompleted {
				// should not happen (Alexa validation would have failed?)
			} else {
				b.AddDirective(&alexa.Directive{
					Type: alexa.DirectiveTypeDialogDelegate,
					// UpdatedIntent:  only needed when changing intent
				})
			}
			return
		}

		title, text, ssmlText := app.AWSStatus(l, region)

		if len(l.GetErrors()) > 0 {
			handleLocaleErrors(b, l.GetErrors())
			l.ResetErrors()
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

// handleMissingLocale makes Alexa respond with a "local not supported" error
func handleMissingLocale(b *alexa.ResponseBuilder, locale string) {
	b.WithSimpleCard("error", fmt.Sprintf("Locale '%s' is not supported!", locale)).
		WithSpeech(l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "Your language is not supported")))
}

// handleLocaleErrors makes Alexa show the last error on the screen
func handleLocaleErrors(b *alexa.ResponseBuilder, errs []error) {
	b.WithSimpleCard("error", fmt.Sprintf("last error: %s", errs[len(errs)-1].Error())).
		WithSpeech(l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "An error occurred")))
}
