package lambda

import (
	"fmt"

	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	Handle()
	Help() (string, string)
}

type Handler func(alexa.Request) (alexa.Response, error)

// HandleRequest is the lambda hander
func HandleRequest(app Application) Handler {
	return func(r alexa.Request) (alexa.Response, error) {
		name := r.Body.Intent.Name

		switch name {
		case "hello":
			return handleHello(r), nil

		case alexa.HelpIntent:
			title, text := app.Help()
			return alexa.NewSimpleResponse(title, text), nil
		}

		return alexa.Response{}, fmt.Errorf("server: unknown intent %s", name)
	}
}

func handleHello(request alexa.Request) alexa.Response {
	title := "Saying Hello"
	var text string
	switch request.Body.Locale {
	case alexa.LocaleAustralianEnglish:
		text = "G'day mate!"
	case alexa.LocaleGermanGerman:
		text = "Hallo Welt"
	case alexa.LocaleJapanese:
		text = "こんにちは世界"
	default:
		text = "Hello, World"
	}
	return alexa.NewSimpleResponse(title, text)
}
