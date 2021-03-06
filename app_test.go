package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationHelp(t *testing.T) {
	a := NewApplication(nil, nil)
	l, err := l10n.Resolve("de-DE")
	title, text, _ := a.Help(l)

	assert.Nil(t, err, "could not resolve locale 'de-DE'!")
	assert.Equal(t, "de-DE", l.GetName())

	title, text, ssmlText := a.SaySomething(l)
	assert.NotEmpty(t, title, "'title' must not be empty")
	assert.NotEmpty(t, text, "'text' must not be empty")
	assert.NotEmpty(t, ssmlText, "'ssmlText' must not be empty")
	//fmt.Printf("ssml text: %s\n", ssmlText)

	title, text, ssmlText = a.SSMLDemo(l)
	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
	assert.NotEmpty(t, ssmlText)
	//fmt.Printf("ssml text: %s\n", ssmlText)

	//assert.Nil(t, l.Name)
}
