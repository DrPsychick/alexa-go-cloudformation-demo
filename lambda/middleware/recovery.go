// Package middleware for lambda requests
package middleware

import (
	"fmt"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/pkg/log"
)

// Recovery is a middleware that will recover from panics and logs the error.
type Recovery struct {
	handler alexa.Handler
	l       log.Logger
}

// WithRecovery recovers from panics and log the error.
func WithRecovery(h alexa.Handler, lable log.Loggable) alexa.Handler {
	return &Recovery{
		handler: h,
		l:       lable.Logger(),
	}
}

// Serve serves the request.
func (m Recovery) Serve(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
	defer func() {
		if v := recover(); v != nil {
			m.l.Error(fmt.Sprintf("%+v", v))
		}
	}()

	m.handler.Serve(b, r)
}
