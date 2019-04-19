package gen_test

import (
	"encoding/json"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntentBuilder(t *testing.T) {
	loc := registry.GetDefault()

	// define the intent
	ib := gen.NewModelIntentBuilder("MyIntent").
		WithLocaleRegistry(registry).
		WithSamples("MyIntent_Samples")
	ib.AddSlot("SlotName").
		WithType("TypeSlotOne").
		WithSamples("SlotOneSamples")

	// validate alexa.ModelIntent
	li, err := ib.BuildLanguageIntent(loc.GetName())
	assert.NoError(t, err)
	assert.IsType(t, alexa.ModelIntent{}, li)
	assert.Equal(t, "MyIntent", li.Name)
	assert.Equal(t, loc.GetAll("MyIntent_Samples"), li.Samples)
	assert.NotEmpty(t, li.Slots)

	res, err := json.MarshalIndent(li, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	assert.NotContains(t, string(res), "null")
	fmt.Printf("MyIntent LanguageModel = %s\n", string(res))

	// validate alexa.DialogIntent
	di := ib.BuildDialogIntent(registry.GetDefault().GetName())
	assert.IsType(t, alexa.DialogIntent{}, di)
	assert.Equal(t, len(li.Slots), len(di.Slots))

	res, err = json.MarshalIndent(di, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	// rendering JSON should never contain "null"
	assert.NotContains(t, string(res), "null")
	fmt.Printf("MyIntent Dialog = %s\n", string(res))
}

func TestTypeBuilder(t *testing.T) {
	tb := gen.NewModelTypeBuilder("MY_type").
		WithLocaleRegistry(registry).
		WithValuesName("MY_type_values")
	mt, err := tb.Build(registry.GetDefault().GetName())
	assert.NoError(t, err)
	assert.IsType(t, alexa.ModelType{}, mt)
	assert.Equal(t, "MY_type", mt.Name)

	res, err := json.MarshalIndent(mt, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	fmt.Printf("MY_type = %s\n", string(res))
}

func TestModelBuilder_AddIntent(t *testing.T) {
	mb := gen.NewModelBuilder().
		WithLocaleRegistry(registry)

	mb.AddIntent("MyIntent").
		WithSamples("MyIntent_Samples")

}

// Use ModelBuilder by manually passing locale strings
func TestLocaleModelBuilder(t *testing.T) {
	loc, err := registry.Resolve("en-US")
	assert.NoError(t, err)

	mb := gen.NewModelBuilder().
		AddLocale("en-US", "my skill").
		AddLocale("de-DE", "mein skill")

	mb.AddType("TypeSlotOne").
		WithLocaleValues("en-US", []string{"One"}).
		WithLocaleValues("de-DE", []string{"Eins"})

	mb.AddIntent("MyIntent").
		WithLocaleSamples(loc.GetName(), loc.GetAll("MyIntent_Samples")).
		WithLocaleSamples("de-DE", []string{"sample eins", "sample zwei"}).
		AddSlot("SlotName").
		WithType("TypeSlotOne")

	m, err := mb.BuildLocale(loc.GetName())
	assert.NoError(t, err)
	res, err := json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, "TypeSlotOne", m.Model.Language.Types[0].Name)
	assert.Equal(t, "MyIntent", m.Model.Language.Intents[0].Name)
	assert.NotContains(t, "null", string(res))
	fmt.Printf("%s = %s\n", loc.GetName(), string(res))

	m, err = mb.BuildLocale("de-DE")
	assert.NoError(t, err)
	res, err = json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, "sample eins", m.Model.Language.Intents[0].Samples[0])
	fmt.Printf("%s = %s\n", "de-DE", string(res))
}

func TestModelBuilder_Build(t *testing.T) {
	mb := gen.NewModelBuilder().
		WithLocaleRegistry(registry)
	mb.AddIntent("MyIntent")

	ms, err := mb.Build()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ms))

	res, err := json.MarshalIndent(ms["en-US"], "", "  ")
	assert.NoError(t, err)

	fmt.Printf("en-US: %s\n", string(res))
}
