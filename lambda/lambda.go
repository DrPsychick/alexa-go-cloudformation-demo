package lambda

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
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
	AWSStatusRegionElicit(l l10n.LocaleInstance, r string) (string, string, string)
	AWSStatusAreaElicit(l l10n.LocaleInstance, r string) (string, string, string)
	SaySomething(l l10n.LocaleInstance, opts ...alfalfa.ResponseFunc) (alfalfa.ApplicationResponse, error)
	AWSStatus(l l10n.LocaleInstance, area string, region string) (alfalfa.ApplicationResponse, error)
}

// NewMux returns a new handler for defined intents
func NewMux(app Application, sb *gen.SkillBuilder) alexa.Handler {
	mux := alexa.NewServerMux(app.Logger())
	sb.WithModel()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
	mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)
	mux.HandleRequestTypeFunc(alexa.TypeSessionEndedRequest, handleEnd(app))

	// new approach:
	mux.HandleIntent(alexa.HelpIntent, handleHelp(app, sb))
	mux.HandleIntent(alexa.CancelIntent, handleStop(app, sb))
	mux.HandleIntent(alexa.StopIntent, handleStop(app, sb))
	mux.HandleIntent(loca.DemoIntent, handleSSMLResponse(app, sb))
	mux.HandleIntent(loca.SaySomething, handleSaySomethingResponse(app, sb))
	mux.HandleIntent(loca.AWSStatus, handleAWSStatus(app, sb)) //, WithSlot(loca.TypeArea))

	return mux
}

func handleCanFulfillIntent(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
	intent := r.Request.Intent.Name
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
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
			return
		}
		title, text := app.Launch(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(false)
	})
}

func handleEnd(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
			return
		}

		title, text, _ := app.Stop(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

// handleHelp calls the app help method, it does not close the session
func handleHelp(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(alexa.HelpIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
			return
		}
		title, text, _ := app.Help(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(false)
	})
}

func handleStop(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(alexa.StopIntent)
	sb.Model().WithIntent(alexa.CancelIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
			return
		}
		title, text, _ := app.Stop(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

func handleSSMLResponse(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(loca.DemoIntent)
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
			return
		}

		title, text, ssmlText := app.SSMLDemo(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc.GetErrors())
			loc.ResetErrors()
			return
		}

		b.WithSpeech(ssmlText).
			WithSimpleCard(title, text)
	})
}

// simple: one specific function per intent
func handleSaySomethingResponse(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(loca.SaySomething)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			handleMissingLocale(b, r.Request.Locale)
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

// SlotValue always returns a string, it will be empty if the slot is not found
func SlotValue(r *alexa.RequestEnvelope, n string) string {
	s, ok := r.Request.Intent.Slots[n]
	if !ok {
		return ""
	}
	return s.Value
}

// SlotAuthorities always returns a PerAuthority slice
func SlotAuthorities(r *alexa.RequestEnvelope, n string) []*alexa.PerAuthority {
	s, ok := r.Request.Intent.Slots[n]
	if !ok {
		return []*alexa.PerAuthority{}
	}
	if s.Resolutions == nil || s.Resolutions.PerAuthority == nil {
		return []*alexa.PerAuthority{}
	}
	return s.Resolutions.PerAuthority
}

func SlotResolutionFirstValue(r *alexa.RequestEnvelope, n string) string {
	sa := SlotAuthorities(r, n)
	if len(sa) == 0 || len(sa[0].Values) == 0 || sa[0].Values[0] == nil || sa[0].Values[0].Value == nil {
		return ""
	}
	return sa[0].Values[0].Value.Name
}

// func SlotResolutionValues

func SlotMatch(r *alexa.RequestEnvelope, n string) bool {
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

func SlotNoMatch(r *alexa.RequestEnvelope, n string) bool {
	sa := SlotAuthorities(r, n)
	if len(sa) == 0 {
		return false
	}
	if sa[0].Status == nil {
		return false
	}
	return sa[0].Status.Code == alexa.ResolutionStatusNoMatch
}

func handleAWSStatus(app Application, sb *gen.SkillBuilder) alexa.Handler {
	// TODO: the mux should know about slots and "pass" it to the handler via request
	// register intent, slots, types with the model
	sb.Model().WithIntent(loca.AWSStatus)
	sb.Model().
		WithType(loca.TypeArea).
		WithType(loca.TypeRegion)

	sb.Model().Intent(loca.AWSStatus).
		WithDelegation(alexa.DelegationSkillResponse).
		WithSlot(loca.TypeAreaName, loca.TypeArea).
		WithSlot(loca.TypeRegionName, loca.TypeRegion)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		tags := []interface{}{"intent", loca.AWSStatus, "locale", r.Request.Locale}

		loc, err := loca.Registry.Resolve(r.Request.Locale)
		if err != nil {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			handleMissingLocale(b, r.Request.Locale)
			return
		}

		//// elicit the slot value through Alexa
		if !SlotMatch(r, loca.TypeAreaName) { // using 'not SlotMatch' because that includes a missing slot
			// failed validation or missing -> elicit - but need to provide prompt!
			title, text, ssml := app.AWSStatusAreaElicit(loc, SlotValue(r, loca.TypeAreaName))

			if len(loc.GetErrors()) > 0 {
				handleLocaleErrors(b, loc.GetErrors())
				loc.ResetErrors()
				return
			}

			b.AddDirective(&alexa.Directive{
				Type:         alexa.DirectiveTypeDialogElicitSlot,
				SlotToElicit: loca.TypeAreaName,
			}).
				WithSpeech(ssml).
				WithSimpleCard(title, text)
			return
		}
		area := SlotResolutionFirstValue(r, loca.TypeAreaName)

		//// if slot is empty and dialog still open, respond with Dialog:Delegate
		//if area == "" {
		//	if r.Request.DialogState == alexa.DialogStateCompleted {
		//		// should not happen (Alexa validation would have failed?)
		//		// how to respond if it does happen?
		//	} else {
		//		b.AddDirective(&alexa.Directive{
		//			Type: alexa.DirectiveTypeDialogDelegate,
		//			// UpdatedIntent:  only needed when changing intent
		//		})
		//	}
		//	return
		//}

		// elicit the slot value through Alexa
		if !SlotMatch(r, loca.TypeRegionName) { // using 'not SlotMatch' because that includes a missing slot
			// failed validation or missing -> elicit - but need to provide prompt!
			title, text, ssml := app.AWSStatusRegionElicit(loc, SlotValue(r, loca.TypeRegionName))

			if len(loc.GetErrors()) > 0 {
				handleLocaleErrors(b, loc.GetErrors())
				loc.ResetErrors()
				return
			}

			b.AddDirective(&alexa.Directive{
				Type:         alexa.DirectiveTypeDialogElicitSlot,
				SlotToElicit: loca.TypeRegionName,
			}).
				WithSpeech(ssml).
				WithSimpleCard(title, text)
			return
		}

		// -> r.Intent.Slots["Region"].Resolutions.PerAuthority[0].Values[0].Value.Name
		region := SlotResolutionFirstValue(r, loca.TypeRegionName)

		//// if slot is empty and dialog still open, respond with Dialog:Delegate
		//if region == "" {
		//	if r.Request.DialogState == alexa.DialogStateCompleted {
		//		// should not happen (Alexa validation would have failed?)
		//		// how to respond if it does happen?
		//	} else {
		//		b.AddDirective(&alexa.Directive{
		//			Type: alexa.DirectiveTypeDialogDelegate,
		//			// UpdatedIntent:  only needed when changing intent
		//		})
		//	}
		//	return
		//}

		resp, err := app.AWSStatus(loc, area, region)
		if err != nil {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			switch err {
			default:
				fallthrough
			case alfalfa.ErrorNoTranslation:
				if len(loc.GetErrors()) > 0 {
					handleLocaleErrors(b, loc.GetErrors())
					loc.ResetErrors()
					return
				}
				resp = alfalfa.ApplicationResponse{}
				resp.Title = loc.GetAny(l10n.KeyErrorNoTranslationTitle)
				resp.Text = loc.GetAny(l10n.KeyErrorNoTranslationText, loca.AWSStatus)
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
func handleError(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope, err error) {
	loc := localeDefaults(r.Request.Locale)
	switch err {
	default:
		b.WithSimpleCard(loc.GetAny(l10n.KeyErrorTitle), loc.GetAny(l10n.KeyErrorText, err.Error())).
			WithShouldEndSession(true)
	}
}

func localeDefaults(locale string) l10n.LocaleInstance {
	loc, err := loca.Registry.Resolve(locale)
	if err != nil {
		loc = l10n.NewLocale(locale)
		loca.Registry.Register(loc)
	}
	if loc.Get(l10n.KeyErrorTitle) == "" {
		loc.Set(l10n.KeyErrorTitle, []string{"Error"})
	}
	if loc.Get(l10n.KeyErrorText) == "" {
		loc.Set(l10n.KeyErrorText, []string{"The app returned an error:\n%s"})
	}
	if loc.Get(l10n.KeyErrorMissingPlaceholder) == "" {
		loc.Set(l10n.KeyErrorMissingPlaceholder, []string{"the string is missing a placeholder %%s: '%s'"})
	}
	return loc
}

// handleMissingLocale makes Alexa respond with a "local not supported" error
func handleMissingLocale(b *alexa.ResponseBuilder, locale string) {
	b.WithSimpleCard("error", fmt.Sprintf("Locale '%s' is not supported!", locale)).
		WithSpeech(l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "Your language is not supported"))).
		WithShouldEndSession(true)
}

// handleLocaleErrors makes Alexa show the last error on the screen
func handleLocaleErrors(b *alexa.ResponseBuilder, errs []error) {
	b.WithSimpleCard("error", fmt.Sprintf("last error: %s", errs[len(errs)-1].Error())).
		WithSpeech(l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "An error occurred"))).
		WithShouldEndSession(true)
}
