// Package lambda defines intents, handles requests and calls Application functions accordingly.
package lambda

import (
	"errors"

	alfalfa "github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/go-alexa-lambda"
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/skill"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Application defines the interface used of the app.
type Application interface {
	log.Loggable
	stats.Statable

	Launch(l l10n.LocaleInstance) (alexa.Response, error)
	Help(l l10n.LocaleInstance) (alexa.Response, error)
	Stop(l l10n.LocaleInstance) (alexa.Response, error)
	SSMLDemo(l l10n.LocaleInstance) (alexa.Response, error)
	Demo(l l10n.LocaleInstance) (alexa.Response, error)
	AWSStatusRegionElicit(l l10n.LocaleInstance, r string) (alexa.Response, error)
	AWSStatusAreaElicit(l l10n.LocaleInstance, r string) (alexa.Response, error)
	SaySomething(l l10n.LocaleInstance, opts ...alfalfa.ResponseFunc) (alexa.Response, error)
	AWSStatus(l l10n.LocaleInstance, area, region string) (alexa.Response, error)
}

// NewMux returns a new handler for defined intents.
func NewMux(app Application, sb *skill.SkillBuilder) alexa.Handler {
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

func launch(app Application, b *alexa.ResponseBuilder, loc l10n.LocaleInstance) error {
	resp, err := app.Launch(loc)
	if err != nil {
		return err
	}

	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleLaunch(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := launch(app, b, loc); err != nil {
			log.Error(app, "could not handle Stop: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func handleEnd(app Application) alexa.HandlerFunc {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := stop(app, b, loc); err != nil {
			log.Error(app, "could not handle Stop: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func help(app Application, b *alexa.ResponseBuilder, loc l10n.LocaleInstance) error {
	resp, err := app.Help(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

// handleHelp calls the app help method, it does not close the session.
func handleHelp(app Application, sb *skill.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(alexa.HelpIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := help(app, b, loc); err != nil {
			log.Error(app, "could not handle Stop: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func stop(app Application, b *alexa.ResponseBuilder, loc l10n.LocaleInstance) error {
	resp, err := app.Stop(loc)
	if err != nil {
		return err
	}

	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleStop(app Application, sb *skill.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(alexa.StopIntent)
	sb.Model().WithIntent(alexa.CancelIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := stop(app, b, loc); err != nil {
			log.Error(app, "could not handle Stop: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

func ssmlResponse(app Application, b *alexa.ResponseBuilder, loc l10n.LocaleInstance) error {
	resp, err := app.SaySomething(loc)
	if err != nil {
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		return err
	}

	b.With(resp)
	return nil
}

func handleSSMLResponse(app Application, sb *skill.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(loca.DemoIntent)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		// get loc, call app func, check loca errors -> always return err
		if err := ssmlResponse(app, b, loc); err != nil {
			// isResponse
			log.Error(app, "could not handle SSMLResponse: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}
	})
}

// simple: one specific function per intent
func handleSaySomethingResponse(app Application, sb *skill.SkillBuilder) alexa.Handler {
	sb.Model().WithIntent(loca.SaySomething)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		var resp alexa.Response
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		responseFuncs := []alfalfa.ResponseFunc{}
		per, err := r.ContextPerson()
		if err == nil {
			responseFuncs = append(responseFuncs, alfalfa.WithUser(per.PersonID))
		}
		resp, err = app.SaySomething(loc, responseFuncs...)
		if res := alexa.HandleError(b, loc, err); res {
			return
		}

		b.With(resp)
	})
}

func awsStatus(app Application, b *alexa.ResponseBuilder, loc l10n.LocaleInstance, r *alexa.RequestEnvelope) error { //nolint:funlen,gocognit,lll,cyclop
	tags := []string{"intent", loca.AWSStatus, "locale", r.RequestLocale()}
	var resp alexa.Response

	// TODO: if err := validateSlot(loca.TypeAreaSlot); err != nil {}
	slotArea, err := r.Slot(loca.TypeAreaName)
	_, err2 := slotArea.FirstAuthorityWithMatch()
	// elicit the slot value through Alexa
	if err != nil || err2 != nil { //nolint:nestif
		// failed validation or missing -> elicit - but need to provide prompt!
		resp, err = app.AWSStatusAreaElicit(loc, r.SlotValue(loca.TypeAreaName))
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return nil
			}
			return err
		}
		if err := alexa.CheckForLocaleError(loc); err != nil {
			if alexa.HandleError(b, loc, err) {
				return nil
			}
			return err
		}
		b.With(resp)
		return nil
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
	if err != nil || err2 != nil { //nolint:nestif
		// failed validation or missing -> elicit - but need to provide prompt!
		resp, err = app.AWSStatusRegionElicit(loc, r.SlotValue(loca.TypeRegionName))
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return nil
			}
			return err
		}
		if err := alexa.CheckForLocaleError(loc); err != nil {
			for _, e := range loc.GetErrors() {
				log.Error(app, e.Error())
			}
			if alexa.HandleError(b, loc, err) {
				return nil
			}
			return err
		}
		b.With(resp)
		return nil
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

	resp, err = app.AWSStatus(loc, area, region)
	if err != nil {
		stats.Inc(app, "handleAWSStatus.error", 1, 1.0, tags...)
		if alexa.HandleError(b, loc, err) {
			return nil
		}
		return err
	}
	if err := alexa.CheckForLocaleError(loc); err != nil {
		if alexa.HandleError(b, loc, err) {
			return nil
		}
		return err
	}

	b.With(resp)
	return nil
}

func handleAWSStatus(app Application, sb *skill.SkillBuilder) alexa.Handler {
	// TODO: the mux should know about slots and "pass" it to the handler via request
	// register intent, slots, types with the model
	sb.Model().WithIntent(loca.AWSStatus)
	sb.Model().
		WithType(loca.TypeArea).
		WithType(loca.TypeRegion)

	sb.Model().Intent(loca.AWSStatus).
		WithDelegation(skill.DelegationSkillResponse).
		WithSlot(loca.TypeAreaName, loca.TypeArea).
		WithSlot(loca.TypeRegionName, loca.TypeRegion)

	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
		// var resp alexa.Response
		var loc l10n.LocaleInstance
		loc, err := loca.Registry.Resolve(r.RequestLocale())
		if err != nil {
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		if err := awsStatus(app, b, loc, r); err != nil {
			// isResponse
			log.Error(app, "could not handle AWSStatus: "+err.Error())
			if alexa.HandleError(b, loc, err) {
				return
			}
			alexa.HandleError(b, loc, &DefaultError{loc})
			return
		}

		monitorLocaleErrors(app, loc)
	})
}

// monitorLocaleErrors logs and stats every locale error.
func monitorLocaleErrors(app Application, loc l10n.LocaleInstance) {
	if len(loc.GetErrors()) > 0 {
		for _, err := range loc.GetErrors() {
			log.Error(app, err.Error())
			var locaErr l10n.LocaleError
			var tags []string
			if errors.As(err, &locaErr) {
				tags = []string{"locale", locaErr.GetLocale()}
			}
			stats.Inc(app, "locale_error", 1, 1, tags...)
		}
	}
}

// DefaultError is a generic error.
type DefaultError struct {
	Locale l10n.LocaleInstance
}

// Error returns a string.
func (m DefaultError) Error() string {
	return "an error occurred"
}

// Response returns a default error response.
func (m DefaultError) Response(loc l10n.LocaleInstance) alexa.Response {
	return alexa.Response{
		Title:  loc.GetAny(l10n.KeyErrorTitle),
		Text:   loc.GetAny(l10n.KeyErrorText),
		Speech: loc.GetAny(l10n.KeyErrorSSML),
		End:    true,
	}
}
