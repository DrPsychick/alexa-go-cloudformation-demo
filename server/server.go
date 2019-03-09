package server

import (
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/aws/aws-lambda-go/lambda"
)

type Application interface {
	Handle(alexa.Request) (alexa.Response, error)
}

// NewMux creates a new Mux instance.
func NewServer(app Application) {
	lambda.Start(Handler)
}

// Handler is the lambda hander
func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

// DispatchIntents dispatches each intent to the right handler
func DispatchIntents(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "hello":
		response = handleHello(request)
	case alexa.HelpIntent:
		response = handleHelp()
	}

	return response
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

func handleHelp() alexa.Response {
	return alexa.NewSimpleResponse("Help for Hello", "To receive a greeting, ask hello to say hello")
}
