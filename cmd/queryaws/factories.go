package main

import (
	"github.com/hamba/cmd"
)

// Application =============================

func newApplication(c *cmd.Context) (*queryaws.Application, error) {
	app := queryaws.NewApplication(
		c.Logger(),
		c.Statter(),
	)

	return app, nil
}
