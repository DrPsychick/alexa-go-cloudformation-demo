// Package loca contains all localization for the skill.
package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// keys of the project.
const (
	ByeBye       string = "byebye"
	GenericTitle string = "Alexa"

	// Intents.
	SaySomething              string = "SaySomething"
	SaySomethingSamples       string = "SaySomething_Samples"
	SaySomethingTitle         string = "SaySomething_Title"
	SaySomethingText          string = "SaySomething_Text"
	SaySomethingSSML          string = "SaySomething_SSML"
	SaySomethingUserTitle     string = "SaySomethingUser_Title"
	SaySomethingUserText      string = "SaySomethingUser_Text"
	SaySomethingUserSSML      string = "SaySomethingUser_SSML"
	DemoIntent                string = "DemoIntent"
	DemoIntentSamples         string = "DemoIntent_Samples"
	DemoIntentTitle           string = "DemoIntent_Title"
	DemoIntentText            string = "DemoIntent_Text"
	DemoIntentSSML            string = "DemoIntent_SSML"
	AWSStatus                 string = "AWSStatus"
	AWSStatusSamples          string = "AWSStatus_Samples"
	AWSStatusTitle            string = "AWSStatus_Title"
	AWSStatusText             string = "AWSStatus_Text"
	AWSStatusSSML             string = "AWSStatus_SSML"
	AWSStatusTextGood         string = "AWSStatus_Text_Good"
	AWSStatusSSMLGood         string = "AWSStatus_SSML_Good"
	AWSStatusAreaSamples      string = "AWSStatus_Area_Samples"
	AWSStatusAreaElicitText   string = "AWSStatus_Area_Elicit_Text"
	AWSStatusAreaElicitSSML   string = "AWSStatus_Area_Elicit_SSML"
	AWSStatusRegionSamples    string = "AWSStatus_Region_Samples"
	AWSStatusRegionElicitText string = "AWSStatus_Region_Elicit_Text"
	AWSStatusRegionElicitSSML string = "AWSStatus_Region_Elicit_SSML"
	AWSStatusAreaConfirmSSML  string = "AWSStatus_Area_Confirm_SSML"
	RegionValidateText        string = "_Region_Validate_Text"

	// Types.
	TypeArea        string = "AWSArea"
	TypeAreaName    string = "Area"
	TypeAreaValues  string = "AWSArea_Values"
	TypeAreaSamples string = "AWSArea_Samples"

	TypeRegion        string = "AWSRegion"
	TypeRegionName    string = "Region"
	TypeRegionValues  string = "AWSRegion_Values"
	TypeRegionSamples string = "AWSRegion_Samples"

	AMAZONStopSamples   string = "AMAZON.StopIntent_Samples"
	AMAZONHelpSamples   string = "AMAZON.HelpIntent_Samples"
	AMAZONCancelSamples string = "AMAZON.CancelIntent_Samples"
)

// Registry is the global l10n registry.
var Registry = l10n.NewRegistry()

func init() {
	// default first
	locales := []*l10n.Locale{
		enUS, deDE, // frFR,
	}
	for _, l := range locales {
		if err := Registry.Register(l); err != nil {
			panic("registration of locale failed")
		}
	}
}
