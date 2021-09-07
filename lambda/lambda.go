// Package lambda defines intents, handles requests and calls Application functions accordingly.
package lambda

import (
	"errors"
	"fmt"

	alfalfa "github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Application defines the interface used of the app.
type Application interface {
	log.Loggable
	stats.Statable

	Launch(l l10n.LocaleInstance) (string, string)
	Help(l l10n.LocaleInstance) (string, string, string)
	Stop(l l10n.LocaleInstance) (string, string, string)
	SSMLDemo(l l10n.LocaleInstance) (string, string, string)
	Demo(l l10n.LocaleInstance) (string, string, string)
	AWSStatusRegionElicit(l l10n.LocaleInstance, r string) (alfalfa.ApplicationResponse, error)
	AWSStatusAreaElicit(l l10n.LocaleInstance, r string) (alfalfa.ApplicationResponse, error)
	SaySomething(l l10n.LocaleInstance, opts ...alfalfa.ResponseFunc) (alfalfa.ApplicationResponse, error)
	AWSStatus(l l10n.LocaleInstance, area, region string) (alfalfa.ApplicationResponse, error)
}

// NewMux returns a new handler for defined intents.
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
	mux.HandleIntent(loca.AWSStatus, handleAWSStatus(app, sb))

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
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}
		title, text := app.Launch(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(false)
	})
}

func handleEnd(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}

		title, text, _ := app.Stop(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
			return
		}

		b.WithSpeech(text).
			WithSimpleCard(title, text).
			WithShouldEndSession(true)
	})
}

// handleHelp calls the app help method, it does not close the session.
func handleHelp(app Application, sb *gen.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(alexa.HelpIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}
		title, text, _ := app.Help(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
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
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}
		title, text, _ := app.Stop(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
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
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}

		title, text, ssmlText := app.SSMLDemo(loc)

		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
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
		loc := locale(b, r.RequestLocale())
		if loc == nil {
			return
		}

		responseFuncs := []alfalfa.ResponseFunc{}
		if r.Context.System != nil && r.Context.System.Person != nil {
			responseFuncs = append(responseFuncs, alfalfa.WithUser(r.Context.System.Person.PersonID))
		}
		resp, err := app.SaySomething(loc, responseFuncs...)
		if res := handleError(b, r, err); res {
			return
		}

		response(b, resp)
	})
}

func handleAWSStatus(app Application, sb *gen.SkillBuilder) alexa.Handler { //nolint:funlen
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
		tags := []string{"intent", loca.AWSStatus, "locale", r.RequestLocale()}

		loc := locale(b, r.RequestLocale())
		if loc == nil {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			return
		}

		slotArea, err := r.Slot(loca.TypeAreaName)
		_, err2 := slotArea.FirstAuthorityWithMatch()
		// elicit the slot value through Alexa
		if err != nil || err2 != nil {
			// failed validation or missing -> elicit - but need to provide prompt!
			resp, err := app.AWSStatusAreaElicit(loc, r.SlotValue(loca.TypeAreaName))
			if res := handleError(b, r, err); res {
				return
			}

			response(b, resp)
			return
		}
		area := slotArea.Value

		// if slot is empty and dialog still open, respond with Dialog:Delegate
		// if area == "" {
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
		// }

		slotRegion, err := r.Slot(loca.TypeRegionName)
		_, err2 = slotRegion.FirstAuthorityWithMatch()
		// elicit the slot value through Alexa
		if err != nil || err2 != nil {
			// failed validation or missing -> elicit - but need to provide prompt!
			resp, err := app.AWSStatusRegionElicit(loc, r.SlotValue(loca.TypeRegionName))
			if res := handleError(b, r, err); res {
				return
			}
			response(b, resp)
			return
		}
		region := slotRegion.Value

		// if slot is empty and dialog still open, respond with Dialog:Delegate
		// if region == "" {
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
		// }

		resp, err := app.AWSStatus(loc, area, region)
		if res := handleError(b, r, err); res {
			stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
			return
		}

		response(b, resp)
	})
}

// TODO: allow registering error handlers which are then all checked
// mux.registerError(NotFoundError) {
// error.As(err, nfe) -> if nfe.HandleError(err) -> return
// func HandleError() (ApplicationResponse)
// TODO: generic ApplicationResponse!
// you could register objects with the mux that handle an error
// you would then call them one by one and if they return true, you stop handling
// this would be the simplest way.
func handleError(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope, err error) bool { //nolint:funlen
	var resp alfalfa.ApplicationResponse
	// locale error is already handled (err will not be nil)
	loc, _ := loca.Registry.Resolve(r.RequestLocale())
	if loc == nil {
		loc = loca.Registry.GetDefault()
	}

	if err != nil { //nolint:nestif
		// ignore previous locale errors as we're handling an error
		loc.ResetErrors()

		if errors.Is(err, alexa.ErrUnknown) {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorUnknownTitle),
				Text:   loc.GetAny(l10n.KeyErrorUnknownText),
				Speech: loc.GetAny(l10n.KeyErrorUnknownSSML),
				End:    true,
			}
		}
		var notFoundError alexa.NotFoundError
		if errors.As(err, &notFoundError) {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorNotFoundTitle),
				Text:   loc.GetAny(l10n.KeyErrorNotFoundText, notFoundError.Error()),
				Speech: loc.GetAny(l10n.KeyErrorNotFoundSSML),
				End:    true,
			}
		}
		var localeNotFoundError l10n.LocaleNotFoundError
		if errors.As(err, &localeNotFoundError) {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorLocaleNotFoundTitle),
				Text:   loc.GetAny(l10n.KeyErrorLocaleNotFoundText, localeNotFoundError.Locale),
				Speech: loc.GetAny(l10n.KeyErrorLocaleNotFoundSSML, localeNotFoundError.Locale),
				End:    true,
			}
		}
		var noTranslationError l10n.NoTranslationError
		if errors.As(err, &noTranslationError) {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorNoTranslationTitle),
				Text:   loc.GetAny(l10n.KeyErrorNoTranslationText, noTranslationError.Key),
				Speech: loc.GetAny(l10n.KeyErrorNoTranslationSSML, noTranslationError.Key),
				End:    true,
			}
		}
		var placeholderError l10n.MissingPlaceholderError
		if errors.As(err, &placeholderError) {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorMissingPlaceholderTitle),
				Text:   loc.GetAny(l10n.KeyErrorMissingPlaceholderText, placeholderError.Key),
				Speech: loc.GetAny(l10n.KeyErrorMissingPlaceholderSSML, placeholderError.Key),
				End:    true,
			}
		}

		// errors in error locales
		if len(loc.GetErrors()) > 0 {
			handleLocaleErrors(b, loc)
			return true
		}

		// fallback
		if resp.Title == "" {
			resp = alfalfa.ApplicationResponse{
				Title:  loc.GetAny(l10n.KeyErrorTitle),
				Text:   loc.GetAny(l10n.KeyErrorText, err.Error()),
				Speech: loc.GetAny(l10n.KeyErrorSSML),
				End:    true,
			}
		}

		response(b, resp)
		return true
	}

	// handle locale errors
	if len(loc.GetErrors()) > 0 {
		handleLocaleErrors(b, loc)
		return true
	}
	return false
}

func locale(b *alexa.ResponseBuilder, l string) l10n.LocaleInstance {
	loc, err := loca.Registry.Resolve(l)
	if err != nil {
		handleError(b, &alexa.RequestEnvelope{Request: &alexa.Request{Locale: (alexa.RequestLocale(l))}}, err)
		return nil
	}
	return loc
}

// TODO: needs to go into library.
func response(b *alexa.ResponseBuilder, resp alfalfa.ApplicationResponse) {
	if resp.Image != "" {
		b.WithStandardCard(resp.Title, resp.Text, &alexa.Image{
			SmallImageURL: fmt.Sprintf(resp.Image, "small"),
			LargeImageURL: fmt.Sprintf(resp.Image, "large"),
		})
	}
	b.WithSimpleCard(resp.Title, resp.Text)
	if resp.Speech != "" {
		b.WithSpeech(resp.Speech)
	}
	b.WithShouldEndSession(resp.End)
}

// handleLocaleErrors makes Alexa show the last error on the screen.
func handleLocaleErrors(b *alexa.ResponseBuilder, loc l10n.LocaleInstance) {
	errs := loc.GetErrors()
	loc.ResetErrors()
	resp := alfalfa.ApplicationResponse{
		Title:  loc.GetAny(l10n.KeyErrorTitle),
		Text:   loc.GetAny(l10n.KeyErrorText, errs[len(errs)-1].Error()),
		Speech: loc.GetAny(l10n.KeyErrorSSML),
		End:    true,
	}
	// fallback: even basic loca is missing
	if len(loc.GetErrors()) > 0 {
		resp = alfalfa.ApplicationResponse{
			Title: "error",
			Text:  fmt.Sprintf("An error occurred: %s", errs[len(errs)-1].Error()),
			End:   true,
		}
	}
	response(b, resp)
}
