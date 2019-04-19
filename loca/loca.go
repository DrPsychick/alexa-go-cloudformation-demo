package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// keys of the project
const (
	GreetingTitle string = "greeting_title"
	Greeting      string = "greeting"
	GreetingSSML  string = "greeting_ssml"
	ByeBye        string = "byebye"
	StopTitle     string = "stop_title"
	Stop          string = "stop"
	GenericTitle  string = "Alexa"

	// Intents
	SaySomething        string = "SaySomething"
	SaySomethingSamples string = "SaySomething_Samples"
	SaySomethingTitle   string = "SaySomething_Title"
	SaySomethingText    string = "SaySomething_Text"
	SaySomethingSSML    string = "SaySomething_SSML"
	DemoIntent          string = "DemoIntent"
	DemoIntentText      string = "DemoIntentText"
	DemoIntentSSML      string = "DemoIntentSSML"

	AWSStatusIntent string = "AWSStatus"

	// Types
	TypeArea       string = "AWS_Area"
	TypeAreaName   string = "Area"
	TypeAreaValues string = "AWS_AreaValues"

	TypeRegion       string = "AWS_Region"
	TypeRegionName   string = "Region"
	TypeRegionValues string = "AWS_RegionValues"
)

func init() {
	var locales = []*l10n.Locale{
		deDE, enUS, frFR,
	}
	for _, l := range locales {
		if err := l10n.Register(l); err != nil {
			panic("registration of locale failed")
		}
	}
}
