package main

import (
	alfalfa "github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/skill"
	"github.com/hamba/cmd"
)

func newApplication(c *cmd.Context) (*alfalfa.Application, error) { //nolint:unparam
	app := alfalfa.NewApplication(
		c.Logger(),
		c.Statter(),
	)

	return app, nil
}

func newSkill() *skill.SkillBuilder {
	return alfalfa.NewSkill()
}

func createSkillModels(s *skill.SkillBuilder) (map[string]*skill.Model, error) {
	return alfalfa.CreateSkillModels(s)
}
