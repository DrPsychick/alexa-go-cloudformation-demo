package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

var deDE = &l10n.Locale{
	Name:     "de-DE",
	Fallback: enUS,
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
			"Hör' zu!",
		},
		SaySomethingSSML: []string{
			"<speak>" +
				"Sie: <voice name=\"Marlene\">Schatz? Ich fühl mich in letzter Zeit so dick und hässlich, " +
				"ich brauch dringend ein Kompliment!</voice> " +
				"Er: <voice name=\"Hans\">Du hast eine hervorragende Beobachtungsgabe, mein Schatz!</voice>" +
				"</speak>",
			"<speak>" +
				"Er: <voice name=\"Hans\">Wenn meine Frau singt, gehe ich immer aus dem Haus, " +
				"damit die Nachbarn sehen, dass ich sie nicht schlage!</voice>" +
				"</speak>",
		},
	},
	// Utterances: map[loca.Key][]string{ SaySomethingUtterances: []string{}, },
}
