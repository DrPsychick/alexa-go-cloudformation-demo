package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"github.com/hamba/cmd"
)

import _ "github.com/joho/godotenv/autoload"

var version = "v0.0.1"

var commands = []cli.Command{
	{
		Name: "server",
		Usage: "Run the lambda server",
		Action: runServer,
	},
	{
		Name: "generate",
		Usage: "Generate Alexa skill files",
		Action: runMake,
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "unknown"
	app.Version = version
	app.Commands = commands

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}