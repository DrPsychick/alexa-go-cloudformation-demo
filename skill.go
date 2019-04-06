package alfalfa

import "github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/alexa"

// setup the skill
func setupSkill() {
	var skill alexa.Skill

	skill.Category = alexa.CategoryOrganizersAndAssistants
	skill.Locales = []alexa.Locale{"de-DE", "en-US"}

}
