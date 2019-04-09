package server

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

type Application interface {
	Handle()
	Help() (string, string)
}

// NewMux creates a new Mux instance.
func NewServer(app Application) alexa.Handler {
	return nil
}

type Handler func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error)

// HandleRequest is the lambda hander
//func HandleRequest(app Application) Handler {
//	return func(r alexa.RequestEnvelope) (alexa.ResponseEnvelope, error) {
//		name := r.Request.Intent.Name
//
//		switch name {
//		case "hello":
//			return handleHello(r), nil
//
//		case alexa.HelpIntent:
//			title, text := app.Help()
//			return alexa.NewSimpleResponse(title, text), nil
//		}
//
//		return alexa.ResponseEnvelope{}, fmt.Errorf("server: unknown intent %s", name)
//	}
//}
//
//func handleHello(request alexa.RequestEnvelope) alexa.ResponseEnvelope {
//	title := "Saying Hello"
//	var text string
//	switch request.Request.Locale {
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
