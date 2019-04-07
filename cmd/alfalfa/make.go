package main

import (
	"encoding/json"
	"fmt"
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"log"
)

// Alexa skill definition (to generate skill.json)
var skill = alexa.Skill{
	Manifest: alexa.Manifest{
		Version: "1.0",
		Publishing: alexa.Publishing{
			Locales: map[alexa.Locale]alexa.LocaleDef{
				"de-DE": {
					Name:         "DemoSkill",
					Description:  "Demo for the golang meetup",
					Summary:      "Demo for deploying Alexa + Lambda with cloudformation",
					Keywords:     []string{"Cloudformation Demo"},
					Examples:     []string{"Schiess los"},
					SmallIconURI: "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_small.png",
					LargeIconURI: "https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_large.png",
				},
			},
			Category:            alexa.CategoryOrganizersAndAssistants,
			Countries:           []alexa.Country{"DE"},
			TestingInstructions: "Demo Alexa skill...",
		},
		//Apis: alexa.Apis{
		//	Custom: alexa.Custom{
		//		Endpoint: alexa.Endpoint{
		//			URI: "arn:...",
		//		},
		//	},
		//	Interfaces: []string{},
		//},
		Permissions: &[]alexa.Permission{},
		Privacy: &alexa.Privacy{
			IsExportCompliant: true,
			//Locales: &map[alexa.Locale]alexa.PrivacyLocaleDef{
			//	"de-DE": {
			//
			//	},
			//},
		},
	},
}

var models = map[alexa.Locale]alexa.Model{
	"de-DE": modelGerman,
}

var modelGerman = alexa.Model{
	Model: alexa.InteractionModel{
		Language: alexa.LanguageModel{
			Invocation: "golang meetup",
			Intents: []alexa.ModelIntent{
				{Name: "AMAZON.CancelIntent", Samples: []string{}},
				{Name: "AMAZON.HelpIntent", Samples: []string{}},
				{Name: "AMAZON.StopIntent", Samples: []string{}},
				{Name: "DemoIntent", Samples: []string{
					"Schiess' los",
					"Auf geht's",
					"Hop hop",
				}},
				{Name: "SSMLDemoIntent", Samples: []string{
					"Zeig' was du kannst!",
					"Immer her damit.",
					"Was kann SSML?",
				}},
				{Name: "SaySomething", Samples: []string{
					"Erzaehl' mir was",
					"Sag was",
				}},
				{
					Name: "BeerStatsIntent",
					Samples: []string{
						"wieviel Bier trinkt man in {Country}",
						"wieviel trinkt {Country}",
						"wieviel Bier trinken {PeopleCategory} in {Country}",
					},
					Slots: &[]alexa.ModelSlot{
						{Name: "Country", Type: "BEER_Countries", Samples: []string{"{Country}"}},
						{Name: "PeopleCategory", Type: "BEER_PeopleCategory", Samples: []string{"{PeopleCategory}", "von {PeopleCategory}"}},
					},
				},
			},
			Types: []alexa.ModelType{
				{Name: "BEER_Countries", Values: []alexa.TypeValue{
					{Name: alexa.NameValue{Value: "Deutschland"}},
				}},
				{Name: "BEER_PeopleCategory", Values: []alexa.TypeValue{
					{Name: alexa.NameValue{Value: "Alle"}},
					{Name: alexa.NameValue{Value: "Frauen"}},
					{Name: alexa.NameValue{Value: "Männern"}},
					{Name: alexa.NameValue{Value: "Teenagern"}},
					{Name: alexa.NameValue{Value: "Intellektuellen"}},
				}},
			},
		},
		Dialog: &alexa.Dialog{
			Delegation: alexa.SkillResponse,
			Intents: &[]alexa.DialogIntent{
				{Name: "BeerStatsIntent", Confirmation: false, Slots: []alexa.DialogIntentSlot{
					{Name: "Country", Type: "BEER_Countries", Prompts: alexa.SlotPrompts{
						Elicitation: "Elicit.Intent-BeerStatsIntent.IntentSlot-Country",
					}},
					{Name: "PeopleCategory", Type: "BEER_PeopleCategory", Prompts: alexa.SlotPrompts{
						Elicitation: "Elicit.Intent-BeerStatsIntent.IntentSlot-PeopleCategory",
					}},
				}},
			},
		},
		Prompts: &[]alexa.ModelPrompt{
			{Id: "Elicit.Intent-BeerStatsIntent.IntentSlot-Country", Variations: []alexa.PromptVariations{
				{Type: "PlainText", Value: "Für welches Land möchtest du Bier Statistiken?"},
			}},
			{Id: "Elicit.Intent-BeerStatsIntent.IntentSlot-PeopleCategory", Variations: []alexa.PromptVariations{
				{Type: "PlainText", Value: "Für welche Personengruppe möchtest du Bier Statistiken?"},
			}},
		},
	},
}

func runMake(c *cli.Context) error {
	if c.Bool("skill") {
		res, _ := json.Marshal(skill)
		fmt.Println(string(res))
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0644); err != nil {
			log.Fatal(err)
		}

	}

	if c.Bool("models") {
		for l, m := range models {
			res, _ := json.Marshal(m)
			fmt.Println(string(res))
			if err := ioutil.WriteFile("./alexa/interactionModels/custom/"+string(l)+".json", res, 0644); err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
