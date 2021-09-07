package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/ssml"
)

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"Voice control demo"},
		l10n.KeySkillDescription: {"Demonstrationsskill für das Meetup"},
		l10n.KeySkillSummary:     {"Dieser Skill demonstriert was man mit dem DrPsychick/alexa package machen kann"},
		l10n.KeySkillInvocation:  {"alfalfa demo"},
		l10n.KeySkillExamplePhrases: {
			"Alexa, starte alfalfa demo und sag etwas",
			"schiess los",
			"hopp hopp",
		},
		// fallback to enUS
		l10n.KeySkillSmallIconURI:     enUS.GetAll(l10n.KeySkillSmallIconURI),
		l10n.KeySkillLargeIconURI:     enUS.GetAll(l10n.KeySkillLargeIconURI),
		l10n.KeySkillKeywords:         enUS.GetAll(l10n.KeySkillKeywords),
		l10n.KeySkillPrivacyPolicyURL: enUS.GetAll(l10n.KeySkillPrivacyPolicyURL),
		// l10n.KeySkillTermsOfUse:          enUS.GetAll(l10n.KeySkillTermsOfUse),
		l10n.KeySkillTestingInstructions: enUS.GetAll(l10n.KeySkillTestingInstructions),

		// Errors
		l10n.KeyErrorTitle:               {"Fehler"},
		l10n.KeyErrorText:                {"Es ist folgender Fehler aufgetreten:\n%s"},
		l10n.KeyErrorSSML:                {"<speak>Es ist ein Fehler aufgetreten.</speak>"},
		l10n.KeyErrorLocaleNotFoundTitle: {"Sprache fehlt"},
		l10n.KeyErrorLocaleNotFoundText:  {"Sprache für '%s' nicht gefunden!"},
		l10n.KeyErrorLocaleNotFoundSSML: {
			ssml.Speak("Die Sprache '%s' wird nicht unterstützt."),
		},
		l10n.KeyErrorNoTranslationTitle: {"Übersetzung fehlt"},
		l10n.KeyErrorNoTranslationText:  {"Keine Übersetzung für '%s' gefunden!"},
		l10n.KeyErrorNoTranslationSSML: {
			ssml.Speak("Keine Übersetzung für '%s' gefunden!"),
		},

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
			ssml.Speak(ssml.UseVoice("Marlene", "Hallo!")),
			ssml.Speak("Guten <emphasis level=\"strong\">Tag!</emphasis>"),
			ssml.Speak(ssml.UseVoice("Marlene", "Willkommen bei der <emphasis level=\"strong\">Voice</emphasis> Demo!")),
		},

		// Intent "AMAZON.StopIntent"
		StopTitle: {"Ende Gelände"},
		Stop:      {"Ende.", "Tschüss.", "Bis bald."},
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
			ssml.Speak(
				ssml.UseVoiceLang(ssml.USVoiceKendra, "en-US", "<emphasis level=\"strong\">pace</emphasis> ") +
					ssml.UseVoice(ssml.USVoiceSalli, "iss <emphasis level=\"strong\">geil!</emphasis>"),
			),
			ssml.Speak(ssml.UseVoiceLang(ssml.USVoiceKendra, "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
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
		SaySomethingUserSSML:  {ssml.Speak("Mir <emphasis level=\"strong\">gefällt</emphasis> dein neues Aussehen, %s.")},
		// Intent "AWSStatusIntent"
		AWSStatusSamples: {
			"wie geht's A.W.S.", "sag mir den A.W.S. Status in {Area} {Region}",
			"nach dem A.W.S. Status in {Area} {Region}",
		},
		AWSStatusTitle: {"AWS Status"},
		AWSStatusText:  {"AWS Status in %s, %s: okay"},
		AWSStatusSSML: {
			ssml.Speak("A.W.S. Status in Region %s, %s: SNAFU"),
			ssml.Speak("A.W.S. Status in %s, %s: alles ok"),
		},
		AWSStatusTextGood: {
			"AWS Status in %s: alles bestens",
			"In %s läuft alles rund",
		},
		AWSStatusSSMLGood: {
			ssml.Speak("A.W.S. Status in %s: alles <emphasis level=\"strong\">super</emphasis>"),
			ssml.Speak("In %s: alles " + ssml.UseVoiceLang("Kendra", "en-US", "geil")),
		},
		AWSStatusAreaSamples: {"in {Area}", "von {Area}"},
		AWSStatusAreaElicitText: {
			"In welchem Gebiet? (Europa, Nordamerika, ...)",
			"Welches Gebiet interessiert dich? (Europa, Nordamerika, ...)",
		},
		AWSStatusAreaElicitSSML: {
			ssml.Speak("In welchem Gebiet?"),
			ssml.Speak("Zu welchem Gebiet möchtest du den Status wissen?"),
		},
		AWSStatusAreaConfirmSSML: {
			ssml.Speak("Sicher in {Area}?"),
		},
		AWSStatusRegionSamples: {"in {Region}", "der {Region}"},
		AWSStatusRegionElicitText: {
			"In welcher Region? (Frankfurt, Irland, ...)",
			"Zu welcher Region möchtest du den Status wissen? (Frankfurt, Nord Virginia, ...)",
		},
		AWSStatusRegionElicitSSML: {
			ssml.Speak("In welcher Region?"), // not working?
			ssml.Speak("Zu welcher Region möchtest du den Status wissen?"),
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
