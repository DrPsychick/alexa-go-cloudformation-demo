package lambda

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	Help() (string, string)
	Stop(l *l10n.Locale) (string, string, string)
	SSMLDemo(l *l10n.Locale) (string, string, string)
	SaySomething(l *l10n.Locale) (string, string, string)
}

func NewMux(app Application) alexa.Handler {
	mux := alexa.NewServerMux()

	mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch)

	mux.HandleIntent(alexa.HelpIntent, handleHelp(app))
	mux.HandleIntent(alexa.CancelIntent, handleStop(app))
	mux.HandleIntent(alexa.StopIntent, handleStop(app))

	mux.HandleIntent("SSMLDemoIntent", handleSSMLResponse(app))
	mux.HandleIntent("SaySomething", handleSSMLResponse(app))
	mux.HandleIntentFunc("DemoIntent", handleDemo)

	return mux
}

func handleLaunch(r *alexa.Request) alexa.Response {
	title := "Launch"
	//text := "Willkommen beim Karlsruhe Golang Meetup #3!"
	text := "Ja?"

	resp := alexa.NewSimpleResponse(title, text)
	resp.Body.ShouldEndSession = false

	return resp
}

func handleHelp(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(r *alexa.Request) alexa.Response {
		title, text := app.Help()

		return alexa.NewSimpleResponse(title, text)
	})
}

func handleStop(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(r *alexa.Request) alexa.Response {
		title, text, _ := app.Stop(r.Locale)
		return alexa.NewSimpleResponse(title, text)
	})
}

func handleSSMLResponse(app Application) alexa.Handler {
	return alexa.HandlerFunc(func(r *alexa.Request) alexa.Response {
		title, text, ssmlText := app.SSMLDemo(r.Locale)

		resp := alexa.NewSimpleResponse(title, text)
		resp.Body.OutputSpeech.Type = "SSML"
		resp.Body.OutputSpeech.SSML = ssmlText
		return resp
	})
}

func handleDemo(r *alexa.Request) alexa.Response {
	title := "Test"
	text := "Pace ist geil!"

	resp := alexa.NewSimpleResponse(title, text)
	resp.Body.OutputSpeech.Type = "SSML"
	resp.Body.OutputSpeech.SSML = "<speak>" +
		"<voice name=\"Kendra\"><lang xml:lang=\"en-US\">" +
		"<emphasis level=\"strong\">pace</emphasis>" +
		"</lang></voice>" +
		"<voice name=\"Marlene\">iss <emphasis level=\"strong\">geil!</emphasis></voice>" +
		"</speak>"
	return resp
}
