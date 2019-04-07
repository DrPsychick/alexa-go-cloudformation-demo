package lambda

import (
	"fmt"
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"

	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	//Handle()
	Help() (string, string)
	Stop(l *l10n.Locale) (string, string, string)
	SSMLDemo(l *l10n.Locale) (string, string, string)
	SaySomething(l *l10n.Locale) (string, string)
}

type Handler func(alexa.Request) (alexa.Response, error)

// HandleRequest is the lambda hander
func HandleRequest(app Application) Handler {
	return func(r alexa.Request) (alexa.Response, error) {
		l, err := l10n.Resolve(string(r.Body.Locale))
		if err != nil {
			return alexa.Response{}, fmt.Errorf("could not resolve locale %s", string(r.Body.Locale))
		}

		if r.Body.Type == alexa.TypeLaunchRequest {
			return handleLaunch(r), nil
		}
		if r.Body.Type == alexa.TypeIntentRequest {
			name := r.Body.Intent.Name

			switch name {
			case alexa.HelpIntent:
				title, text := app.Help()
				return alexa.NewSimpleResponse(title, text), nil

			case alexa.CancelIntent:
				fallthrough
			case alexa.StopIntent:
				title, text, _ := app.Stop(l)
				return alexa.NewSimpleResponse(title, text), nil

			case "SSMLDemoIntent":
				return handleSSMLResponse(app.SSMLDemo(l)), nil

			case "SaySomething":
				return handleSimpleResponse(app.SaySomething(l)), nil

			case "DemoIntent":
				return handleDemo(r), nil
			}
			return alexa.Response{}, fmt.Errorf("server: unknown intent %s", name)
		}
		if r.Body.Type == alexa.TypeSessionEndedRequest {

		}
		if r.Body.Type == alexa.TypeCanFulfillIntentRequest {

		}
		return alexa.Response{}, fmt.Errorf("server: unknown intent type %s", r.Body.Type)
	}
}

func handleSimpleResponse(title string, text string) alexa.Response {
	return alexa.NewSimpleResponse(title, text)
}

func handleSSMLResponse(title string, text string, ssmlText string) alexa.Response {
	r := alexa.NewSimpleResponse(title, text)
	r.Body.OutputSpeech.Type = "SSML"
	r.Body.OutputSpeech.SSML = ssmlText
	return r
}

func handleLaunch(request alexa.Request) alexa.Response {
	title := "Launch"
	//text := "Willkommen beim Karlsruhe Golang Meetup #3!"
	text := "Ja?"
	r := alexa.NewSimpleResponse(title, text)
	r.Body.ShouldEndSession = false
	return r
}

func handleDemo(request alexa.Request) alexa.Response {
	title := "Test"
	text := "Pace ist geil!"
	r := alexa.NewSimpleResponse(title, text)
	r.Body.OutputSpeech.Type = "SSML"
	r.Body.OutputSpeech.SSML = "<speak>" +
		"<voice name=\"Kendra\"><lang xml:lang=\"en-US\">" +
		"<emphasis level=\"strong\">pace</emphasis>" +
		"</lang></voice>" +
		"<voice name=\"Marlene\">iss <emphasis level=\"strong\">geil!</emphasis></voice>" +
		"</speak>"
	return r
}

//func handleHello(request alexa.Request) alexa.Response {
//	title := "Saying Hello"
//	var text string
//	switch request.Body.Locale {
//	case alexa.LocaleAustralianEnglish:
//		text = "G'day mate!"
//	case alexa.LocaleGerman:
//		text = "Hallo Welt"
//	case alexa.LocaleJapanese:
//		text = "こんにちは世界"
//	default:
//		text = "Hello, World"
//	}
//	return alexa.NewSimpleResponse(title, text)
//}
