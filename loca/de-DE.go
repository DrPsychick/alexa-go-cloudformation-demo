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
			"hopp hopp",
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
		TypeRegionValues: {"Frankfurt", "Irland", "London", "Paris", "Stockholm", "Nord Virginia"},

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

		// Intent "AMAZON.StopIntent"
		StopTitle: {"Ende Gelände"},
		Stop:      {"Ende."},
		HelpTitle: {"Hilfe"},
		Help:      {"Probier mal 'hopp hopp' oder 'sag etwas' oder 'erzähl mir was'"},

		// Intent: "DemoIntent"
		DemoIntentSamples: {"schiess' los", "auf geht's", "hopp hopp"},
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
		SaySomethingUserTitle: {"Hey %s!"},
		SaySomethingUserText:  {"Mir gefällt dein neues Aussehen, %s."},
		SaySomethingUserSSML:  {l10n.Speak("Mir <emphasis level=\"strong\">gefällt</emphasis> dein neues Aussehen, %s.")},
		// Intent "AWSStatusIntent"
		AWSStatusSamples: {"wie geht's A.W.S.", "sag mir den A.W.S. Status in {Area}, {Region}", "nach dem A.W.S. Status in {Area}, {Region}"},
		AWSStatusTitle:   {"AWS Status"},
		AWSStatusText:    {"AWS Status in %s, %s: okay"},
		AWSStatusSSML: {
			l10n.Speak("A.W.S. Status in Region %s, %s: SNAFU"),
			l10n.Speak("A.W.S. Status in %s, %s: alles ok"),
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
		AWSStatusAreaElicitText: {
			"In welchem Gebiet? (Europa, Nordamerika, ...)",
			"Welches Gebiet interessiert dich? (Europa, Nordamerika, ...)",
		},
		AWSStatusAreaElicitSSML: {
			l10n.Speak("In welchem Gebiet?"),
			l10n.Speak("Zu welchem Gebiet möchtest du den Status wissen?"),
		},
		AWSStatusAreaConfirmSSML: {
			l10n.Speak("Sicher in {Area}?"),
		},
		AWSStatusRegionSamples: {"in {Region}", "der {Region}"},
		AWSStatusRegionElicitText: {
			"In welcher Region? (Frankfurt, Irland, ...)",
			"Zu welcher Region möchtest du den Status wissen? (Frankfurt, Nord Virginia, ...)",
		},
		AWSStatusRegionElicitSSML: {
			l10n.Speak("In welcher Region?"), // not working?
			l10n.Speak("Zu welcher Region möchtest du den Status wissen?"),
		},
		RegionValidateText: {
			"Bitte wähle eine gültige Region, zum Beispiel Frankfurt, Irland, Nord Virginia.",
		},

		// required for tests to work (delegated to Alexa in real use)
		AMAZONStopSamples:   {"stop", "beenden"},
		AMAZONHelpSamples:   {"hilfe", "hilf mir"},
		AMAZONCancelSamples: {"brich ab"},
	},
}
