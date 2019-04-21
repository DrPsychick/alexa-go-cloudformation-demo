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

// Alexa skill definition (to generate skill.json)
var skill = alexa.Skill{
	Manifest: alexa.Manifest{
		Version: "1.0",
		Publishing: alexa.Publishing{
			Locales: map[string]alexa.LocaleDef{
				"de-DE": {
					Name:         "Demo Skill",
					Description:  "Demo for the golang meetup",
					Summary:      "This skill demonstrates what you can do with the alexa package and cloudformation",
					Keywords:     []string{"demo", "test", "SSML", "cloudformation", "automation"},
					Examples:     []string{"Alexa, start demo skill and say something"},
					SmallIconURI: "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_small.png",
					LargeIconURI: "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_large.png",
				},
			},
			Category:            alexa.CategoryOrganizersAndAssistants,
			Countries:           []string{"DE"},
			TestingInstructions: "Alexa, open demo skill. Yes? Say something.",
		},
		//Apis: alexa.Apis{
		//	Custom: alexa.Custom{
		//		Endpoint: alexa.Endpoint{
		//			URI: "arn:...",
		//		},
		//	},
		//	Interfaces: []string{},
		//},
		Permissions: []alexa.Permission{},
		Privacy: &alexa.Privacy{
			IsExportCompliant: true,
			//Locales: map[alexa.Locale]alexa.PrivacyLocaleDef{
			//	"de-DE": {
			//
			//	},
			//},
		},
	},
}

var models = map[string]alexa.Model{
	"de-DE": modelGerman,
}

var modelGerman = alexa.Model{
	Model: alexa.InteractionModel{
		Language: alexa.LanguageModel{
			Invocation: "demo skill",
			Intents: []alexa.ModelIntent{
				{Name: "AMAZON.CancelIntent", Samples: []string{}},
				{Name: "AMAZON.HelpIntent", Samples: []string{}},
				{Name: "AMAZON.StopIntent", Samples: []string{}},
				{Name: "DemoIntent", Samples: []string{
					"Schiess' los",
					"Auf geht's",
					"Hop hop",
				}},
				//{Name: "SSMLDemoIntent", Samples: []string{
				//	"Zeig' was du kannst",
				//	"Immer her damit",
				//	"Was kann SSML",
				//}},
				{Name: "SaySomething", Samples: []string{
					"Erzähl' mir was",
					"Sag was",
				}},
				{
					Name: "AWSStatus",
					Samples: []string{
						"A.W.S. status of {Area}",
						"status of {Area}",
						"give me the status of {Region}",
						"status of {Region}",
						"{Region} status",
					},
					Slots: []alexa.ModelSlot{
						{Name: "Area", Type: "AWSArea", Samples: []string{"of {Area}"}},
						{Name: "Region", Type: "AWSRegion", Samples: []string{"of {Region}", "in {Region}"}},
					},
				},
			},
			Types: []alexa.ModelType{
				{Name: "AWSArea", Values: []alexa.TypeValue{
					{Name: alexa.NameValue{Value: "Europe"}},
					{Name: alexa.NameValue{Value: "North America"}},
					{Name: alexa.NameValue{Value: "South America"}},
					{Name: alexa.NameValue{Value: "Asia Pacific"}},
				}},
				{Name: "AWSRegion", Values: []alexa.TypeValue{
					{Name: alexa.NameValue{Value: "Frankfurt"}},
					{Name: alexa.NameValue{Value: "Ireland"}},
					{Name: alexa.NameValue{Value: "London"}},
					{Name: alexa.NameValue{Value: "Paris"}},
					{Name: alexa.NameValue{Value: "Stockholm"}},
					{Name: alexa.NameValue{Value: "North Virginia"}},
				}},
			},
		},
		Dialog: &alexa.Dialog{
			Delegation: alexa.DelegationSkillResponse,
			Intents: []alexa.DialogIntent{
				{Name: "AWSStatus", Confirmation: false, Slots: []alexa.DialogIntentSlot{
					{Name: "Area", Type: "AWSArea", Prompts: alexa.SlotPrompts{
						Elicitation: "Elicit.Intent-AWSStatus.IntentSlot-Area",
					}},
					{Name: "Region", Type: "AWSRegion", Prompts: alexa.SlotPrompts{
						Elicitation: "Elicit.Intent-AWSStatus.IntentSlot-Region",
					}},
				}},
			},
		},
		Prompts: []alexa.ModelPrompt{
			{Id: "Elicit.Intent-AWSStatus.IntentSlot-Area", Variations: []alexa.PromptVariation{
				{Type: "PlainText", Value: "From what area do you seek status?"},
			}},
			{Id: "Elicit.Intent-AWSStatus.IntentSlot-Region", Variations: []alexa.PromptVariation{
				{Type: "PlainText", Value: "From what region do you want to know the status?"},
			}},
		},
	},
}

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
	m := s.AddModel().
		WithDelegationStrategy(alexa.DelegationSkillResponse)

	m.AddType(loca.TypeArea)
	m.AddType(loca.TypeRegion)

	m.AddIntent(loca.DemoIntent)
	m.AddIntent(loca.SaySomething)

	i := m.AddIntent(loca.AWSStatus)
	i.AddSlot(loca.TypeAreaName, loca.TypeArea)
	i.AddSlot(loca.TypeRegionName, loca.TypeRegion)

	pb := m.AddElicitationSlotPrompt(loca.AWSStatus, loca.TypeRegionName)
	pb.AddVariation("PlainText").
		AddVariation("SSML")
	pb = m.AddConfirmationSlotPrompt(loca.AWSStatus, loca.TypeAreaName)
	pb.AddVariation("SSML")
	if pb == nil {
		return nil, fmt.Errorf("Elicitation prompt failed to add")
	}

	return s.BuildModels()
}
