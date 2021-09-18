package alfalfa

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/drpsychick/go-alexa-lambda/skill"
)

// NewSkill returns a configured SkillBuilder.
func NewSkill() *skill.SkillBuilder {
	return skill.NewSkillBuilder().
		WithLocaleRegistry(loca.Registry).
		WithCategory(skill.CategoryOrganizersAndAssistants).
		WithPrivacyFlag(skill.FlagIsExportCompliant, true)
}

// CreateSkillModels generates and returns a list of Models.
func CreateSkillModels(s *skill.SkillBuilder) (map[string]*skill.Model, error) {
	m := s.Model().
		WithDelegationStrategy(skill.DelegationSkillResponse)

	// we define intents, slots, types in lambda,
	// that's why `newLambda` must be called before this, if not it will panic.

	// this "breaks" `ask dialog --replay` when starting the intent without any slots
	// (as Alexa will try to get the Region slot from the user)
	// Prompts are part of the Alexa dialog, so independent of lambda.
	m.WithElicitationSlotPrompt(loca.AWSStatus, loca.TypeRegionName)
	// add variations (texts) to the prompt
	m.ElicitationPrompt(loca.AWSStatus, loca.TypeRegionName).
		WithVariation("PlainText").
		WithVariation("SSML")

	// do not require elicitation for Region
	m.Intent(loca.AWSStatus).Slot(loca.TypeRegionName).WithElicitation(false)

	// this "breaks" the `ask dialog --replay` testing as alexa asks the user to validate the input
	// m.WithConfirmationSlotPrompt(loca.AWSStatus, loca.TypeAreaName)
	// m.ConfirmationPrompt(loca.AWSStatus, loca.TypeAreaName).
	//	WithVariation("SSML")

	// create a Validation prompt, connected to type-values
	m.WithValidationSlotPrompt(loca.TypeRegionName, skill.ValidationTypeHasMatch)
	m.ValidationPrompt(loca.TypeRegionName, skill.ValidationTypeHasMatch).
		WithVariation("PlainText")

	// ValidationTypeInSet requires values -> we need to pass a key
	m.WithValidationSlotPrompt(loca.TypeRegionName, skill.ValidationTypeInSet, loca.TypeRegionValues)
	m.ValidationPrompt(loca.TypeRegionName, skill.ValidationTypeInSet).
		WithVariation("PlainText")

	// m.Intent(loca.AWSStatus).Slot(loca.TypeRegionName).
	//	WithValidationRule(alexa.ValidationTypeHasMatch)

	return s.BuildModels()
}
