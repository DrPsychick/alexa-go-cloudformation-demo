package main

import (
	lambda2 "github.com/DrPsychick/alexa-go-cloudformation-demo/lambda"
	"github.com/aws/aws-lambda-go/lambda"
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

	lambda.Start(lambda2.HandleRequest(app))

	return nil
}
