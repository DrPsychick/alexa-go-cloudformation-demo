package main

import (
	"github.com/DrPsychick/alexa-go-cloudformation-demo"
	"github.com/hamba/cmd"
)

// Application =============================

func newApplication(c *cmd.Context) (*alfalfa.Application, error) {
	app := alfalfa.NewApplication(
		c.Logger(),
		c.Statter(),
	)

	return app, nil
}
