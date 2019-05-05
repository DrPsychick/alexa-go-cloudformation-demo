package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"Demo Skill"},
		l10n.KeySkillDescription: {"Demonstrationsskill für das Meetup"},
		l10n.KeySkillSummary:     {"Dieser Skill demonstriert was man mit dem DrPsychick/alexa package machen kann"},
		l10n.KeySkillInvocation:  {"demo skill"},
		l10n.KeySkillExamplePhrases: {
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
		l10n.KeyErrorTitle:               {"Fehler"},
		l10n.KeyErrorText:                {"Es ist folgender Fehler aufgetreten:\n%s"},
		l10n.KeyErrorSSML:                {"<speak>Es ist ein Fehler aufgetreten.</speak>"},
		l10n.KeyErrorNoTranslationText:   {"Keine Übersetzung für '%s' gefunden!"},

		// Type values
		TypeAreaValues:   {"Europa", "Nordamerika", "Südamerika", "Asien"},
		TypeRegionValues: {"Frankfurt", "Irland", "London", "Paris", "Stockholm", "Nordvirginia"},

		// Launch request
		LaunchTitle: {
			"Begrüßung",
		},
		LaunchText: {
			"Hallo!",
			"Guten Tag!",
		},
		LaunchSSML: {
			"<speak><voice name=\"Marlene\">Hallo!</voice></speak>",
			"<speak>Guten <emphasis level=\"strong\">Tag!</emphasis></speak>",
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
				l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">pace</emphasis>") +
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
		SaySomethingUserTitle: {"Hey %s!"},
		SaySomethingUserText:  {"Mir gefällt dein neues Aussehen, %s."},
		SaySomethingUserSSML:  {l10n.Speak("Mir <emphasis level=\"strong\">gefällt</emphasis> dein neues Aussehen, %s.")},
		// Intent "AWSStatusIntent"
		AWSStatusSamples:          {"wie geht's A.W.S."},
		AWSStatusTitle:            {"AWS Status"},
		AWSStatusText:             {"AWS Status in %s %s"},
		AWSStatusSSML:             {l10n.Speak("Foo")},
		AWSStatusAreaSamples:      {"in {Area}", "von {Area}"},
		AWSStatusRegionSamples:    {"in {Region}", "der {Region}"},
		AWSStatusRegionElicitText: {"In welcher Region?", "Wo nochmal?"},
		AWSStatusRegionElicitSSML: {
			l10n.Speak("In welcher Region?"), l10n.Speak("Wo bitte?")},
		AWSStatusAreaConfirmSSML: {
			l10n.Speak("Sicher in Region?"),
		},
	},
}
