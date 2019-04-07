package l10n

import "github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"

var enUS = l10n.Locale{
	Name:     "en-US",
	Fallback: nil,
	TextSnippets: map[l10n.Key][]string{
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
