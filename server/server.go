// Package server is a standalone http server (as replacement for lambda)
package server

import (
	"github.com/drpsychick/go-alexa-lambda"
)

// Application defines the interface to the app.
type Application interface {
	Handle()
	Help() (string, string)
}

// NewServer creates a new Mux instance.
func NewServer(app Application) alexa.Handler {
	return nil
}

// Handler defines the handler function.
type Handler func(alexa.RequestEnvelope) (alexa.ResponseEnvelope, error)
