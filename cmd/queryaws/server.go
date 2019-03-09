package main

import (
	"fmt"
	"github.com/hamba/cmd"
	"github.com/hamba/pkg/log"
	"gopkg.in/urfave/cli.v1"
	"net/http"
	"net/rpc"
)

func runServer(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	s := rpc.NewServer(app)
	log.Info(ctx, fmt.Sprintf("Starting lambda server"))
	// if err :=
}

func runMake(c *cli.Context) error {

}

func newServer(app *queryaws.Application) http.Handler
