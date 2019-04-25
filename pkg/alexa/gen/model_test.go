package gen_test

import (
	"encoding/json"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Case 1: input multiple languages directly
func TestModelBuilder_BuildLocale(t *testing.T) {
	loc, err := registry.Resolve("en-US")
	loc.Set("MyIntent_SlotName_Samples", []string{"of {Slot}"})
	assert.NoError(t, err)

	mb := gen.NewModelBuilder().
		WithDelegationStrategy(alexa.DelegationSkillResponse).
		AddLocale("en-US", "my skill").
		AddLocale("de-DE", "mein skill")

	mb.AddType("TypeSlotOne").
		WithLocaleValues("en-US", []string{"One"}).
		WithLocaleValues("de-DE", []string{"Eins"})

	mb.AddIntent("MyIntent").
		WithLocaleSamples(loc.GetName(), loc.GetAll("MyIntent_Samples")).
		WithLocaleSamples("de-DE", []string{"sample eins", "sample zwei"}).
		AddSlot("SlotName", "TypeSlotOne").
		WithLocaleSamples(loc.GetName(), loc.GetAll("MyIntent_SlotName_Samples")).
		WithLocaleSamples("de-DE", []string{"von {Slot}"})

	mb.AddElicitationSlotPrompt("MyIntent", "SlotName").
		AddVariation("PlainText").
		WithLocaleValue("de-DE", "PlainText", []string{"Was?", "Wie bitte?"}).
		WithLocaleValue(loc.GetName(), "PlainText", []string{"What?"})

	mb.AddConfirmationSlotPrompt("MyIntent", "SlotName").
		AddVariation("PlainText").
		WithLocaleValue(loc.GetName(), "PlainText", []string{"Sure?"}).
		WithLocaleValue("de-DE", "PlainText", []string{"Sicher?"})

	m, err := mb.BuildLocale(loc.GetName())
	assert.NoError(t, err)
	res, err := json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, "TypeSlotOne", m.Model.Language.Types[0].Name)
	assert.Equal(t, "MyIntent", m.Model.Language.Intents[0].Name)
	assert.Equal(t, "Elicit.Intent-MyIntent.IntentSlot-SlotName", m.Model.Prompts[0].Id)
	assert.Equal(t, "Sure?", m.Model.Prompts[1].Variations[0].Value)
	assert.NotContains(t, "null", string(res))
	fmt.Printf("%s = %s\n", loc.GetName(), string(res))

	m, err = mb.BuildLocale("de-DE")
	assert.NoError(t, err)
	res, err = json.MarshalIndent(m, "", "  ")
	assert.NoError(t, err)
	assert.Equal(t, "sample eins", m.Model.Language.Intents[0].Samples[0])
	assert.Equal(t, "von {Slot}", m.Model.Language.Intents[0].Slots[0].Samples[0])
	assert.Equal(t, "Wie bitte?", m.Model.Prompts[0].Variations[1].Value)
	assert.Equal(t, "Sicher?", m.Model.Prompts[1].Variations[0].Value)
	fmt.Printf("%s = %s\n", "de-DE", string(res))
}

// Case 2: input LocaleRegistry
func TestModelBuilder_Build(t *testing.T) {
	mb := gen.NewModelBuilder().
		WithLocaleRegistry(registry).
		WithDelegationStrategy(alexa.DelegationSkillResponse)
	mb.AddType("MyType")
	mb.AddIntent("MyIntent")
	mb.AddIntent("SlotIntent").
		AddSlot("SlotName", "MyType")

	mb.AddElicitationSlotPrompt("SlotIntent", "SlotName").
		AddVariation("PlainText").
		AddVariation("SSML")

	ms, err := mb.Build()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(ms))

	res, err := json.MarshalIndent(ms["en-US"], "", "  ")
	assert.NoError(t, err)
	// must contain translations from enUS
	assert.Contains(t, string(res), "what about slot {SlotName}")
	assert.Contains(t, string(res), "I'm sorry")
	assert.Contains(t, string(res), "Which slot")

	fmt.Printf("en-US: %s\n", string(res))

	mb.AddElicitationSlotPrompt("SlotIntent", "SlotName")
	_, err = mb.Build()
	assert.Error(t, err)
}

// individual functions
func TestModelWith(t *testing.T) {
	mb := gen.NewModelBuilder().
		WithLocaleRegistry(registry).
		WithInvocation("MyInvocationKey").
		AddLocale("en-US", "invoke me")
	// fails because "en-US" is already in registry!
	assert.Nil(t, mb)

	mb = gen.NewModelBuilder().
		WithInvocation("MyInvocationKey").
		AddLocale("fo-BA", "invoke me")
	assert.IsType(t, &gen.ModelBuilder{}, mb)

	ms, err := mb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "invoke me", ms["fo-BA"].Model.Language.Invocation)
}

func TestIntentBuilder(t *testing.T) {
	loc := registry.GetDefault()

	// define the intent
	ib := gen.NewModelIntentBuilder("MyIntent").
		WithLocaleRegistry(registry).
		WithSamples("MyIntent_Samples")
	ib.AddSlot("SlotName", "TypeSlotOne").
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

	// validate alexa.DialogIntent
	di, err := ib.BuildDialogIntent(registry.GetDefault().GetName())
	assert.NoError(t, err)
	assert.IsType(t, alexa.DialogIntent{}, di)
	assert.Equal(t, len(li.Slots), len(di.Slots))

	res, err = json.MarshalIndent(di, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	// rendering JSON should never contain "null"
	assert.NotContains(t, string(res), "null")
}

func TestTypeBuilder(t *testing.T) {
	tb := gen.NewModelTypeBuilder("MY_type").
		WithLocaleRegistry(registry).
		WithValues("MY_type_values")
	mt, err := tb.Build(registry.GetDefault().GetName())
	assert.NoError(t, err)
	assert.IsType(t, alexa.ModelType{}, mt)
	assert.Equal(t, "MY_type", mt.Name)

	res, err := json.MarshalIndent(mt, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
}

func TestNewElicitationModelPromptBuilder(t *testing.T) {
	r := l10n.NewRegistry()
	r.Register(l10n.NewLocale("en-US"))

	pb := gen.NewElicitationPromptBuilder("MyIntent", "MySlot").
		WithLocaleRegistry(r)
	pb.AddVariation("PlainText").
		WithLocaleValue("en-US", "PlainText", []string{"What?"})

	mp, err := pb.BuildLocale("en-US")
	assert.NoError(t, err)
	assert.IsType(t, alexa.ModelPrompt{}, mp)
	assert.Contains(t, mp.Id, "Elicit")
	assert.Contains(t, mp.Id, "MyIntent")
	assert.Contains(t, mp.Id, "MySlot")
	assert.Equal(t, "What?", mp.Variations[0].Value)

	pb = gen.NewConfirmationPromptBuilder("MyIntent", "MySlot").
		WithLocaleRegistry(r)
	pb.AddVariation("SSML").
		WithLocaleValue("en-US", "SSML", []string{"Confirm!"})
	mp, err = pb.BuildLocale("en-US")
	assert.NoError(t, err)
	assert.Contains(t, mp.Id, "Confirm")
	assert.Contains(t, mp.Id, "MyIntent")
	assert.Contains(t, mp.Id, "MySlot")
	assert.Equal(t, "Confirm!", mp.Variations[0].Value)

	// fails without locale
	pb = gen.NewElicitationPromptBuilder("MyIntent", "MySlot")
	_, err = pb.BuildLocale("en-US")
	assert.Error(t, err)

	// fails without variations value
	pb.WithLocaleRegistry(r).
		AddVariation("NotExists")
	_, err = pb.BuildLocale("en-US")
	assert.Error(t, err)
}

func TestModelBuilder_AddIntent(t *testing.T) {
	mb := gen.NewModelBuilder().
		WithLocaleRegistry(registry)

	mb.AddIntent("MyIntent").
		WithSamples("MyIntent_Samples")

	m, err := mb.Build()
	assert.NoError(t, err)
	assert.NotEmpty(t, m["en-US"].Model.Language.Intents)
	i := m["en-US"].Model.Language.Intents
	assert.Equal(t, "MyIntent", i[0].Name)
	assert.Equal(t, []string{"say one", "say two"}, i[0].Samples)
}
