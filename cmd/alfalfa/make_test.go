package main

import (
	"encoding/json"
	"github.com/drpsychick/alexa-go-cloudformation-demo"
	"github.com/hamba/logger"
	"github.com/hamba/statter/l2met"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMakeSkill(t *testing.T) {
	sb := newSkill()

	s, err := sb.Build()
	assert.NoError(t, err)

	res, err := json.MarshalIndent(s, "", "  ")
	assert.NotEmpty(t, string(res))
}

func TestMakeModels(t *testing.T) {
	l := logger.New(logger.StreamHandler(os.Stdout, logger.LogfmtFormat()))
	app := alfalfa.NewApplication(
		l,
		l2met.New(l, ""),
	)
	sb := newSkill()
	newLambda(app, sb)

	ms, err := createSkillModels(sb)
	assert.NoError(t, err)

	for _, m := range ms {
		res, err := json.MarshalIndent(m, "", "  ")
		assert.NoError(t, err)
		assert.NotEmpty(t, string(res))
	}
}
