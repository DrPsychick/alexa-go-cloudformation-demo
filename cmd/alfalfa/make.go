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
			Locales: map[string]alexa.LocaleDef{
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

var modelGerman alexa.Model = alexa.Model{
	Model: alexa.InteractionModel{
		Language: alexa.Language{
			Invocation: "Bier fakten",
			Intents: []alexa.ModelIntent{
				{Name: "AMAZON.CancelIntent", Samples: []string{}},
			},
			Types: []alexa.ModelType{
				{Name: "BEER_Countries", Values: []alexa.TypeValue{
					{Name: alexa.NameValue{Value: "foo"}},
				}},
			},
		},
	},
}

func runMake(c *cli.Context) error {
	res, _ := json.Marshal(skill)
	fmt.Println(string(res))

	res, _ = json.Marshal(modelGerman)
	fmt.Println(string(res))
	return nil
}
