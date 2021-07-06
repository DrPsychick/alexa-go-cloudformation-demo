package main

import (
	"encoding/json"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/hamba/cmd"
	"github.com/hamba/logger"
	"github.com/hamba/pkg/log"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
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

	ms, err := createModels(sk)
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
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0644); err != nil {
			log.Fatal(ctx, err)
		}
	}

	if c.Bool("models") {
		os.MkdirAll("./alexa/interactionModels/custom", 0755)
		for l, m := range ms {
			var filename = "./alexa/interactionModels/custom/" + string(l) + ".json"

			res, _ := json.MarshalIndent(m, "", "  ")
			if err := ioutil.WriteFile(filename, res, 0644); err != nil {
				log.Fatal(ctx, err)
			}
		}
	}

	return nil
}

// createModels generates and returns a list of Models.
func createModels(s *gen.SkillBuilder) (map[string]*alexa.Model, error) {
	m := s.Model().
		WithDelegationStrategy(alexa.DelegationSkillResponse)

	// we define intents, slots, types in lambda,
	// that's why `newLambda` must be called before this, if not it will panic.

	// Prompts are part of the Alexa dialog, so independent of lambda.
	m.WithElicitationSlotPrompt(loca.AWSStatus, loca.TypeRegionName)
	// add variations (texts) to the prompt
	m.ElicitationPrompt(loca.AWSStatus, loca.TypeRegionName).
		WithVariation("PlainText").
		WithVariation("SSML")

	m.WithConfirmationSlotPrompt(loca.AWSStatus, loca.TypeAreaName)
	m.ConfirmationPrompt(loca.AWSStatus, loca.TypeAreaName).
		WithVariation("SSML")

	// create a Validation prompt, connected to type-values
	m.WithValidationSlotPrompt(loca.TypeRegionName, alexa.ValidationTypeHasMatch)
	m.ValidationPrompt(loca.TypeRegionName, alexa.ValidationTypeHasMatch).
		WithVariation("PlainText")

	// ValidationTypeInSet requires values -> we need to pass a key
	m.WithValidationSlotPrompt(loca.TypeRegionName, alexa.ValidationTypeInSet, loca.TypeRegionValues)
	m.ValidationPrompt(loca.TypeRegionName, alexa.ValidationTypeInSet).
		WithVariation("PlainText")

	//m.Intent(loca.AWSStatus).Slot(loca.TypeRegionName).
	//	WithValidationRule(alexa.ValidationTypeHasMatch)

	return s.BuildModels()
}
