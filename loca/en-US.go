package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

var enUS = &l10n.Locale{
	Name:     "en-US",
	Fallback: nil,
	Countries: []alexa.Country{
		alexa.CountryUnitedStates,
	},
	// TODO: move simple text (no list) to separate key?
	TextSnippets: map[string][]string{
		l10n.KeySkillName:        []string{"Demo Skill"},
		l10n.KeySkillDescription: []string{"Demo for the golang meetup"},
		l10n.KeySkillSummary: []string{
			"This skill demonstrates what you can do with the alexa package and cloudformation",
		},
		l10n.KeySkillSmallIconURI: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_small.png",
		},
		l10n.KeySkillLargeIconURI: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/alexa/assets/images/de-DE_large.png",
		},
		l10n.KeySkillInvocation:          []string{"demo skill"},
		l10n.KeySkillTestingInstructions: []string{"Alexa, open demo skill. Yes? Say something."},
		l10n.KeySkillPrivacyPolicyURL: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		},
		l10n.KeySkillTermsOfUse: []string{
			"https://raw.githubusercontent.com/DrPsychick/alexa-go-cloudformation-demo/master/LICENSE",
		},
		l10n.KeySkillExamplePhrases: []string{
			"Alexa, start demo skill and say something",
		},
		l10n.KeySkillKeywords: []string{
			"demo", "test", "SSML", "cloudformation", "automation",
		},
		// Type values
		TypeAreaValues:   []string{"Europe", "North America", "Asia Pacific", "South America"},
		TypeRegionValues: []string{"Frankfurt", "Ireland", "London", "Paris", "Stockholm", "North Virginia"},

		GreetingTitle: []string{
			"Greeting",
		},
		Greeting: []string{
			"Hello!",
			"Hi!",
		},
		GreetingSSML: []string{
			"<speak><voice name=\"Marlene\">Hello!</voice></speak>",
			"<speak><emphasis level=\"strong\">Hi!</emphasis></speak>",
		},
	},
	IntentResponses: l10n.IntentResponses{
		DemoIntent: l10n.IntentResponse{
			Samples: []string{"schiess' los", "auf geht's", "hop hop"},
			Title:   []string{"Demo"},
			Text: []string{
				"PACE ist geil!",
				"Jawoll",
			},
			SSML: []string{
				l10n.Speak(
					l10n.UseVoiceLang("Joanna", "en-US", "<emphasis level=\"strong\">pace</emphasis>") +
						l10n.UseVoiceLang("Kendra", "en-US", "iss <emphasis level=\"strong\">geil!</emphasis>"),
				),
				l10n.Speak(l10n.UseVoiceLang("Kendra", "en-US", "<emphasis level=\"strong\">geil</emphasis>")),
			},
		},
		SaySomething: l10n.IntentResponse{
			Samples: []string{"say something", "tell me a story"},
			Title:   []string{"Answer", "Title 2"},
			Text: []string{
				"I will tell you something... Can you believe it?",
				"I never thought this would be possible!",
				"Listen!",
			},
			SSML: []string{
				l10n.Speak(
					"Sie: <voice name=\"Joanna\">Schatz? Ich fühl mich in letzter Zeit so dick und hässlich, " +
						"ich brauch dringend ein Kompliment!</voice> " +
						"Er: <voice name=\"Matthew\">Du hast eine hervorragende Beobachtungsgabe, mein Schatz!</voice>",
				),
				l10n.Speak("He:" +
					l10n.UseVoice("Matthew",
						"When my wife sings, I always leave the house, "+
							"so that my neighbours see that I don't beat her!"),
				),
				l10n.Speak("I <emphasis level=\"strong\">great</emphasis> you!"),
			},
		},
		AWSStatusIntent: l10n.IntentResponse{
			Title: []string{},
			Samples: []string{
				"A.W.S. status of {Area}",
				"status of {Area}",
				"give me the status of {Region}",
				"status of {Region}",
				"{Region} status",
			},
			Text: []string{},
			Slots: map[string]l10n.Slot{
				TypeAreaName: {
					Samples: []string{"of {Area}", "in {Area}"},
					PromptElicitations: []alexa.PromptVariations{
						{Type: "PlainText", Value: "From what area do you seek status?"},
						{Type: "PlainText", Value: "Which area do you want to get A.W.S. stats for?"},
						{Type: "SSML", Value: l10n.Speak("Which area?")},
					},
				},
				TypeRegionName: {
					Samples: []string{"of {Region}", "in {Region}"},
					PromptElicitations: []alexa.PromptVariations{
						{Type: "PlainText", Value: "From what region do you want to know the status?"},
					},
				},
			},
		},
	},
}
