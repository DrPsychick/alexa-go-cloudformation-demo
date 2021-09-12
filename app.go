// Package alfalfa contains base elements of the skill project (app, skill).
package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

// Config defines additional data that can be provided and used in requests.
type Config struct {
	User string
}

// type AppResponseFunc func(locale l10n.LocaleInstance, opts ...ResponseFunc) (Response, error)

// ResponseFunc defines the function that can optionally be passed to responses.
type ResponseFunc func(cfg *Config)

// WithUser returns a ResponseFunc that sets the user.
func WithUser(user string) ResponseFunc {
	return func(cfg *Config) {
		cfg.User = user
	}
}

// Application defines the base application.
type Application struct {
	logger  log.Logger
	statter stats.Statter
}

// NewApplication returns an Application with the logger and statter.
func NewApplication(l log.Logger, s stats.Statter) *Application {
	return &Application{
		logger:  l,
		statter: s,
	}
}

// Launch is the response to the launch request.
func (a *Application) Launch(l l10n.LocaleInstance) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(l10n.KeyLaunchTitle),
		Text:   l.GetAny(l10n.KeyLaunchText),
		Speech: l.GetAny(l10n.KeyLaunchSSML),
		End:    false,
	}, nil
}

// Help is the response to a help request.
func (a *Application) Help(l l10n.LocaleInstance) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(l10n.KeyHelpTitle),
		Text:   l.GetAny(l10n.KeyHelpText),
		Speech: l.GetAny(l10n.KeyHelpSSML),
		End:    false,
	}, nil
}

// Stop is the response to stop the skill.
func (a *Application) Stop(l l10n.LocaleInstance) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(l10n.KeyStopTitle),
		Text:   l.GetAny(l10n.KeyStopText),
		Speech: l.GetAny(l10n.KeyStopSSML),
		End:    true,
	}, nil
}

// SSMLDemo is the intent to demonstrate SSML output with Alexa.
func (a *Application) SSMLDemo(l l10n.LocaleInstance) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(loca.LaunchTitle),
		Text:   l.GetAny(loca.LaunchText),
		Speech: l.GetAny(loca.LaunchSSML),
		End:    true,
	}, nil
}

// Demo is a simple demo response.
func (a *Application) Demo(l l10n.LocaleInstance) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.Get(loca.DemoIntentTitle),
		Text:   l.GetAny(loca.DemoIntentText),
		Speech: l.GetAny(loca.DemoIntentSSML),
		End:    true,
	}, nil
}

// SaySomething handles simple title + text response.
func (a *Application) SaySomething(loc l10n.LocaleInstance, opts ...ResponseFunc) (alexa.Response, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	var tit, msg, msgSSML string
	if cfg.User != "" {
		// personalized response
		tit = loc.GetAny(loca.SaySomethingUserTitle, cfg.User)
		msg = loc.GetAny(loca.SaySomethingUserText, cfg.User)
		msgSSML = loc.GetAny(loca.SaySomethingUserSSML, cfg.User)
	} else {
		tit = loc.GetAny(loca.SaySomethingTitle)
		msg = loc.GetAny(loca.SaySomethingText)
		msgSSML = loc.GetAny(loca.SaySomethingSSML)
	}

	if msg == "" {
		if cfg.User != "" {
			return alexa.Response{}, &l10n.NoTranslationError{Locale: loc.GetName(), Key: loca.SaySomethingUserText}
		}
		return alexa.Response{}, &l10n.NoTranslationError{Locale: loc.GetName(), Key: loca.SaySomethingText}
	}

	return alexa.Response{
		Title:  tit,
		Text:   msg,
		Speech: msgSSML,
		End:    true,
	}, nil
}

// AWSStatus responds with messages containing 2 slots.
func (a *Application) AWSStatus(loc l10n.LocaleInstance, area, region string) (alexa.Response, error) {
	title := loc.GetAny(loca.AWSStatusTitle)
	msg := loc.GetAny(loca.AWSStatusText, area, region)
	msgSSML := loc.GetAny(loca.AWSStatusSSML, area, region)

	return alexa.Response{
		Title:  title,
		Text:   msg,
		Speech: msgSSML,
		Image:  "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_%s.png", //nolint:lll
		End:    true,
	}, nil
}

// AWSStatusRegionElicit will ask for the Region value.
func (a *Application) AWSStatusRegionElicit(l l10n.LocaleInstance, region string) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(loca.AWSStatusTitle),
		Text:   l.GetAny(loca.AWSStatusRegionElicitText),
		Speech: l.GetAny(loca.AWSStatusRegionElicitSSML),
		End:    false,
	}, nil
}

// AWSStatusAreaElicit will ask for the Area value.
func (a *Application) AWSStatusAreaElicit(l l10n.LocaleInstance, area string) (alexa.Response, error) {
	return alexa.Response{
		Title:  l.GetAny(loca.AWSStatusTitle),
		Text:   l.GetAny(loca.AWSStatusAreaElicitText),
		Speech: l.GetAny(loca.AWSStatusAreaElicitSSML),
		End:    false,
	}, nil
}

// Logger returns the application logger.
func (a *Application) Logger() log.Logger {
	return a.logger
}

// Statter returns the application statter.
func (a *Application) Statter() stats.Statter {
	return a.statter
}
