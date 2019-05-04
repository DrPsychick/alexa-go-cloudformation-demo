package alfalfa

import (
	"errors"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

var (
	ErrorNoTranslation = errors.New("no '%s' translation for %s")
)

// Application defines the base application
type Application struct {
	logger  log.Logger
	statter stats.Statter
}

// NewApplication returns an Application with the logger and statter
func NewApplication(l log.Logger, s stats.Statter) *Application {
	return &Application{
		logger:  l,
		statter: s,
	}
}

func (a *Application) SaySomething2() AppResponseFunc {
	return AppResponseFunc(func(loc l10n.LocaleInstance, opts ...ResponseFunc) (ApplicationResponse, error) {
		// run all ResponseFuncs
		cfg := &Config{}
		for _, opt := range opts {
			opt(cfg)
		}

		tit := ""
		msg := ""
		msgSSML := ""
		if cfg.user != "" {
			// personalized response
			tit = loc.GetAny("SaySomething_UserTitle", cfg.user)
			msg = loc.GetAny("SaySomething_UserResponse", cfg.user)
			msgSSML = loc.GetAny("SaySomething_UserResponse"+l10n.KeyPostfixSSML, cfg.user)
		} else {
			tit = loc.GetAny(loca.SaySomethingTitle)
			msg = loc.GetAny(loca.SaySomethingText)
			msgSSML = loc.GetAny(loca.SaySomethingSSML)
		}

		if msg == "" {
			return ApplicationResponse{}, ErrorNoTranslation
		}

		return ApplicationResponse{
			Title:  tit,
			Text:   msg,
			Speech: msgSSML,
			End:    true,
		}, nil
	})
}

// Launch is the response to the launch request.
func (a *Application) Launch(l l10n.LocaleInstance) (string, string) {
	return l.Get(loca.LaunchTitle), l.GetAny(loca.LaunchText)
}

// Help is the response to a help request.
func (a *Application) Help() (string, string) {
	return "Help", "No help available!"
}

// Stop is the response to stop the skill.
func (a *Application) Stop(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.StopTitle), l.GetAny(loca.Stop), ""
}

// SimpleResponse handles simple title + text response.
func (a *Application) SaySomething(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.SaySomethingTitle), l.GetAny(loca.SaySomethingText), l.GetAny(loca.SaySomethingSSML)
}

// SSMLDemo is the intent to demonstrate SSML output with Alexa.
func (a *Application) SSMLDemo(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.LaunchTitle), l.GetAny(loca.LaunchText), l.GetAny(loca.LaunchSSML)
}

// Demo is a simple demo response.
func (a *Application) Demo(l l10n.LocaleInstance) (string, string, string) {
	return l.Get(loca.GenericTitle), l.GetAny(loca.DemoIntentText), l.GetAny(loca.DemoIntentSSML)
}

// Logger returns the application logger.
func (a *Application) Logger() log.Logger {
	return a.logger
}

// Statter returns the application statter.
func (a *Application) Statter() stats.Statter {
	return a.statter
}

type ApplicationResponse struct {
	Title  string
	Text   string
	Speech string
	Image  string
	End    bool
}
type Config struct {
	user string
}
type ResponseFunc func(cfg *Config)
type AppResponseFunc func(locale l10n.LocaleInstance, opts ...ResponseFunc) (ApplicationResponse, error)
