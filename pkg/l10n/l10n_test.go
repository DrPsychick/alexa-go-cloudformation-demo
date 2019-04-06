package l10n_test

import (
	"fmt"
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

const Greeting l10n.Key = "greeting"
const FuckYou l10n.Key = "fuckyou"
const ByeBye l10n.Key = "byebye"

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: l10n.Snippets{
		Greeting: []string{
			"Hi",
			"Hallo",
			"Howdi",
		},
		FuckYou: []string{
			"Schleich' di!",
			"Zisch ab!",
			"Ficken Sie sich!",
			"Willst a watschn?",
		},
	},
}

var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: l10n.Snippets{
		Greeting: []string{
			"Hi",
			"Hello",
		},
		ByeBye: []string{
			"Have a nice day!",
			"Bugger off...",
			"Hasta la vista, baby!",
		},
	},
}

//func TestLocale_GetSnippet(t *testing.T) {
//	rnd := 1
//	patch := monkey.Patch(rand.Intn, func() int {
//		return rnd
//	})
//	defer patch.Unpatch()
//
//	text := deDE.GetSnippet(Greeting)
//	assert.Equal(t, "Hallo", text)
//}

func TestRegistry(t *testing.T) {
	l10n.Register(enUS, l10n.AsDefault())
	l, err := l10n.Resolve(enUS.Name)
	assert.Nil(t, err, "register locale %s must resolve", enUS.Name)
	assert.Equal(t, enUS, l)
	assert.Equal(t, "Hello", l.GetSnippet(Greeting))
}

func TestFallback(t *testing.T) {
	l10n.Register(deDE, l10n.AsDefault())
	err := l10n.Register(enUS, l10n.AsDefault(), l10n.AsFallbackFor("de-DE"))
	assert.Nil(t, err, "Register of locale %s failed: %s", enUS.Name, err)
	fmt.Printf("Default %s\n", l10n.DefaultRegistry.DefaultLocale)

	l, _ := l10n.Resolve("de-DE")
	assert.NotNil(t, l.Fallback)

	assert.Equal(t, "Bugger off...", l.GetSnippet(ByeBye))
}
