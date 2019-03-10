package main

import (
	"encoding/json"
	"fmt"
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"gopkg.in/urfave/cli.v1"
)

// Alexa skill definition (to generate skill.json)
var skill = alexa.Skill{
	Manifest: alexa.Manifest{
		Version: "1.0",
		Publishing: alexa.Publishing{
			Locales: map[alexa.Locale]alexa.LocaleDef{
				"de-DE": {
					Name:        "name",
					Description: "description",
					Summary:     "summary",
					Keywords:    []string{"Demo"},
					Examples:    []string{"tell me how much beer people drink in germany"},
				},
			},
			Category:  "mycategory",
			Countries: []alexa.Country{"DE"},
		},
		//Apis: alexa.Apis{
		//	Custom: alexa.Custom{
		//		Endpoint: alexa.Endpoint{
		//			Uri: "arn:...",
		//		},
		//	},
		//	Interfaces: []string{},
		//},
		Permissions: []string{},
		Privacy: alexa.Privacy{
			Compliant: true,
		},
	},
}

var models = []alexa.Model{
	modelGerman,
}

var modelGerman = alexa.Model{
	Model: alexa.InteractionModel{
		Language: alexa.LanguageModel{
			Invocation: "Bier fakten",
			Intents: []alexa.ModelIntent{
				{Name: "AMAZON.CancelIntent", Samples: []string{}},
				{Name: "CustomIntent", Samples: []string{
					"Schiess' los",
					"Auf geht's",
					"Hop hop",
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
					{Name: alexa.NameValue{Value: "Frauen"}},
					{Name: alexa.NameValue{Value: "Männern"}},
					{Name: alexa.NameValue{Value: "Transen"}},
					{Name: alexa.NameValue{Value: "Homosexuellen"}},
					{Name: alexa.NameValue{Value: "Intellektuellen"}},
				}},
			},
		},
		Dialog: &alexa.Dialog{
			Intents: []alexa.DialogIntent{
				{Name: "BeerStatsIntent", Confirmation: false},
			},
		},
		Prompts: &[]alexa.ModelPrompt{},
	},
}

func runMake(c *cli.Context) error {
	if c.Bool("skill") {
		res, _ := json.Marshal(skill)
		fmt.Println(string(res))
	}

	if c.Bool("models") {
		for m := range models {
			res, _ := json.Marshal(models[m])
			fmt.Println(string(res))
		}
	}
	return nil
}
