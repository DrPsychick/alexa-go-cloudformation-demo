package main

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/drpsychick/alexa-go-cloudformation-demo/lambda/middleware"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/hamba/cmd"
	"github.com/hamba/pkg/log"
	"gopkg.in/urfave/cli.v1"
)

func runLambda(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	l := newLambda(app)
	if err := alexa.Serve(l); err != nil {
		log.Fatal(ctx, err)
	}

	return nil
}

func newLambda(app *alfalfa.Application) alexa.Handler {
	h := lambda.NewMux(app)

	h = middleware.WithRequestStats(h, app)
	return middleware.WithRecovery(h, app)
}
