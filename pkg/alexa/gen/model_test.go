package gen_test

import (
	"encoding/json"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

var registry = l10n.NewRegistry()
var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[string][]string{
		"MyIntent_Samples": []string{"say one", "say two"},
		"MY_type_values":   []string{"Value 1", "Value 2"},
	},
}

func init() {
	registry.Register(enUS, l10n.AsDefault())
}

func TestSetup(t *testing.T) {
	assert.NotEmpty(t, registry.GetLocales())
	assert.NotEmpty(t, registry.GetDefault())
}

//func TestIntentBuilder(t *testing.T) {
//	loc := registry.GetDefault()
//
//	// define the intent
//	ib := gen.NewModelIntentBuilder("MyIntent").
//		WithLocaleRegistry(registry).
//		WithSamples("MyIntent_Samples")
//	ib.AddSlot("SlotName").
//		WithType("TypeSlotOne").
//		WithSamples("SlotOneSamples")
//
//	// validate alexa.ModelIntent
//	li := ib.BuildLanguageIntent(loc.Name)
//	assert.IsType(t, alexa.ModelIntent{}, li)
//	assert.Equal(t, "MyIntent", li.Name)
//	assert.Equal(t, loc.GetAllSnippets("MyIntent_Samples"), li.Samples)
//	assert.NotEmpty(t, li.Slots)
//
//	res, err := json.MarshalIndent(li, "", "  ")
//	assert.NoError(t, err)
//	assert.NotEmpty(t, string(res))
//	assert.NotContains(t, string(res), "null")
//
//	// validate alexa.DialogIntent
//	di := ib.BuildDialogIntent(registry.GetDefault().Name)
//	assert.IsType(t, alexa.DialogIntent{}, di)
//	assert.Nil(t, di.Slots)
//
//	res, err = json.MarshalIndent(di, "", "  ")
//	assert.NoError(t, err)
//	assert.NotEmpty(t, string(res))
//	// rendering JSON should never contain "null"
//	assert.NotContains(t, string(res), "null")
//}

//func TestTypeBuilder(t *testing.T) {
//	tb := gen.NewModelTypeBuilder("MY_type").
//		WithLocaleRegistry(registry).
//		WithValuesName("MY_type_values")
//	mt := tb.Build(registry.GetDefault())
//	assert.IsType(t, alexa.ModelType{}, mt)
//	assert.Equal(t, "MY_type", mt.Name)
//
//	res, _ := json.MarshalIndent(mt, "", "  ")
//	assert.NotEmpty(t, string(res))
//
//}

func TestModelBuilder(t *testing.T) {
	loc, err := registry.Resolve("en-US")
	assert.NoError(t, err)

	mb := gen.NewModelBuilder().
		AddLocale("en-US", "my skill").
		AddLocale("de-DE", "mein skill")

	mb.AddType("TypeSlotOne").
		WithLocaleValues("en-US", []string{"One"}).
		WithLocaleValues("de-DE", []string{"Eins"})

	mb.AddIntent("MyIntent").
		WithLocaleSamples(loc.Name, loc.GetAll("MyIntent_Samples")).
		WithLocaleSamples("de-DE", []string{"sample eins", "sample zwei"}).
		AddSlot("SlotName").
		WithType("TypeSlotOne")

	m := mb.BuildLocale(loc.Name)
	res, err := json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, "TypeSlotOne", m.Model.Language.Types[0].Name)
	assert.Equal(t, "MyIntent", m.Model.Language.Intents[0].Name)
	fmt.Printf("%s = %s\n", loc.Name, string(res))

	//assert.Empty(t, string(res), "locale: %s", loc.Name)

	m = mb.BuildLocale("de-DE")
	res, err = json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	fmt.Printf("%s = %s\n", "de-DE", string(res))

	fmt.Printf("locales = %s\n", mb.GetLocales())
}

//func TestModelBuilderLocale(t *testing.T) {
//	mb := gen.NewModelBuilder()
//	mb.AddLocaleIntent("en-US", "MyIntent").
//		WithLocaleSamples([]string{"serve my intent", "do what i mean"})
//
//	for l, m := range mb.Build() {
//		res, err := json.MarshalIndent(m, "", "  ")
//		assert.NoError(t, err)
//		assert.Empty(t, string(res), "locale %s", l)
//	}
//}
