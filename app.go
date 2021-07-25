package alfalfa

import (
	"errors"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
)

var (
	// ErrorNoTranslation is the text for missing translations
	ErrorNoTranslation = errors.New("translation missing")
)

// Config defines additional data that can be provided and used in requests
type Config struct {
	user string
}

//type AppResponseFunc func(locale l10n.LocaleInstance, opts ...ResponseFunc) (ApplicationResponse, error)

// ResponseFunc defines the function that can optionally be passed to responses
type ResponseFunc func(cfg *Config)

// WithUser returns a ResponseFunc that sets the user
func WithUser(user string) ResponseFunc {
	return func(cfg *Config) {
		cfg.user = user
	}
}

// ApplicationResponse defines the reponse returned to lambda
type ApplicationResponse struct {
	Title  string
	Text   string
	Speech string
	Image  string
	End    bool
}

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

// Launch is the response to the launch request.
func (a *Application) Launch(l l10n.LocaleInstance) (string, string) {
	return l.Get(loca.LaunchTitle), l.GetAny(loca.LaunchText)
}

// Help is the response to a help request.
func (a *Application) Help(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.HelpTitle), l.GetAny(loca.Help), ""
}

// Stop is the response to stop the skill.
func (a *Application) Stop(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.StopTitle), l.GetAny(loca.Stop), ""
}

// SaySomething handles simple title + text response.
func (a *Application) SaySomething(loc l10n.LocaleInstance, opts ...ResponseFunc) (ApplicationResponse, error) {
	// run all ResponseFuncs
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}
	//stats.Inc(a, "SaySomething", 1, 1.0)
	//stats.Timing(a,"SaySomething", 1.0, 1.0)

	tit := ""
	msg := ""
	msgSSML := ""
	if cfg.user != "" {
		// personalized response
		tit = loc.GetAny(loca.SaySomethingUserTitle, cfg.user)
		msg = loc.GetAny(loca.SaySomethingUserText, cfg.user)
		msgSSML = loc.GetAny(loca.SaySomethingUserSSML, cfg.user)
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
		End:    false,
	}, nil
}

// SSMLDemo is the intent to demonstrate SSML output with Alexa.
func (a *Application) SSMLDemo(l l10n.LocaleInstance) (string, string, string) {
	return l.GetAny(loca.LaunchTitle), l.GetAny(loca.LaunchText), l.GetAny(loca.LaunchSSML)
}

// Demo is a simple demo response.
func (a *Application) Demo(l l10n.LocaleInstance) (string, string, string) {
	return l.Get(loca.DemoIntentTitle), l.GetAny(loca.DemoIntentText), l.GetAny(loca.DemoIntentSSML)
}

// AWSStatus responds with messages containting 2 slots
func (a *Application) AWSStatus(loc l10n.LocaleInstance, area string, region string) (ApplicationResponse, error) {
	title := loc.GetAny(loca.AWSStatusTitle)
	msg := loc.GetAny(loca.AWSStatusText, area, region)
	msgSSML := loc.GetAny(loca.AWSStatusSSML, area, region)

	if title == "" || msg == "" || msgSSML == "" {
		return ApplicationResponse{}, ErrorNoTranslation
	}

	return ApplicationResponse{
		Title:  title,
		Text:   msg,
		Speech: msgSSML,
		Image:  "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_%s.png",
		End:    true,
	}, nil
}

//func (a *Application) AWSStatus(l l10n.LocaleInstance, region string) (string, string, string) {
//	// TODO: we need to have access to slot values here!
//	// request (with slot values) status from AWS status provider
//	// decide how to respond based on status results
//	// return response texts
//	text := loca.AWSStatusText
//	ssml := loca.AWSStatusSSML
//	if region == "Frankfurt" {
//		text = loca.AWSStatusTextGood
//		ssml = loca.AWSStatusSSMLGood
//	}
//	return l.GetAny(loca.AWSStatusTitle), l.GetAny(text, region), l.GetAny(ssml, region)
//}

// Logger returns the application logger.
func (a *Application) Logger() log.Logger {
	return a.logger
}

// Statter returns the application statter.
func (a *Application) Statter() stats.Statter {
	return a.statter
}
