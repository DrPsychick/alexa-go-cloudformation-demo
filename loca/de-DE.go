package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

var deDE = &l10n.Locale{
	Name:     "de-DE",
	Fallback: enUS,
	Countries: []alexa.Country{
		"DE", "AT",
	},
	TextSnippets: map[l10n.Key][]string{
		l10n.KeySkillName:         []string{"Mein Skill"},
		l10n.KeySkillDescription:  []string{"Demonstrationsskill"},
		l10n.KeySkillSummary:      []string{"Dieser Skill demonstriert was man mit dem DrPsychick/alexa package machen kann"},
		l10n.KeySkillSmallIconURI: []string{"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_small.png"},
		l10n.KeySkillLargeIconURI: []string{"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/development/alexa/assets/images/de-DE_large.png"},
		l10n.KeySkillInvocation:   []string{"demo skill"},
		l10n.KeySkillExamplePhrases: []string{
			"Alexa, starte demo skill und sag etwas",
			"schiess los",
			"sag' was",
		},
		TypeBeerCountriesValues: []string{"Deutschland", "Frankreich"},
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
	},
	// Combine localization used in a single response
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
		SaySomething: l10n.IntentResponse{
			Samples: []string{"sag' etwas", "erzähl' mir was"},
			Title:   []string{"Antwort", "Titel 2"},
			Text: []string{
				"Jetzt sag ich dir mal was... Kannst du das wirklich glauben?",
				"Ich hätte das nie für möglich gehalten!",
				"Hör' zu!",
			},
			SSML: []string{
				"<speak>" +
					"Sie: <voice name=\"Marlene\">Schatz? Ich fühl mich in letzter Zeit so dick und hässlich, " +
					"ich brauch dringend ein Kompliment!</voice> " +
					"Er: <voice name=\"Hans\">Du hast eine hervorragende Beobachtungsgabe, mein Schatz!</voice>" +
					"</speak>",
				"<speak>" +
					"Er: <voice name=\"Hans\">Wenn meine Frau singt, gehe ich immer aus dem Haus, " +
					"damit die Nachbarn sehen, dass ich sie nicht schlage!</voice>" +
					"</speak>",
				"<speak>Ich <emphasis level=\"strong\">grüße</emphasis> dich!</speak>",
			},
		},
	},
}
