package main

import (
	"fmt"

	"github.com/hamba/cmd"
	"github.com/hamba/pkg/log"
	"github.com/urfave/cli/v2"
)

func runServer(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	_, err = newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// TODO!
	//s := rpc.NewServer(app)
	log.Info(ctx, fmt.Sprintf("Starting lambda server"))
	// if err :=

	return nil
}
