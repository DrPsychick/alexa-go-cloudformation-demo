package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
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
	return l.Get(loca.DemoIntentTitle), l.GetAny(loca.DemoIntentText), l.GetAny(loca.DemoIntentSSML)
}

func (a *Application) AWSStatus(l l10n.LocaleInstance, region string) (string, string, string) {
	// TODO: we need to have access to slot values here!
	// request (with slot values) status from AWS status provider
	// decide how to respond based on status results
	// return response texts
	text := loca.AWSStatusText
	ssml := loca.AWSStatusSSML
	if region == "Frankfurt" {
		text = loca.AWSStatusTextGood
		ssml = loca.AWSStatusSSMLGood
	}
	return l.GetAny(loca.AWSStatusTitle), l.GetAny(text, region), l.GetAny(ssml, region)
}

func (a *Application) AWSStatusRegionElicit(l l10n.LocaleInstance, region string) (string, string, string) {
	text := loca.AWSStatusRegionElicitText
	ssml := loca.AWSStatusRegionElicitSSML
	return l.GetAny(loca.AWSStatusTitle), l.GetAny(text, region), l.GetAny(ssml, region)
}

// Logger returns the application logger.
func (a *Application) Logger() log.Logger {
	return a.logger
}

// Statter returns the application statter.
func (a *Application) Statter() stats.Statter {
	return a.statter
}
