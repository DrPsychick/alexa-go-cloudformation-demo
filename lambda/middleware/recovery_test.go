package middleware_test

import (
	"errors"
	"testing"

	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/go-alexa-lambda"
	"github.com/hamba/pkg/log"
)

func TestWithRecovery(t *testing.T) {
	h := middleware.WithRecovery(
		alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
			panic("panic")
		}),
		log.NewMockLoggable(log.Null),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{}

	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Expected the panic to be handled.")
		}
	}()

	h.Serve(bdr, req)
}

func TestWithRecovery_Error(t *testing.T) {
	h := middleware.WithRecovery(
		alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.RequestEnvelope) {
			panic(errors.New("panic"))
		}),
		log.NewMockLoggable(log.Null),
	)

	bdr := &alexa.ResponseBuilder{}
	req := &alexa.RequestEnvelope{}

	defer func() {
		if err := recover(); err != nil {
			t.Fatal("Expected the panic to be handled.")
		}
	}()

	h.Serve(bdr, req)
}
