package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"Voice control demo"},
		l10n.KeySkillDescription: {"Demonstrationsskill für das Meetup"},
		l10n.KeySkillSummary:     {"Dieser Skill demonstriert was man mit dem DrPsychick/alexa package machen kann"},
		l10n.KeySkillInvocation:  {"voice demo"},
		l10n.KeySkillExamplePhrases: {
			"Alexa, starte voice demo und sag etwas",
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
		TypeAreaValues:   {"Europa", "Nordamerika", "Südamerika", "Asien"},
		TypeRegionValues: {"Frankfurt", "Irland", "London", "Paris", "Stockholm", "Nordvirginia"},

		// Launch request
		LaunchTitle: {
			"Begrüßung",
			"Willkommen",
		},
		LaunchText: {
			"Hallo!",
			"Guten Tag!",
			"Willkommen bei der Voice Demo!",
		},
		LaunchSSML: {
			l10n.Speak(l10n.UseVoice("Marlene", "Hallo!")),
			l10n.Speak("Guten <emphasis level=\"strong\">Tag!</emphasis>"),
			l10n.Speak(l10n.UseVoice("Marlene", "Willkommen bei der <emphasis level=\"strong\">Voice</emphasis> Demo!")),
		},

		// Intent: "DemoIntent"
		DemoIntentSamples: {"schiess' los", "auf geht's", "hop hop"},
		DemoIntentTitle:   {"Demo"},
		DemoIntentText: {
			"PACE ist geil!",
			"Jawoll",
		},
		DemoIntentSSML: {
			l10n.Speak(
				l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">pace</emphasis> ") +
					l10n.UseVoice("Marlene", "iss <emphasis level=\"strong\">geil!</emphasis>"),
			),
			l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
		},
		// Intent: "SaySomething"
		SaySomethingSamples: {"sag' etwas", "erzähl' mir was"},
		SaySomethingTitle:   {"Antwort", "Titel 2"},
		SaySomethingText: {
			"Jetzt sag ich dir mal was... Kannst du das wirklich glauben?",
			"Ich hätte das nie für möglich gehalten!",
			"Hör' zu!",
		},
		SaySomethingSSML: {
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
		AWSStatusSamples: {"wie geht's A.W.S.", "A.W.S. Status in {Region}", "A.W.S. Status in {Area}"},
		AWSStatusTitle:   {"AWS Status"},
		AWSStatusText:    {"AWS Status in %s: okay"},
		AWSStatusSSML: {
			l10n.Speak("A.W.S. Status in Region %s: SNAFU"),
			l10n.Speak("A.W.S. Status in %s: alles ok"),
		},
		AWSStatusTextGood: {
			"AWS Status in %s: alles bestens",
			"In %s läuft alles rund",
		},
		AWSStatusSSMLGood: {
			l10n.Speak("A.W.S. Status in %s: alles <emphasis level=\"strong\">super</emphasis>"),
			l10n.Speak("In %s: alles " + l10n.UseVoiceLang("Kendra", "en-US", "geil")),
		},
		AWSStatusAreaSamples: {"in {Area}", "von {Area}"},
		AWSStatusAreaConfirmSSML: {
			l10n.Speak("Sicher in {Area}?"),
		},
		AWSStatusRegionSamples: {"in {Region}", "der {Region}"},
		AWSStatusRegionElicitText: {
			"In welcher Region? (Europa, Nordamerika, ...)",
			"Zu welcher Region möchtest du den Status wissen? (Europa, Nordamerika, ...)",
		},
		AWSStatusRegionElicitSSML: {
			l10n.Speak("In welcher Region?"),
			l10n.Speak("Zu welcher Region möchtest du den Status wissen?"),
		},

		// Intent "AMAZON.StopIntent"
		StopTitle: {"Ende Gelände"},
		Stop:      {"Ende."},
		HelpTitle: {"Hilfe"},
		Help:      {"Probier mal 'Hop hop' oder 'sag etwas' oder 'erzähl mir was'"},
	},
}
