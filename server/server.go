package server

import (
	"fmt"

	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

type Application interface {
	Handle()
	Help() (string, string)
}

// NewMux creates a new Mux instance.
func NewServer(app Application) {
	lambda.Start(HandleRequest)
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
	case alexa.LocaleGerman:
		text = "Hallo Welt"
	case alexa.LocaleJapanese:
		text = "こんにちは世界"
	default:
		text = "Hello, World"
	}
	return alexa.NewSimpleResponse(title, text)
}
