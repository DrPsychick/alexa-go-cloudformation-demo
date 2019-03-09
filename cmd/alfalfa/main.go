package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

import _ "github.com/joho/godotenv/autoload"

var version = "v0.0.1"

var commands = []cli.Command{
	{
		Name:   "server",
		Usage:  "Run the lambda server",
		Action: runServer,
	},
	{
		Name:   "make",
		Usage:  "Make Alexa skill files",
		Action: runMake,
	},
}

func main() {
	app := cli.NewApp()
	app.Name = "Alfalfa"
	app.Usage = "It does stuff and stuff"
	app.Version = version
	app.Commands = commands
	app.Action = runLambda

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
