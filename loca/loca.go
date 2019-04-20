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
	SaySomething               string = "SaySomething"
	SaySomethingSamples        string = "SaySomething_Samples"
	SaySomethingTitle          string = "SaySomething_Title"
	SaySomethingText           string = "SaySomething_Text"
	SaySomethingSSML           string = "SaySomething_SSML"
	DemoIntent                 string = "DemoIntent"
	DemoIntentSamples          string = "DemoIntent_Samples"
	DemoIntentTitle            string = "DemoIntent_Title"
	DemoIntentText             string = "DemoIntentText"
	DemoIntentSSML             string = "DemoIntentSSML"
	AWSStatus                  string = "AWSStatus"
	AWSStatusSamples           string = "AWSStatus_Samples"
	AWSStatusTitle             string = "AWSStatus_Title"
	AWSStatusText              string = "AWSStatus_Text"
	AWSStatusSSML              string = "AWSStatus_SSML"
	AWSStatusTypeAreaSamples   string = "AWSStatus_AWSArea_Samples"
	AWSStatusTypeRegionSamples string = "AWSStatus_AWSRegion_Samples"
	AWSStatusRegionElicitText  string = "AWSStatus_Region_Elicit_Text"
	AWSStatusRegionElicitSSML  string = "AWSStatus_Region_Elicit_SSML"
	AWSStatusAreaConfirmSSML   string = "AWSStatus_Area_Confirm_SSML"

	// Types
	TypeArea        string = "AWSArea"
	TypeAreaName    string = "Area"
	TypeAreaValues  string = "AWSArea_Values"
	TypeAreaSamples string = "AWSArea_Samples"

	TypeRegion        string = "AWSRegion"
	TypeRegionName    string = "Region"
	TypeRegionValues  string = "AWSRegion_Values"
	TypeRegionSamples string = "AWSRegion_Samples"
)

func init() {
	// default first
	var locales = []*l10n.Locale{
		enUS, deDE, //frFR,
	}
	for _, l := range locales {
		if err := l10n.Register(l); err != nil {
			panic("registration of locale failed")
		}
	}
}
