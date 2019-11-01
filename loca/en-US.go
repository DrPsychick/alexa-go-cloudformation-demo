package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"Demo Skill"},
		l10n.KeySkillDescription: {"Demo for the golang meetup"},
		l10n.KeySkillSummary: {
			"This skill demonstrates what you can do with the alexa package and cloudformation",
		},
		l10n.KeySkillSmallIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_small.png",
		},
		l10n.KeySkillLargeIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_large.png",
		},
		l10n.KeySkillInvocation: {"demo skill"},
		l10n.KeySkillTestingInstructions: {
			"Alexa, open demo skill. Yes? Go ahead.",
		},
		l10n.KeySkillPrivacyPolicyURL: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		},
		// Error: privacyAndCompliance.locales.en-US - object instance has properties which are not allowed by the schema: ["termsOfUse"]
		//l10n.KeySkillTermsOfUse: []string{
		//	"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		//},
		l10n.KeySkillExamplePhrases: {
			"Alexa, start demo skill and go ahead",
			"Here we go",
			"How is AWS",
		},
		l10n.KeySkillKeywords: {
			"demo", "test", "SSML",
		},
		// Type values
		TypeAreaValues:   {"Europe", "North America", "Asia Pacific", "South America"},
		TypeRegionValues: {"Frankfurt", "Ireland", "London", "Paris", "Stockholm", "North Virginia"},

		// Launch request
		LaunchTitle: {
			"Greeting",
		},
		LaunchText: {
			"Hello!",
			"Hi!",
		},
		LaunchSSML: {
			"<speak><voice name=\"Marlene\">Hello!</voice></speak>",
			"<speak><emphasis level=\"strong\">Hi!</emphasis></speak>",
		},

		// default intents
		StopTitle: {"Ending"},
		Stop:      {"End."},
		HelpTitle: {"Help"},
		Help:      {"Try saying 'here we go' or 'go ahead'"},

		// Intent: "DemoIntent"
		DemoIntentSamples: {"here we go", "go ahead"},
		DemoIntentTitle:   {"Demo"},
		DemoIntentText:    {"PACE is geil", "you're right"},
		DemoIntentSSML: {
			l10n.Speak(
				l10n.UseVoiceLang("Joanna", "en-US", "<emphasis level=\"strong\">pace</emphasis>") +
					l10n.UseVoiceLang("Kendra", "en-US", "is <emphasis level=\"strong\">geil!</emphasis>"),
			),
			l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
		},

		// Intent: "SaySomething"
		SaySomethingSamples: {"say something", "tell me a story"},
		SaySomethingTitle:   {"Get this", "Listen up"},
		SaySomethingText: {
			"Some german words sound nice in english...",
		},
		SaySomethingSSML: {
			l10n.Speak(
				l10n.UseVoiceLang("Kendra", "en-US", "I like the Autobahn, it's so geil"),
			),
		},

		// Intent "AWSStatusIntent"
		AWSStatusSamples:          {"how is A.W.S."},
		AWSStatusTitle:            {"AWS Status"},
		AWSStatusText:             {"AWS Status in {Region}"},
		AWSStatusAreaSamples:      {"in {Area}", "of {Area}"},
		AWSStatusRegionSamples:    {"in {Region}", "of {Region}"},
		AWSStatusRegionElicitText: {"In which region?", "Where again?"},
		AWSStatusRegionElicitSSML: {
			l10n.Speak("In which Region?"), l10n.Speak("Sorry, where?")},
		AWSStatusAreaConfirmSSML: {
			l10n.Speak("Are you sure?"),
		},
	},
	//IntentResponses: l10n.IntentResponses{
	//	SaySomething: l10n.IntentResponse{
	//		Samples: []string{"say something", "tell me a story"},
	//		Title:   []string{"Answer", "Title 2"},
	//		Text: []string{
	//			"I will tell you something... Can you believe it?",
	//			"I never thought this would be possible!",
	//			"Listen!",
	//		},
	//		SSML: []string{
	//			l10n.Speak(
	//				"Sie: <voice name=\"Joanna\">Schatz? Ich fühl mich in letzter Zeit so dick und hässlich, " +
	//					"ich brauch dringend ein Kompliment!</voice> " +
	//					"Er: <voice name=\"Matthew\">Du hast eine hervorragende Beobachtungsgabe, mein Schatz!</voice>",
	//			),
	//			l10n.Speak("He:" +
	//				l10n.UseVoice("Matthew",
	//					"When my wife sings, I always leave the house, "+
	//						"so that my neighbours see that I don't beat her!"),
	//			),
	//			l10n.Speak("I <emphasis level=\"strong\">great</emphasis> you!"),
	//		},
	//	},
	//	AWSStatusIntent: l10n.IntentResponse{
	//		Title: []string{},
	//		Samples: []string{
	//			"A.W.S. status of {Area}",
	//			"status of {Area}",
	//			"give me the status of {Region}",
	//			"status of {Region}",
	//			"{Region} status",
	//		},
	//		Text: []string{},
	//		Slots: map[string]l10n.Slot{
	//			TypeAreaName: {
	//				Samples: []string{"of {Area}", "in {Area}"},
	//				PromptElicitations: []alexa.PromptVariations{
	//					{Type: "PlainText", Value: "From what area do you seek status?"},
	//					{Type: "PlainText", Value: "Which area do you want to get A.W.S. stats for?"},
	//					{Type: "SSML", Value: l10n.Speak("Which area?")},
	//				},
	//			},
	//			TypeRegionName: {
	//				Samples: []string{"of {Region}", "in {Region}"},
	//				PromptElicitations: []alexa.PromptVariations{
	//					{Type: "PlainText", Value: "From what region do you want to know the status?"},
	//				},
	//			},
	//		},
	//	},
	//},
}
