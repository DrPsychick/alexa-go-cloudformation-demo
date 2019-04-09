package loca

import "github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"

var enUS = &l10n.Locale{
	Name:     "en-US",
	Fallback: nil,
	// TODO: move simple text (no list) to separate key?
	TextSnippets: map[l10n.Key][]string{
		l10n.KeySkillName:                []string{"My Skill"},
		l10n.KeySkillDescription:         []string{"This skill demonstrates what you can do with the alexa package"},
		l10n.KeySkillSummary:             []string{"Summary of this skill..."},
		l10n.KeySkillSmallIconURI:        []string{"https://smallicon.uri"},
		l10n.KeySkillLargeIconURI:        []string{"https://largeicon.uri"},
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
		TypeBeerCountriesValues: []string{"Germany", "France"},
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
}
