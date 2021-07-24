package main

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/hamba/cmd"
)

// Application =============================

func newApplication(c *cmd.Context) (*alfalfa.Application, error) {
	app := alfalfa.NewApplication(
		c.Logger(),
		c.Statter(),
	)

	return app, nil
}

func newSkill() *gen.SkillBuilder {
	return alfalfa.NewSkill()
}

func createSkillModels(s *gen.SkillBuilder) (map[string]*alexa.Model, error) {
	return alfalfa.CreateSkillModels(s)
}
