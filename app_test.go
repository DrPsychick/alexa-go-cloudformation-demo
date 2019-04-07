package alfalfa

import (
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationHelp(t *testing.T) {
	a := NewApplication(nil, nil)
	title, text := a.Help()

	l, err := l10n.Resolve("de-DE")
	assert.Nil(t, err, "could not resolve locale 'de-DE'!")

	title, text = a.SaySomething(l)
	assert.NotEmpty(t, title, "'title' must not be empty")
	assert.NotEmpty(t, text, "'text' must not be empty")

	title, text, ssmlText := a.SSMLDemo(l)
	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
	assert.NotEmpty(t, ssmlText)
}
