package main

import (
	"github.com/hamba/cmd"
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

import _ "github.com/joho/godotenv/autoload"

var version = "v0.0.1"

// used when launched without command (e.g. lambda)
var default_flags = cmd.Flags{
	cli.IntFlag{
		Name:   "lambda.port",
		Usage:  "Port on which lambda will listen",
		EnvVar: "_LAMBDA_SERVER_PORT",
	},
	cli.StringFlag{
		Name:   cmd.FlagLogFormat,
		Usage:  "Specify the format of logs. Supported formats: 'logfmt', 'json'",
		EnvVar: "LOG_FORMAT",
	},
	cli.StringFlag{
		Name:   cmd.FlagLogLevel,
		Value:  "info",
		Usage:  "Specify the log level. E.g. 'debug', 'warning'.",
		EnvVar: "LOG_LEVEL",
	},
	cli.StringSliceFlag{
		Name:   cmd.FlagLogTags,
		Usage:  "A list of tags appended to every log. Format: key=value.",
		EnvVar: "LOG_TAGS",
	},
	cli.StringFlag{
		Name:   cmd.FlagStatsDSN,
		Value:  "l2met://log",
		Usage:  "The URL of a stats backend.",
		EnvVar: "STATS_DSN",
	},
	cli.StringFlag{
		Name:   cmd.FlagStatsPrefix,
		Usage:  "The prefix of the measurements names.",
		EnvVar: "STATS_PREFIX",
	},
	cli.StringSliceFlag{
		Name:   cmd.FlagStatsTags,
		Usage:  "A list of tags appended to every measurement. Format: key=value.",
		EnvVar: "STATS_TAGS",
	},
}

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
		}.Merge(default_flags, cmd.ServerFlags),
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
	// need to be set for default Action
	app.Flags = default_flags
	app.Action = runLambda

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
