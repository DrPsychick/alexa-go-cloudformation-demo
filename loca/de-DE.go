package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        []string{"Demo Skill"},
		l10n.KeySkillDescription: []string{"Demonstrationsskill für das Meetup"},
		l10n.KeySkillSummary:     []string{"Dieser Skill demonstriert was man mit dem DrPsychick/alexa package machen kann"},
		l10n.KeySkillInvocation:  []string{"demo skill"},
		l10n.KeySkillExamplePhrases: []string{
			"Alexa, starte demo skill und sag etwas",
			"schiess los",
			"hop hop",
		},
		// fallback to enUS
		l10n.KeySkillSmallIconURI:     enUS.GetAll(l10n.KeySkillSmallIconURI),
		l10n.KeySkillLargeIconURI:     enUS.GetAll(l10n.KeySkillLargeIconURI),
		l10n.KeySkillKeywords:         enUS.GetAll(l10n.KeySkillKeywords),
		l10n.KeySkillPrivacyPolicyURL: enUS.GetAll(l10n.KeySkillPrivacyPolicyURL),
		//l10n.KeySkillTermsOfUse:          enUS.GetAll(l10n.KeySkillTermsOfUse),
		l10n.KeySkillTestingInstructions: enUS.GetAll(l10n.KeySkillTestingInstructions),

		// Type values
		TypeAreaValues:   []string{"Europa", "Nordamerika", "Südamerika", "Asien"},
		TypeRegionValues: []string{"Frankfurt", "Irland", "London", "Paris", "Stockholm", "Nordvirginia"},

		// Launch request
		LaunchTitle: []string{
			"Begrüßung",
		},
		LaunchText: []string{
			"Hallo!",
			"Guten Tag!",
		},
		LaunchSSML: []string{
			"<speak><voice name=\"Marlene\">Hallo!</voice></speak>",
			"<speak>Guten <emphasis level=\"strong\">Tag!</emphasis></speak>",
		},

		// Intent: "DemoIntent"
		DemoIntentSamples: []string{"schiess' los", "auf geht's", "hop hop"},
		DemoIntentTitle:   []string{"Demo"},
		DemoIntentText: []string{
			"PACE ist geil!",
			"Jawoll",
		},
		DemoIntentSSML: []string{
			l10n.Speak(
				l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">pace</emphasis>") +
					l10n.UseVoice("Marlene", "iss <emphasis level=\"strong\">geil!</emphasis>"),
			),
			l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
		},
		// Intent: "SaySomething"
		SaySomethingSamples: []string{"sag' etwas", "erzähl' mir was"},
		SaySomethingTitle:   []string{"Antwort", "Titel 2"},
		SaySomethingText: []string{
			"Jetzt sag ich dir mal was... Kannst du das wirklich glauben?",
			"Ich hätte das nie für möglich gehalten!",
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
			"<speak>Ich <emphasis level=\"strong\">grüße</emphasis> dich!</speak>",
		},
		// Intent "AWSStatusIntent"
		AWSStatusSamples:          []string{"wie geht's A.W.S."},
		AWSStatusTitle:            []string{"AWS Status"},
		AWSStatusText:             []string{"AWS Status in {Region}"},
		AWSStatusAreaSamples:      []string{"in {Area}", "von {Area}"},
		AWSStatusRegionSamples:    []string{"in {Region}", "der {Region}"},
		AWSStatusRegionElicitText: []string{"In welcher Region?", "Wo nochmal?"},
		AWSStatusRegionElicitSSML: []string{
			l10n.Speak("In welcher Region?"), l10n.Speak("Wo bitte?")},
		AWSStatusAreaConfirmSSML: []string{
			l10n.Speak("Sicher in Region?"),
		},
	},
}
