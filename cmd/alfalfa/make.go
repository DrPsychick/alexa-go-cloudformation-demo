package main

import (
	"encoding/json"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"log"
)

func runMake(c *cli.Context) error {
	// build skill and models
	sk, err := createSkill(l10n.DefaultRegistry)
	if err != nil {
		return err
	}
	ms, err := createModels(sk)
	if err != nil {
		return err
	}

	// build and write JSON files
	if c.Bool("skill") {
		s, err := sk.Build()
		if err != nil {
			log.Fatal(err)
		}
		res, _ := json.MarshalIndent(s, "", "  ")
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0644); err != nil {
			log.Fatal(err)
		}
	}

	if c.Bool("models") {
		for l, m := range ms {
			var filename = "./alexa/interactionModels/custom/" + string(l) + ".json"

			res, _ := json.MarshalIndent(m, "", "  ")
			if err := ioutil.WriteFile(filename, res, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}

// createSkill generates and returns a SkillBuilder.
func createSkill(r l10n.LocaleRegistry) (*gen.SkillBuilder, error) {
	skill := gen.NewSkillBuilder().
		WithLocaleRegistry(r).
		WithCategory(alexa.CategoryOrganizersAndAssistants).
		WithPrivacyFlag(gen.FlagIsExportCompliant, true)

	return skill, nil
}

// createModels generates and returns a list of Models.
func createModels(s *gen.SkillBuilder) (map[string]*alexa.Model, error) {
	m := s.WithModel().Model().
		WithDelegationStrategy(alexa.DelegationSkillResponse) // general delegation to lambda

	m.WithType(loca.TypeArea)
	m.WithType(loca.TypeRegion)

	m.WithIntent(loca.DemoIntent)
	m.WithIntent(loca.SaySomething)
	m.WithIntent(loca.AWSStatus)

	// requires
	m.WithIntent(alexa.HelpIntent)
	m.WithIntent(alexa.CancelIntent)
	m.WithIntent(alexa.StopIntent)

	m.Intent(loca.AWSStatus).
		WithSlot(loca.TypeAreaName, loca.TypeArea).
		WithSlot(loca.TypeRegionName, loca.TypeRegion).
		WithDelegation(alexa.DelegationAlways) // ALWAYS = delegate specific intent dialog to alexa

	// make an elicitation prompt and link it to AWSStatus intent, AWSRegion slot
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
