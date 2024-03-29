package loca

import (
	"github.com/drpsychick/go-alexa-lambda/l10n"
	"github.com/drpsychick/go-alexa-lambda/ssml"
)

var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        {"Voice control demo"},
		l10n.KeySkillDescription: {"Voice demo for the golang meetup"},
		l10n.KeySkillSummary: {
			"This skill demonstrates what you can do with the alexa package and cloudformation",
		},
		l10n.KeySkillSmallIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_small.png", //nolint:lll
		},
		l10n.KeySkillLargeIconURI: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_large.png", //nolint:lll
		},
		l10n.KeySkillInvocation: {"alfalfa demo"},
		l10n.KeySkillTestingInstructions: {
			"Alexa, open alfalfa demo. Yes? Go ahead.",
		},
		l10n.KeySkillPrivacyPolicyURL: {
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		},
		// Error: privacyAndCompliance.locales.en-US
		// - object instance has properties which are not allowed by the schema: ["termsOfUse"]
		// l10n.KeySkillTermsOfUse: []string{
		//	"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		// },
		l10n.KeySkillExamplePhrases: {
			"Alexa, start alfalfa demo and go ahead",
			"How is A.W.S.",
			"Say something",
		},
		l10n.KeySkillKeywords: {
			"demo", "test", "SSML",
		},

		// Errors
		l10n.KeyErrorTitle:               {"Error"},
		l10n.KeyErrorText:                {"The following error occurred:\n%s"},
		l10n.KeyErrorSSML:                {"<speak>An error occurred.</speak>"},
		l10n.KeyErrorLocaleNotFoundTitle: {"Locale missing"},
		l10n.KeyErrorLocaleNotFoundText:  {"Locale for '%s' not found!"},
		l10n.KeyErrorLocaleNotFoundSSML: {
			ssml.Speak("The locale '%s' is not supported."),
		},
		l10n.KeyErrorTranslationTitle: {"Translation missing"},
		l10n.KeyErrorTranslationText:  {"There was an error in translation. The developer is informed."},
		l10n.KeyErrorTranslationSSML: {ssml.Speak(
			"An error occurred during translation. The developer gets informed about this.",
		)},
		l10n.KeyErrorNoTranslationTitle: {"Translation missing"},
		l10n.KeyErrorNoTranslationText:  {"No translation found for '%s'!"},
		l10n.KeyErrorNoTranslationSSML: {
			ssml.Speak("No translation found for '%s'!"),
		},
		l10n.KeyErrorMissingPlaceholderTitle: {"Placeholder missing"},
		l10n.KeyErrorMissingPlaceholderText:  {"Placeholder missing in '%s'!"},
		l10n.KeyErrorMissingPlaceholderSSML:  {ssml.Speak("Placeholder missing in %s!")},

		// Type values
		TypeAreaValues:   {"Europe", "North America", "Asia Pacific", "South America"},
		TypeRegionValues: {"Frankfurt", "Ireland", "London", "Paris", "Stockholm", "North Virginia"},

		// Launch request
		l10n.KeyLaunchTitle: {
			"Greeting",
		},
		l10n.KeyLaunchText: {
			"Hello!",
			"Hi!",
			"Yes?",
		},
		l10n.KeyLaunchSSML: {
			ssml.Speak("<voice name=\"Marlene\">Hello!</voice>"),
			ssml.Speak("<emphasis level=\"strong\">Hi!</emphasis>"),
		},

		// default intents
		l10n.KeyStopTitle:   {"Ending"},
		l10n.KeyStopText:    {"End.", "Good bye.", "See U!"},
		l10n.KeyStopSSML:    {ssml.Speak("Bye."), ssml.Speak("Ok, I'll stop.")},
		l10n.KeyHelpTitle:   {"Help"},
		l10n.KeyHelpText:    {"Try saying 'here we go' or 'go ahead'"},
		l10n.KeyHelpSSML:    {ssml.Speak("Try saying 'here we go' or 'go ahead'")},
		l10n.KeyCancelTitle: {"Abort"},
		l10n.KeyCancelText:  {"Aborting."},
		l10n.KeyCancelSSML:  {ssml.Speak("Alright, aborting.")},

		// Intent: "DemoIntent"
		DemoIntentSamples: {"here we go", "go ahead"},
		DemoIntentTitle:   {"Demo"},
		DemoIntentText:    {"PACE is geil", "you're right"},
		DemoIntentSSML: {
			ssml.Speak(
				ssml.UseVoiceLang("Joanna", "en-US", "<emphasis level=\"strong\">pace</emphasis>") +
					ssml.UseVoiceLang("Kendra", "en-US", " is <emphasis level=\"strong\">geil!</emphasis>"),
			),
			ssml.Speak(ssml.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
		},

		// Intent: "SaySomething"
		SaySomethingSamples: {"say something", "tell me a story"},
		SaySomethingTitle:   {"Get this", "Listen up"},
		SaySomethingText: {
			"Some german words sound nice in english...",
		},
		SaySomethingSSML: {
			ssml.Speak(
				ssml.UseVoiceLang("Kendra", "en-US", "I like the Autobahn, it's so geil"),
			),
		},
		SaySomethingUserTitle: {"Hey %s!"},
		SaySomethingUserText:  {"I like how you dress %s."},
		SaySomethingUserSSML:  {ssml.Speak("I <emphasis level=\"strong\">like</emphasis> your new look %s!")},

		// Intent "AWSStatusIntent"
		AWSStatusSamples: {
			"how is A.W.S.", "how is A.W.S. in {Region}", "how is A.W.S. in {Area} {Region}",
			"tell me the A.W.S. status", "tell me the A.W.S. status in {Area} {Region}",
			"about A.W.S. status in {Area} {Region}",
		},
		AWSStatusTitle: {"AWS Status"},
		AWSStatusText:  {"AWS Status in region %s, %s: okay", "In %s, %s everything's fine"},
		AWSStatusSSML: {
			ssml.Speak("A.W.S. status in %s, %s: all okay"),
		},
		AWSStatusTextGood: {
			"AWS Status in %s: all good",
			"In %s everything's up and running",
		},
		AWSStatusSSMLGood: {
			ssml.Speak("A.W.S. status in %s: everything <emphasis level=\"strong\">perfect</emphasis>"),
			ssml.Speak("In %s everything's running smoothly"),
		},
		AWSStatusAreaSamples: {"in {Area}", "of {Area}", "{Area}"},
		AWSStatusAreaElicitText: {
			"In which area? (Europe, North America, ...)",
			"What area are you interested in? (Europe, North America, ...)",
		},
		AWSStatusAreaElicitSSML: {
			ssml.Speak("In which Area?"), // not working?
			ssml.Speak("About which area do you want to know the status?"),
		},
		AWSStatusAreaConfirmSSML: {
			ssml.Speak("Are you sure about area {Area}?"),
		},
		AWSStatusRegionSamples: {"in {Region}", "of {Region}", "{Region}"},
		AWSStatusRegionElicitText: {
			"In which region? (Frankfurt, North Virginia, ...)",
			"What region are you interested in? (Ireland, Frankfurt, ...)",
		},
		AWSStatusRegionElicitSSML: {
			ssml.Speak("In which Region?"), // not working?
			ssml.Speak("About which region do you want to know the status?"),
		},
		RegionValidateText: {
			"Please choose a valid region like Frankfurt, Ireland, North Virginia.",
		},
		// required for tests to work (delegated to Alexa in real use)
		AMAZONStopSamples:   {"stop", "terminate"},
		AMAZONHelpSamples:   {"help", "help me"},
		AMAZONCancelSamples: {"abort"},
	},
	// IntentResponses: l10n.IntentResponses{
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
	// },
}
