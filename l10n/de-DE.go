package l10n

import (
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

var deDE = l10n.Locale{
	Name:     "de-DE",
	Fallback: &enUS,
	TextSnippets: map[l10n.Key][]string{
		GreetingTitle: []string{
			"Begrüßung",
		},
		Greeting: []string{
			"Hallo!",
			"Guten Tag!",
		},
		GreetingSSML: []string{
			"<speak><voice name=\"Marlene\">Hallo!</voice></speak>",
			"<speak>Guten <emphasis level=\"strong\">Tag!</emphasis></speak>",
		},
		SaySomethingTitle: []string{
			"Antwort",
		},
		SaySomething: []string{
			"Jetzt sag ich dir mal was... Kannst du das wirklich glauben?" +
				"Ich hätte es nie für möglich gehalten!",
		},
	},
}
