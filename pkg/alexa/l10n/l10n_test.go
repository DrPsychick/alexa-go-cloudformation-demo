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

func TestDefaultRegistry(t *testing.T) {
	// make sure we have an empty default
	l10n.DefaultRegistry = l10n.NewRegistry()

	err := l10n.Register(deDE)
	assert.NoError(t, err)
	deDE2, err := l10n.Resolve("de-DE")
	assert.NoError(t, err)
	assert.Equal(t, deDE, deDE2)
}

func TestNewRegistry(t *testing.T) {
	registry = l10n.NewRegistry()
	assert.Empty(t, registry.GetLocales())
	assert.Nil(t, registry.GetDefault())
}

func TestRegisterLocale(t *testing.T) {
	assert.NotNil(t, registry)
	err := registry.SetDefault("de-DE")
	assert.Error(t, err)

	_, err = registry.Resolve("en-US")
	assert.Error(t, err)

	err = registry.Register(deDE)
	assert.NoError(t, err)
	err = registry.Register(deDE)
	assert.Error(t, err)

	l, err := registry.Resolve(deDE.Name)
	assert.NoError(t, err)

	assert.Equal(t, deDE, l)
	assert.NotEmpty(t, l.GetAny(Greeting))

	err = registry.Register(enUS, l10n.AsDefault())
	assert.NoError(t, err)

	assert.Equal(t, 2, len(registry.GetLocales()))
	assert.Equal(t, enUS, registry.GetDefault())

	err = registry.SetDefault("de-DE")
	assert.NoError(t, err)
	assert.Equal(t, deDE, registry.GetDefault())

}

func TestLocale(t *testing.T) {
	l := l10n.NewLocale("fo-BA")
	assert.IsType(t, &l10n.Locale{}, l)
	assert.Equal(t, "fo-BA", l.Name)

	vals := []string{"bar1", "bar2"}
	l.Set("foo", vals)
	v2 := l.GetAll("foo")
	assert.Equal(t, vals, v2)
}

func TestLocaleKeyNotExists(t *testing.T) {
	assert.NotNil(t, registry)

	l, err := registry.Resolve("de-DE")
	assert.NoError(t, err)

	tx := l.Get("not exists")
	assert.Empty(t, tx)
	tx = l.GetAny("not exists")
	assert.Empty(t, tx)
	txs := l.GetAll("not exists")
	assert.Empty(t, txs)

	_, err = deDE.TextSnippets.GetFirst("not exists")
	assert.Error(t, err)
	_, err = deDE.TextSnippets.GetAny("not exists")
	assert.Error(t, err)
	_, err = deDE.TextSnippets.GetAll("not exists")
	assert.Error(t, err)
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
