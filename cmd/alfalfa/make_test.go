package main

import (
	"encoding/json"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/loca"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeSkill(t *testing.T) {
	sb, err := createSkill(loca.Registry)
	assert.NoError(t, err)

	s, err := sb.Build()
	assert.NoError(t, err)

	res, err := json.MarshalIndent(s, "", "  ")
	assert.NotEmpty(t, string(res))

	fmt.Printf("%s\n", string(res))
}

func TestMakeModels(t *testing.T) {
	sb, err := createSkill(loca.Registry)
	assert.NoError(t, err)

	ms, err := createModels(sb)
	assert.NoError(t, err)

	for l, m := range ms {
		res, err := json.MarshalIndent(m, "", "  ")
		assert.NoError(t, err)
		assert.NotEmpty(t, string(res))

		fmt.Printf("%s: %s\n", l, string(res))
	}
}
