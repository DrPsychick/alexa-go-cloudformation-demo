package l10n_test

import (
	"bou.ke/monkey"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const Greeting string = "greeting"
const OnlyDe string = "only_de"
const ByeBye string = "byebye"
const WithParam string = "withparam"
const FallbackTest string = "fallback_test"

var registry l10n.LocaleRegistry

var deDE = &l10n.Locale{
	Name: "de-DE",
	TextSnippets: l10n.Snippets{
		Greeting: []string{
			"Hi",
			"Hallo",
			"Howdi",
		},
		OnlyDe: []string{
			"Schleich' di!",
			"Zisch ab!",
			"Ficken Sie sich!",
			"Willst a watschn?",
		},
		WithParam: []string{
			"Hello %s",
		},
	},
	//IntentResponses: l10n.IntentResponses{
	//	Greeting: l10n.IntentResponse{
	//		Title: []string{"title"},
	//		Text:  []string{"text a", "text b"},
	//		SSML:  []string{"<speak>foo bar</speak>"},
	//	},
	//},
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

func TestNewRegistry(t *testing.T) {
	registry = l10n.NewRegistry()
	assert.Empty(t, registry.GetLocales())
	assert.Empty(t, registry.GetDefault())
}

func TestRegisterLocale(t *testing.T) {
	assert.NotNil(t, registry)

	err := registry.Register(deDE)
	assert.NoError(t, err)

	l, err := registry.Resolve(deDE.Name)
	assert.NoError(t, err)

	assert.Equal(t, deDE, l)
	assert.NotEmpty(t, l.GetAny(Greeting))

	err = registry.Register(enUS, l10n.AsDefault())
	assert.NoError(t, err)

	assert.Equal(t, 2, len(registry.GetLocales()))
	assert.Equal(t, enUS, registry.GetDefault())
}

func TestLocaleKeyNotExists(t *testing.T) {
	assert.NotNil(t, registry)

	l, err := registry.Resolve("de-DE")
	assert.NoError(t, err)

	trans := l.GetAny("not exists")
	assert.Empty(t, trans)
}

func TestLocaleGetAny(t *testing.T) {
	assert.NotNil(t, registry)

	patch := monkey.Patch(rand.Intn, func(i int) int {
		return 1
	})
	defer patch.Unpatch()

	l, _ := registry.Resolve("de-DE")
	assert.Equal(t, "Hallo", l.GetAny(Greeting))
}

func TestGetWithParam(t *testing.T) {
	assert.NotNil(t, registry)

	l, err := registry.Resolve("de-DE")
	assert.NoError(t, err)

	assert.Equal(t, "Hello there", l.Get(WithParam, "there"))
	assert.Equal(t, "Hello there", l.GetAny(WithParam, "there"))
	assert.Equal(t, "Hello there", l.GetAll(WithParam, "there")[0])
}

//func TestRegisterFallback(t *testing.T) {
//	err := l10n.Register(enUS, l10n.AsFallbackFor("de-DE"))
//	assert.Nil(t, err, "Register of locale %s failed: %s", enUS.Name, err)
//
//	l, _ := l10n.Resolve("de-DE")
//	assert.NotNil(t, l.GetFallback())                                 // fallback Locale is set
//	assert.NotEmpty(t, l.GetAny(ByeBye))                     // fallback key is used
//	assert.Equal(t, "Fallback text", l.GetAny(FallbackTest)) // fallback content is returned
//}

//func TestFallbackToKey(t *testing.T) {
//	l, _ := l10n.Resolve("en-US")
//	assert.Equal(t, "not_found", l.GetAny("not_found")) // no fallback: return key name
//}
