package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
)

// setup the skill
//func setupSkill() {
//	var skill alexa.Skill
//
//	skill.Category = alexa.CategoryOrganizersAndAssistants
//	skill.Locales = []alexa.Locale{"de-DE", "en-US"}
//
//}

// NewSkill returns a configured SkillBuilder
func NewSkill() *gen.SkillBuilder {
	return gen.NewSkillBuilder().
		WithLocaleRegistry(loca.Registry).
		WithCategory(alexa.CategoryOrganizersAndAssistants).
		WithPrivacyFlag(gen.FlagIsExportCompliant, true)

}

// CreateSkillModels generates and returns a list of Models.
func CreateSkillModels(s *gen.SkillBuilder) (map[string]*alexa.Model, error) {
	m := s.Model().
		WithDelegationStrategy(alexa.DelegationSkillResponse)

	// we define intents, slots, types in lambda,
	// that's why `newLambda` must be called before this, if not it will panic.

	// this "breaks" `ask dialog --replay` when starting the intent without any slots (as Alexa will try to get the Region slot from the user)
	// Prompts are part of the Alexa dialog, so independent of lambda.
	m.WithElicitationSlotPrompt(loca.AWSStatus, loca.TypeRegionName)
	// add variations (texts) to the prompt
	m.ElicitationPrompt(loca.AWSStatus, loca.TypeRegionName).
		WithVariation("PlainText").
		WithVariation("SSML")

	// do not require elicitation for Region
	m.Intent(loca.AWSStatus).Slot(loca.TypeRegionName).WithElicitation(false)

	// this "breaks" the `ask dialog --replay` testing as alexa asks the user to validate the input
	//m.WithConfirmationSlotPrompt(loca.AWSStatus, loca.TypeAreaName)
	//m.ConfirmationPrompt(loca.AWSStatus, loca.TypeAreaName).
	//	WithVariation("SSML")

	// create a Validation prompt, connected to type-values
	m.WithValidationSlotPrompt(loca.TypeRegionName, alexa.ValidationTypeHasMatch)
	m.ValidationPrompt(loca.TypeRegionName, alexa.ValidationTypeHasMatch).
		WithVariation("PlainText")

	// ValidationTypeInSet requires values -> we need to pass a key
	m.WithValidationSlotPrompt(loca.TypeRegionName, alexa.ValidationTypeInSet, loca.TypeRegionValues)
	m.ValidationPrompt(loca.TypeRegionName, alexa.ValidationTypeInSet).
		WithVariation("PlainText")

	//m.Intent(loca.AWSStatus).Slot(loca.TypeRegionName).
	//	WithValidationRule(alexa.ValidationTypeHasMatch)

	return s.BuildModels()
}
