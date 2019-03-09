package main

import (
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
		Action: runGenerate,
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "unknown"
	app.Version = version
	app.Commands = commands

	app.Run(os.Args)
}