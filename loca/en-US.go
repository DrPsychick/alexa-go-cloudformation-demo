package loca

import "github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"

var enUS = &l10n.Locale{
	Name:     "en-US",
	Fallback: nil,
	// TODO: move simple text (no list) to separate key?
	TextSnippets: map[l10n.Key][]string{
		l10n.KeySkillName:         []string{"My Skill"},
		l10n.KeySkillDescription:  []string{"This skill demonstrates what you can do with the alexa package"},
		l10n.KeySkillSmallIconURI: []string{"https://smallicon.uri"},
		l10n.KeySkillLargeIconURI: []string{"https://largeicon.uri"},
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
