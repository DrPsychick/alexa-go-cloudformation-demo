package gen_test

import (
	"encoding/json"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

var registry = l10n.NewRegistry()
var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[l10n.Key][]string{
		"MyIntent_Samples": []string{"say one", "say two"},
		"MY_type_values":   []string{"Value 1", "Value 2"},
	},
}

func init() {
	registry.Register(enUS, l10n.AsDefault())
}

func TestSetup(t *testing.T) {
	assert.NotEmpty(t, registry.GetLocales())
	assert.NotEmpty(t, registry.GetDefaultLocale())
}

func TestIntentBuilder(t *testing.T) {
	def := registry.GetDefaultLocale()
	loc, err := registry.Resolve(def)
	assert.NoError(t, err)

	// define the intent
	ib := gen.NewModelIntentBuilder("MyIntent").
		WithRegistry(registry).
		WithSamples("MyIntent_Samples")
	ib.WithSlot("SlotName").
		WithType("TypeSlotOne").
		WithSamples("SlotOneSamples")

	// validate alexa.ModelIntent
	li := ib.BuildLanguageIntent(def)
	assert.IsType(t, alexa.ModelIntent{}, li)
	assert.Equal(t, "MyIntent", li.Name)
	assert.Equal(t, loc.GetAllSnippets(l10n.Key("MyIntent_Samples")), li.Samples)
	assert.Nil(t, li.Slots)

	res, err := json.MarshalIndent(li, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	assert.NotContains(t, string(res), "null")

	// validate alexa.DialogIntent
	di := ib.BuildDialogIntent(registry.GetDefaultLocale())
	assert.IsType(t, alexa.DialogIntent{}, di)
	assert.Nil(t, di.Slots)

	res, err = json.MarshalIndent(di, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	// rendering JSON should never contain "null"
	assert.NotContains(t, string(res), "null")
}

func TestTypeBuilder(t *testing.T) {
	tb := gen.NewModelTypeBuilder("MY_type").
		WithRegistry(registry).
		WithValuesName("MY_type_values")
	mt := tb.Build(registry.GetDefaultLocale())
	assert.IsType(t, alexa.ModelType{}, mt)
	assert.Equal(t, "MY_type", mt.Name)

	res, _ := json.MarshalIndent(mt, "", "  ")
	assert.NotEmpty(t, string(res))

}
