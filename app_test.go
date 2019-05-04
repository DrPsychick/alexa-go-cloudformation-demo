package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApplicationHelp(t *testing.T) {
	a := NewApplication(nil, nil)
	title, text := a.Help()

	l, err := loca.Registry.Resolve("de-DE")
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

func TestApplication_SaySomething2(t *testing.T) {
	loc, err := loca.Registry.Resolve("de-DE")
	assert.NoError(t, err)
	app := NewApplication(nil, nil)

	f := app.SaySomething2()
	res, err := f(loc)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestApplication_SaySomething2_ErrorNoLocale(t *testing.T) {
	r := l10n.NewRegistry()
	loc := l10n.NewLocale("en-US")
	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"Missing translation"})
	loc.Set(l10n.KeyErrorNoTranslationText, []string{"No translation found for '%s'"})
	err := r.Register(loc)
	assert.NoError(t, err)
	app := NewApplication(nil, nil)

	f := app.SaySomething2()
	res, err := f(loc)
	title := ""
	text := ""
	switch err {
	case ErrorNoTranslation:
		title = loc.GetAny(l10n.KeyErrorNoTranslationTitle)
		text = loc.GetAny(l10n.KeyErrorNoTranslationText, loca.SaySomethingText)
	}

	assert.Error(t, err)
	assert.Equal(t, ErrorNoTranslation, err)
	assert.Empty(t, res.Title)
	assert.Empty(t, res.Text)
	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
	assert.Contains(t, text, loca.SaySomethingText)
}
