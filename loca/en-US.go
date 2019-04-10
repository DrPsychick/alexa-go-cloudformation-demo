package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var enUS = &l10n.Locale{
	Name:     "en-US",
	Fallback: nil,
	//Countries: []alexa.Country{},
	// TODO: move simple text (no list) to separate key?
	TextSnippets: map[l10n.Key][]string{
		l10n.KeySkillName:        []string{"Demo Skill"},
		l10n.KeySkillDescription: []string{"Demo for the golang meetup"},
		l10n.KeySkillSummary: []string{
			"This skill demonstrates what you can do with the alexa package and cloudformation",
		},
		l10n.KeySkillSmallIconURI: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_small.png",
		},
		l10n.KeySkillLargeIconURI: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_large.png",
		},
		l10n.KeySkillInvocation:          []string{"demo skill"},
		l10n.KeySkillTestingInstructions: []string{"Alexa, open demo skill. Yes? Say something."},
		l10n.KeySkillPrivacyPolicyURL:    []string{"https://privacy.policy"},
		l10n.KeySkillTermsOfUse:          []string{"https://terms.of.use"},
		l10n.KeySkillExamplePhrases: []string{
			"Alexa, start demo skill and say something",
		},
		l10n.KeySkillKeywords: []string{
			"demo", "test", "SSML", "cloudformation", "automation",
		},
		// Type values
		TypeBeerCountriesValues:  []string{"Germany", "France"},
		TypePeopleCategoryValues: []string{"All", "Women", "Men", "Teenager", "Intellectuals"},
		GreetingTitle: []string{
			"Greeting",
		},
		Greeting: []string{
			"Hello!",
			"Hi!",
		},
		GreetingSSML: []string{
			"<speak><voice name=\"Marlene\">Hello!</voice></speak>",
			"<speak><emphasis level=\"strong\">Hi!</emphasis></speak>",
		},
	},
	IntentResponses: l10n.IntentResponses{
		DemoIntent: l10n.IntentResponse{
			Samples: []string{"schiess' los", "auf geht's", "hop hop"},
			Title:   []string{"Demo"},
			Text: []string{
				"PACE ist geil!",
				"Jawoll",
			},
			SSML: []string{
				"<speak>" +
					"<voice name=\"Kendra\"><lang xml:lang=\"en-US\"><emphasis level=\"strong\">pace</emphasis></lang></voice>" +
					"<voice name=\"Marlene\">iss <emphasis level=\"strong\">geil!</emphasis></voice>" +
					"</speak>",
				"<speak><voice name=\"Kendra\">" +
					"<lang xml:lang=\"en-US\"><emphasis level=\"strong\">geil</emphasis></lang>" +
					"</voice></speak>"},
		},
	},
}
