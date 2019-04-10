package alexa_test

import (
	"testing"

	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/stretchr/testify/assert"
)

func TestHandlerFunc_Serve(t *testing.T) {
	bdr := &alexa.ResponseBuilder{}
	req := &alexa.Request{}

	called := false
	fn := alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
		assert.Equal(t, bdr, b)
		assert.Equal(t, req, r)

		called = true
	})

	fn.Serve(bdr, req)

	assert.True(t, called)
}
