package alexa

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var minimalSkillDef = Skill{
	Manifest: Manifest{
		Version: "1.0",
		Publishing: Publishing{
			Locales: map[Locale]LocaleDef{
				"de-DE": {
					Name:        "name",
					Description: "description",
					Summary:     "summary",
					Keywords:    []string{"Demo"},
					Examples:    []string{"tell me how much beer people drink in germany"},
				},
			},
			Category:  "mycategory",
			Countries: []Country{"DE"},
		},
		Apis: &Apis{
			Custom: &Custom{
				Endpoint: &Endpoint{
					Uri: "arn:...",
				},
			},
			Interfaces: &[]string{},
		},
		Permissions: &[]Permission{},
		Privacy: &Privacy{
			IsExportCompliant: true,
		},
	},
}

// https://developer.amazon.com/de/docs/smapi/skill-manifest.html#sample-manifests
var awsDocsExample = []byte(`{
"manifest": {
"publishingInformation": {
"locales": {
"en-US": {
"summary": "This is a sample Alexa custom skill.",
"examplePhrases": [
"Alexa, open sample custom skill.",
"Alexa, play sample custom skill."
],
"keywords": [
"Descriptive_Phrase_1",
"Descriptive_Phrase_2",
"Descriptive_Phrase_3"
],
"smallIconUri": "https://smallUri.com",
"largeIconUri": "https://largeUri.com",
"name": "Sample custom skill name.",
"description": "This skill does interesting things."
}
},
"isAvailableWorldwide": false,
"testingInstructions": "1) Say 'Alexa, hello world'",
"category": "HEALTH_AND_FITNESS",
"distributionCountries": [
"US",
"GB",
"DE"
]
},
"apis": {
"custom": {
"endpoint": {
"uri": "arn:aws:lambda:us-east-1:040623927470:function:sampleSkill"
},
"interfaces": [
{
"type":"ALEXA_PRESENTATION_APL"
},
{
"type":"AUDIO_PLAYER"
},
{
"type":"CAN_FULFILL_INTENT_REQUEST"
},
{
"type":"GADGET_CONTROLLER"
},
{
"type":"GAME_ENGINE"
},
{
"type":"RENDER_TEMPLATE"
},
{
"type":"VIDEO_APP"
}
],
"regions": {
"NA": {
"endpoint": {
"sslCertificateType": "Trusted",
"uri": "https://customapi.sampleskill.com"
}
}
}
}
},
"manifestVersion": "1.0",
"permissions": [
{
"name": "alexa::devices:all:address:full:read"
},
{
"name": "alexa:devices:all:address:country_and_postal_code:read"
},
{
"name": "alexa::household:lists:read"
},
{
"name": "alexa::household:lists:write"
},
{
"name": "alexa::alerts:reminders:skill:readwrite"
}
],
"privacyAndCompliance": {
"allowsPurchases": false,
"usesPersonalInfo": false,
"isChildDirected": false,
"isExportCompliant": true,
"containsAds": false,
"locales": {
"en-US": {
"privacyPolicyUrl": "http://www.myprivacypolicy.sampleskill.com",
"termsOfUseUrl": "http://www.termsofuse.sampleskill.com"
}
}
},
"events": {
"endpoint": {
"uri": "arn:aws:lambda:us-east-1:040623927470:function:sampleSkill"
},
"subscriptions": [
{
"eventName": "SKILL_ENABLED"
},
{
"eventName": "SKILL_DISABLED"
},
{
"eventName": "SKILL_PERMISSION_ACCEPTED"
},
{
"eventName": "SKILL_PERMISSION_CHANGED"
},
{
"eventName": "SKILL_ACCOUNT_LINKED"
}
],
"regions": {
"NA": {
"endpoint": {
"uri": "arn:aws:lambda:us-east-1:040623927470:function:sampleSkill"
}
}
}
}
}
}`)

func TestMinimalSkillDefinition(t *testing.T) {
	res, _ := json.Marshal(minimalSkillDef)
	assert.NotEmpty(t, string(res), "Generated JSON must not be empty")

}

func TestSampleManifest(t *testing.T) {
	var skill Skill
	err := json.Unmarshal(awsDocsExample, &skill)
	assert.Nil(t, err, "Unmarshal returned error: %s", err)
	res, _ := json.Marshal(skill)
	assert.NotEmpty(t, string(res), "Marshal must return JSON")
}
