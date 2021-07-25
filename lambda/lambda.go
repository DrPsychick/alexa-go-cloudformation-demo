package lambda

import (
	"errors"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

const (
	// SSMLDemoIntent is the name of the SSL demo intent
	SSMLDemoIntent = "SSMLDemoIntent"
	// DemoIntent is the name of the demo intent
	DemoIntent = "DemoIntent"
)

var (
	// ErrorLocaleNotFound is the error text for missing locale
	ErrorLocaleNotFound = errors.New("locale not found")
)

// Application defines the interface used of the app
type Application interface {
	log.Loggable
	stats.Statable

	Launch(l l10n.LocaleInstance) (string, string)
	Help(l l10n.LocaleInstance) (string, string, string)
	Stop(l l10n.LocaleInstance) (string, string, string)
	SSMLDemo(l l10n.LocaleInstance) (string, string, string)
	Demo(l l10n.LocaleInstance) (string, string, string)
	SaySomething(l l10n.LocaleInstance, opts ...alfalfa.ResponseFunc) (alfalfa.ApplicationResponse, error)
	AWSStatus(l l10n.LocaleInstance, area string, region string) (alfalfa.ApplicationResponse, error)
}

// NewMux returns a new handler for defined intents
func NewMux(app Application, sb *gen.SkillBuilder) alexa.Handler {
	mux := alexa.NewServerMux()
	sb.WithModel()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)

	mux.HandleIntent(alexa.HelpIntent, handleHelp(app))
	sb.Model().WithIntent(alexa.HelpIntent)
	mux.HandleIntent(alexa.CancelIntent, handleStop(app))
	sb.Model().WithIntent(alexa.CancelIntent)
	mux.HandleIntent(alexa.StopIntent, handleStop(app))
	sb.Model().WithIntent(alexa.StopIntent)

	mux.HandleIntent(loca.DemoIntent, handleSSMLResponse(app))
	sb.Model().WithIntent(loca.DemoIntent)

	// new approach:
	mux.HandleIntent(loca.SaySomething, handleSaySomethingResponse(app, sb))
	mux.HandleIntent(loca.AWSStatus, handleAWSStatus(app, sb)) //, WithSlot(loca.TypeArea))

	return mux
}

func handleCanFulfillIntent(b *alexa.ResponseBuilder, r *alexa.Request) {
	intent := r.Intent.Name
	if intent == loca.DemoIntent || intent == loca.SaySomething || intent == loca.AWSStatus {
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
		l, err := loca.Registry.Resolve(r.Locale)
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
			WithSimpleCard(title, text)
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
			WithSimpleCard(title, text)
	})
}

func handleStop(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		l, err := loca.Registry.Resolve(r.Locale)
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
		l, err := loca.Registry.Resolve(r.Locale)
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

// simple: one specific function per intent
func handleSaySomethingResponse(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(loca.SaySomething)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		loc, err := loca.Registry.Resolve(r.Locale)
		if err != nil {
			handleMissingLocale(b, r.Locale)
			return
		}

		resp, err := app.SaySomething(loc)
		if err != nil {
			switch err {
			default:
				fallthrough
			case alfalfa.ErrorNoTranslation:
				resp = alfalfa.ApplicationResponse{}
				resp.Title = loc.GetAny(l10n.KeyErrorNoTranslationTitle)
				resp.Text = loc.GetAny(l10n.KeyErrorNoTranslationText, loca.SaySomething)
				resp.Speech = loc.GetAny(l10n.KeyErrorNoTranslationSSML)
				resp.End = true
			}
		}
		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSimpleCard(resp.Title, resp.Text)

		if resp.Speech != "" {
			b.WithSpeech(resp.Speech)
		}

		if resp.End {
			b.WithShouldEndSession(true)
		}
	})
}

func handleAWSStatus(app Application, sb *gen.SkillBuilder) alexa.Handler {
	// TODO: the mux should know about slots and "pass" it to the handler via request
	// register intent, slots, types with the model
	sb.Model().WithIntent(loca.AWSStatus)
	sb.Model().
		WithType(loca.TypeArea).
		WithType(loca.TypeRegion)

	sb.Model().Intent(loca.AWSStatus).
		WithSlot(loca.TypeAreaName, loca.TypeArea).
		WithSlot(loca.TypeRegionName, loca.TypeRegion)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		tags := []interface{}{"intent", loca.AWSStatus, "locale", r.Locale}

		loc, err := loca.Registry.Resolve(r.Locale)
		if err != nil {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			handleMissingLocale(b, r.Locale)
			return
		}

		// require slot input
		area, ok := r.Intent.Slots[loca.TypeArea]
		if !ok {
			// reprompt area slot
			stats.Inc(app, "request.error", 1, 1.0, tags...)
			handleError(b, r, fmt.Errorf("area not defined"))
			return
		}
		ar := area.Value
		if ar == "" && area.SlotValue != nil && area.SlotValue.Resolutions != nil {
			if rpa := area.SlotValue.Resolutions.PerAuthority; rpa != nil && len(rpa) > 0 {
				if vs := rpa[0].Values; vs != nil && len(vs) > 0 {
					ar = vs[0].Value.Name
				}
			}
		}
		// if not provided, respond with Dialog:Delegate
		if ar == "" {
			b.AddDirective(&alexa.Directive{
				Type: alexa.DirectiveTypeDialogDelegate,
			})
			return
		}

		region, ok := r.Intent.Slots[loca.TypeRegion]
		if !ok {
			// reprompt region slot
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			handleError(b, r, fmt.Errorf("region not defined"))
			return
		}
		re := region.Value
		if re == "" && region.SlotValue != nil && region.SlotValue.Resolutions != nil {
			if rpa := region.SlotValue.Resolutions.PerAuthority; rpa != nil && len(rpa) > 0 {
				if vs := rpa[0].Values; vs != nil && len(vs) > 0 {
					re = vs[0].Value.Name
				}
			}
		}
		// if not provided, respond with Dialog:Delegate
		if re == "" {
			b.AddDirective(&alexa.Directive{
				Type: alexa.DirectiveTypeDialogDelegate,
				//UpdatedIntent: &alexa.Intent{ // only needed when changing intent
				//	Name: loca.AWSStatus,
				//	ConfirmationStatus: "NONE",
				//	Slots: map[string]Slot,
				//},
			})
			return
		}

		resp, err := app.AWSStatus(loc, ar, re)
		if err != nil {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			switch err {
			default:
				fallthrough
			case alfalfa.ErrorNoTranslation:
				resp = alfalfa.ApplicationResponse{}
				resp.Title = loc.GetAny(l10n.KeyErrorNoTranslationTitle)
				resp.Text = loc.GetAny(l10n.KeyErrorNoTranslationText, loca.SaySomething)
				resp.Speech = loc.GetAny(l10n.KeyErrorNoTranslationSSML)
				resp.End = true
			}
		}

		b.WithSimpleCard(resp.Title, resp.Text)
		if resp.Image != "" {
			b.WithStandardCard(resp.Title, resp.Text, &alexa.Image{
				SmallImageURL: fmt.Sprintf(resp.Image, "small"),
				LargeImageURL: fmt.Sprintf(resp.Image, "large"),
			})
		}

		if resp.Speech != "" {
			b.WithSpeech(resp.Speech)
		}

		if resp.End {
			b.WithShouldEndSession(true)
		}
	})
}

// TODO: handle errors individually to be of more use to the user
func handleError(b *alexa.ResponseBuilder, r *alexa.Request, err error) {
	l := localeDefaults(r.Locale)
	switch err {
	default:
		b.WithSimpleCard(l.GetAny(l10n.KeyErrorTitle), l.GetAny(l10n.KeyErrorText, err.Error())).
			WithShouldEndSession(true)
	}
}

func localeDefaults(locale string) l10n.LocaleInstance {
	l, err := loca.Registry.Resolve(locale)
	if err != nil {
		l = l10n.NewLocale(locale)
		loca.Registry.Register(l)
	}
	if l.Get(l10n.KeyErrorTitle) == "" {
		l.Set(l10n.KeyErrorTitle, []string{"Error"})
	}
	if l.Get(l10n.KeyErrorText) == "" {
		l.Set(l10n.KeyErrorText, []string{"The app returned an error:\n%s"})
	}
	if l.Get(l10n.KeyErrorMissingPlaceholder) == "" {
		l.Set(l10n.KeyErrorMissingPlaceholder, []string{"the string is missing a placeholder %%s: '%s'"})
	}
	return l
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
