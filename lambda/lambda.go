package lambda

import (
	"fmt"

	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	Handle()
	Help() (string, string)
	Stop() (string, string)
}

type Handler func(alexa.Request) (alexa.Response, error)

// HandleRequest is the lambda hander
func HandleRequest(app Application) Handler {
	return func(r alexa.Request) (alexa.Response, error) {
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
				title, text := app.Stop()
				return alexa.NewSimpleResponse(title, text), nil

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

func handleLaunch(request alexa.Request) alexa.Response {
	title := "Launch"
	text := "Willkommen beim Karlsruhe Golang Meetup #3!"
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
		"<voice name=\"Justin\"><lang xml:lang=\"en-US\">pace</lang></voice>" +
		"<voice name=\"Hans\">iss <emphasis level=\"strong\">geil!</emphasis></voice>" +
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
