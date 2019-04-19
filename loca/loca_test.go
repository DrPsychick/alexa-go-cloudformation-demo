package loca_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestL10NSetup(t *testing.T) {
	l, err := l10n.Resolve("de-DE")
	assert.NoError(t, err)
	assert.NotEmpty(t, l.GetName())
	assert.Equal(t, "de-DE", l.GetName())
}
