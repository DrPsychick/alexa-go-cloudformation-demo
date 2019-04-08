package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

var deDE = &l10n.Locale{
	Name:       "de-DE",
	Invocation: "meine demo",
	Fallback:   enUS,
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
	// Combine localization used in a single response
	IntentResponses: l10n.IntentResponses{
		SaySomethingIntent: l10n.IntentResponse{
			Title: []string{"title1", "title2"},
			Text:  []string{"text one", "text two", "text three"},
			SSML:  []string{}, // no SSML defined
		},
	},
	// Utterances: map[l10n.Key][]string{ SaySomethingUtterances: []string{}, },
}

func init() {
	deDE.IntentResponses[SaySomethingIntent] = l10n.IntentResponse{Title: []string{"Foo"}}
	foo := deDE.IntentResponses[SaySomethingIntent]
	foo.Title = []string{"text 1", "text 2"}

	deDE.IntentResponses[Greeting] = l10n.IntentResponse{
		Title: []string{"Hi"},
		Text:  []string{"Hallo"},
	}
}
