package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/hamba/cmd"
	"github.com/hamba/logger"
	"github.com/hamba/pkg/log"
	"github.com/urfave/cli/v2"
)

func runMake(c *cli.Context) error {
	ctx, err := cmd.NewContext(c)
	if err != nil {
		return err
	}

	// attach a unbuffered logger:
	lg := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	ctx.AttachLogger(func(l log.Logger) log.Logger {
		return lg
	})

	app, err := newApplication(ctx)
	if err != nil {
		log.Fatal(ctx, err.Error())
	}

	// build skill and models
	sk := newSkill()

	// lambda injects supported intents, slots, types
	newLambda(app, sk)

	ms, err := createSkillModels(sk)
	if err != nil {
		return err
	}

	// build and write JSON files
	if c.Bool("skill") {
		s, err := sk.Build()
		if err != nil {
			log.Fatal(ctx, err)
		}
		res, _ := json.MarshalIndent(s, "", "  ")
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0o644); err != nil {
			log.Fatal(ctx, err)
		}
	}

	if c.Bool("models") {
		if err := os.MkdirAll("./alexa/interactionModels/custom", 0o755); err != nil {
			log.Fatal(ctx, "could not create directory ./alexa/interactionModels/custom")
			return err
		}

		for l, m := range ms {
			filename := "./alexa/interactionModels/custom/" + l + ".json"

			res, _ := json.MarshalIndent(m, "", "  ")
			if err := ioutil.WriteFile(filename, res, 0o644); err != nil {
				log.Fatal(ctx, err)
			}
		}
	}

	return nil
}
