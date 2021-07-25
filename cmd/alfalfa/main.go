package main

import (
	"github.com/hamba/cmd"
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
		Flags:  cmd.Flags{}.Merge(cmd.CommonFlags, cmd.ServerFlags),
	},
	{
		Name:   "lambda",
		Usage:  "Run the lambda server",
		Action: runLambda,
		Flags: cmd.Flags{
			cli.IntFlag{
				Name:   "lambda.port",
				Usage:  "Port on which lambda will listen",
				EnvVar: "_LAMBDA_SERVER_PORT",
			},
		}.Merge(cmd.CommonFlags, cmd.ServerFlags),
	},
	{
		Name:  "make",
		Usage: "Make Alexa skill files",
		Flags: cmd.Flags{
			cli.BoolFlag{
				Name:   "skill",
				Usage:  "Generate Alexa skill.json",
				EnvVar: "ALFALFA_MAKE_SKILL",
			},
			cli.BoolFlag{
				Name:   "models",
				Usage:  "Generate Alexa interaction model JSON files",
				EnvVar: "ALFALFA_MAKE_MODELS",
			},
		}.Merge(cmd.CommonFlags, cmd.ServerFlags),
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
