package l10n_test

import (
	"math/rand"
	"testing"

	"bou.ke/monkey"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
	"github.com/stretchr/testify/assert"
)

const Greeting l10n.Key = "greeting"
const FuckYou l10n.Key = "fuckyou"

var deDE = &l10n.Locale{
	Name: "deDE",
	TextSnippets: l10n.Snippets{
		Greeting: []string{
			"Hi",
			"Hallo",
		},
		FuckYou: []string{
			"Schleich' di!",
			"Zisch ab!",
			"Ficken Sie sich!",
			"Willst a watschn?",
		},
	},
}

func TestLocale_GetSnippet(t *testing.T) {
	rnd := 1
	patch := monkey.Patch(rand.Intn, func() int {
		return rnd
	})
	defer patch.Unpatch()

	text := deDE.GetSnippet(Greeting)
	assert.Equal(t, "Hallo", text)
}
