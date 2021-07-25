package alfalfa_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/hamba/logger"
	"github.com/hamba/statter/l2met"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApplication_SaySomething(t *testing.T) {
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	loc, err := loca.Registry.Resolve("de-DE")
	assert.NoError(t, err)

	resp, err := app.SaySomething(loc)

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Title)
	assert.NotEmpty(t, resp.Text)
	assert.NotEmpty(t, resp.Speech)
}

func TestApplication_AWSStatus(t *testing.T) {
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	loc, err := loca.Registry.Resolve("de-DE")
	assert.NoError(t, err)

	resp, err := app.AWSStatus(loc, "Europa", "Frankfurt")

	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Title)
	assert.NotEmpty(t, resp.Text)
	assert.NotEmpty(t, resp.Speech)
	assert.Contains(t, resp.Text, "Europa")
	assert.Contains(t, resp.Text, "Frankfurt")
}

//func TestApplication_SaySomething2_Personalized(t *testing.T) {
//	loc, err := loca.Registry.Resolve("de-DE")
//	assert.NoError(t, err)
//	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
//	app := alfalfa.NewApplication(
//		l,
//		l2met.New(l, ""),
//	)
//
//	f := app.SaySomething2()
//	res, err := f(loc, alfalfa.WithUser("Fred"))
//
//	assert.NoError(t, err)
//	assert.NotEmpty(t, res.Title)
//	assert.NotEmpty(t, res.Text)
//	assert.NotEmpty(t, res.Speech)
//	assert.Contains(t, res.Text, "Fred")
//}
//
//func TestApplication_SaySomething2_ErrorNoLocale(t *testing.T) {
//	r := l10n.NewRegistry()
//	loc := l10n.NewLocale("en-US")
//	loc.Set(l10n.KeyErrorNoTranslationTitle, []string{"missing translation"})
//	loc.Set(l10n.KeyErrorNoTranslationText, []string{"no translation found for '%s'"})
//	err := r.Register(loc)
//	assert.NoError(t, err)
//	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
//	app := alfalfa.NewApplication(
//		l,
//		l2met.New(l, ""),
//	)
//
//	f := app.SaySomething2()
//	res, err := f(loc)
//	title := ""
//	text := ""
//	switch err {
//	case alfalfa.ErrorNoTranslation:
//		title = loc.GetAny(l10n.KeyErrorNoTranslationTitle)
//		text = loc.GetAny(l10n.KeyErrorNoTranslationText, loca.SaySomethingText)
//	}
//
//	assert.Error(t, err)
//	assert.Equal(t, alfalfa.ErrorNoTranslation, err)
//	assert.Empty(t, res.Title)
//	assert.Empty(t, res.Text)
//	assert.NotEmpty(t, title)
//	assert.NotEmpty(t, text)
//	assert.Contains(t, text, loca.SaySomethingText)
//}
