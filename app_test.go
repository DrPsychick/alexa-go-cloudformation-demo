package alfalfa_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/hamba/logger"
	"github.com/hamba/pkg/log"
	"github.com/hamba/pkg/stats"
	"github.com/hamba/statter/l2met"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApplication_ResponseFunc(t *testing.T) {
	c := alfalfa.Config{}
	r := alfalfa.WithUser("Martin")
	r(&c)

	assert.Equal(t, alfalfa.Config{"Martin"}, c)
}

func TestApplication_Launch(t *testing.T) {
	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	title, text := app.Launch(loc)

	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
}

func TestApplication_Help(t *testing.T) {
	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	title, text, _ := app.Help(loc)

	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
}

func TestApplication_Stop(t *testing.T) {
	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	title, text, _ := app.Stop(loc)

	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
}

func TestApplication_SSLDemo(t *testing.T) {
	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	title, text, ssml := app.SSMLDemo(loc)

	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
	assert.NotEmpty(t, ssml)
}

func TestApplication_Demo(t *testing.T) {
	app := alfalfa.NewApplication(log.Null, stats.Null)
	loc, err := loca.Registry.Resolve("en-US")
	assert.NoError(t, err)

	title, text, ssml := app.Demo(loc)

	assert.NotEmpty(t, title)
	assert.NotEmpty(t, text)
	assert.NotEmpty(t, ssml)
}

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

	// with user
	r := alfalfa.WithUser("John")
	resp, err = app.SaySomething(loc, r)

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
