package l10n_test

import (
	"bou.ke/monkey"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const Greeting l10n.Key = "greeting"
const FuckYou l10n.Key = "fuckyou"
const ByeBye l10n.Key = "byebye"
const FallbackTest l10n.Key = "fallback_test"

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
		FallbackTest: []string{
			"Fallback text",
		},
	},
}

func init() {
}

func TestRegisterLocale(t *testing.T) {
	l10n.Register(deDE)
	l, err := l10n.Resolve(deDE.Name)
	assert.Nil(t, err, "register locale %s must resolve", deDE.Name)
	assert.Equal(t, deDE, l)
	assert.NotEmpty(t, l.GetSnippet(Greeting))
}

func TestLocaleGetSnippet(t *testing.T) {
	patch := monkey.Patch(rand.Intn, func(i int) int {
		return 1
	})
	defer patch.Unpatch()

	l, _ := l10n.Resolve("de-DE")
	assert.Equal(t, "Hallo", l.GetSnippet(Greeting))
}

func TestRegisterFallback(t *testing.T) {
	err := l10n.Register(enUS, l10n.AsFallbackFor("de-DE"))
	assert.Nil(t, err, "Register of locale %s failed: %s", enUS.Name, err)

	l, _ := l10n.Resolve("de-DE")
	assert.NotNil(t, l.Fallback)                                 // fallback Locale is set
	assert.NotEmpty(t, l.GetSnippet(ByeBye))                     // fallback Key is used
	assert.Equal(t, "Fallback text", l.GetSnippet(FallbackTest)) // fallback content is returned
}

func TestFallbackToKey(t *testing.T) {
	l, _ := l10n.Resolve("en-US")
	assert.Equal(t, "not_found", l.GetSnippet(l10n.Key("not_found"))) // no fallback: return key name
}
