package gen_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = gen.NewSkillBuilder()

func TestSkill(t *testing.T) {
	s.SetCategory(alexa.CategoryCommunication)
	s.SetTestingInstructions("My instructions")
	skill, _ := s.Build()

	assert.Equal(t, alexa.CategoryCommunication, skill.Manifest.Publishing.Category)
	assert.Equal(t, "My instructions", skill.Manifest.Publishing.TestingInstructions)
}

func TestLocales(t *testing.T) {
	var deDE = &l10n.Locale{
		Name:      "de-DE",
		Countries: []alexa.Country{"US", "CA", "DE"},
		TextSnippets: map[l10n.Key][]string{
			l10n.KeySkillName: []string{"My skill name"},
		},
		IntentResponses: map[l10n.Key]l10n.IntentResponse{
			"WithSlots": l10n.IntentResponse{},
		},
	}
	err := l10n.Register(deDE)
	assert.NoError(t, err)

	s.AddLocale("de-DE", deDE)
	skill, _ := s.Build()

	for l, _ := range skill.Manifest.Publishing.Locales {
		loc, err := l10n.Resolve(l)
		assert.NoError(t, err)
		assert.Equal(t, "My skill name", loc.GetSnippet(l10n.KeySkillName))
		assert.NotEmpty(t, loc.GetSnippet("Test"))
		_, err = loc.GetIntent("DoesNotExist")
		assert.Error(t, err)
	}

	s.AddCountry(alexa.CountryCanada)
	s.AddCountry(alexa.CountryGermany)

	skill, _ = s.Build()
	assert.Contains(t, skill.Manifest.Publishing.Countries, alexa.CountryCanada)
	assert.Contains(t, skill.Manifest.Publishing.Countries, alexa.CountryGermany)
}

func TestIntentWithSlots(t *testing.T) {
	ty := gen.NewType("MY_Type")
	i := gen.NewIntent("WithSlots")
	i.AddSlot(gen.NewSlot("MySlot-1", ty))
	s.AddIntent(i)

	_, err := s.Build()
	assert.NoError(t, err)
	_, err = s.BuildModels()
	assert.NoError(t, err)

	// 3 basic + 2 added in this test case:
	// TODO: fix this if _, ok := models["de-DE"]; ok {
	//	assert.Equal(t, 5, len(models["de-DE"].Model.Language.Intents))
	//}
}

func TestValidateTypes(t *testing.T) {
	assert.Error(t, s.ValidateTypes())
	s.AddTypeString("MY_Type")
	assert.NoError(t, s.ValidateTypes())
	s.AddType(gen.NewType("MY_Type2"))
	assert.NoError(t, s.ValidateTypes())
}

func TestBuildImmutability(t *testing.T) {
	s1, _ := s.Build()
	s2, _ := s.Build()

	assert.Equal(t, s1, s2, "Building skill is not immutable!")
}
