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
			Countries:           []alexa.Country{"DE"},
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
		Permissions: &[]alexa.Permission{},
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
				{Name: "SSMLDemoIntent", Samples: []string{
					"Zeig' was du kannst",
					"Immer her damit",
					"Was kann SSML",
				}},
				{Name: "SaySomething", Samples: []string{
					"Erzähl' mir was",
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
					{Name: alexa.NameValue{Value: "Frankreich"}},
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
	sk, _ := createSkill(*l10n.DefaultRegistry)
	ms, _ := createModels(sk)

	if c.Bool("skill") {
		res, _ := json.MarshalIndent(skill, "", "  ")
		//fmt.Println(string(res))
		if err := ioutil.WriteFile("./alexa/skill.json", res, 0644); err != nil {
			log.Fatal(err)
		}
		res, _ = json.MarshalIndent(sk.Build(), "", "  ")
		//fmt.Printf(string(res))
		ioutil.WriteFile(("./newskill.json"), res, 0644)
		fmt.Printf("vimdiff ./newskill.json ./alexa/skill.json\n")
	}

	if c.Bool("models") {
		for l, m := range models {
			res, _ := json.MarshalIndent(m, "", "  ")
			//fmt.Println(string(res))
			if err := ioutil.WriteFile("./alexa/interactionModels/custom/"+string(l)+".json", res, 0644); err != nil {
				log.Fatal(err)
			}
		}
		for l, m := range ms {
			res, _ := json.MarshalIndent(m, "", "  ")
			//fmt.Printf(string(res))
			ioutil.WriteFile("./new"+string(l)+".json", res, 0644)
			fmt.Printf("vimdiff ./new" + string(l) + ".json ./alexa/interactionModels/custom/" + string(l) + ".json\n")
		}
	}

	return nil
}

// createSkill generates and returns a Skill.
func createSkill(r l10n.Registry) (*gen.Skill, error) {
	skill := gen.NewSkill()
	skill.SetCategory(alexa.CategoryOrganizersAndAssistants)
	skill.SetDefaultLocale(r.GetDefaultLocale())
	skill.AddIntentString(string(loca.DemoIntent))
	skill.AddIntentString(string(loca.SaySomething))

	for n, l := range r.GetLocales() {
		skill.AddLocale(n, l)
		for _, c := range l.Countries {
			skill.AddCountry(c)
		}
	}

	// Types will automatically add the values from l10n.Key
	skill.AddTypeString(string(loca.TypeBeerCountries))
	skill.AddTypeString(string(loca.TypePeopleCategory))

	skill.Privacy.SetIsExportCompliant(true)
	//skill.Privacy.SetContainsAds(false)

	return skill, nil
}

// createModels generates and returns a list of Models.
func createModels(s *gen.Skill) (map[string]*alexa.Model, error) {
	return s.BuildModels(), nil
}
