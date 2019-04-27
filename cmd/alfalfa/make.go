package main

import (
	"encoding/json"
	"fmt"
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
		WithDelegationStrategy(alexa.DelegationSkillResponse)

	m.AddType(loca.TypeArea)
	m.AddType(loca.TypeRegion)

	m.AddIntent(loca.DemoIntent)
	m.AddIntent(loca.SaySomething)

	i := m.AddIntent(loca.AWSStatus)
	i.AddSlot(loca.TypeAreaName, loca.TypeArea)
	i.AddSlot(loca.TypeRegionName, loca.TypeRegion)

	pb := m.AddElicitationSlotPrompt(loca.AWSStatus, loca.TypeRegionName)
	if pb != nil {
		pb.AddVariation("PlainText").AddVariation("SSML")
	}
	pb = m.AddConfirmationSlotPrompt(loca.AWSStatus, loca.TypeAreaName)
	if pb != nil {
		pb.AddVariation("SSML")
	}
	if pb == nil {
		return nil, fmt.Errorf("Elicitation prompt failed to add")
	}

	return s.BuildModels()
}
