package gen_test

import (
	"encoding/json"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = gen.NewSkill()

func TestSkill(t *testing.T) {
	s.SetCategory(alexa.CategoryCommunication)
	s.SetTestingInstructions("My instructions")
	assert.NotEmpty(t, s.Category)
	assert.NotEmpty(t, s.TestingInstructions)
}

func TestIntent(t *testing.T) {
	s.AddIntent(gen.NewIntent("Foo"))
	assert.NotEmpty(t, s.Intents)
}

func TestIntentWithSlots(t *testing.T) {
	ty := gen.NewType("MY_Type")
	i := gen.NewIntent("WithSlots")
	i.AddSlot(gen.NewSlot("MySlot-1", ty))
	s.AddIntent(i)
	// 3 basic + 2 added in this test case:
	assert.Equal(t, 5, len(s.Intents))
}

func TestLocales(t *testing.T) {
	var deDE = &l10n.Locale{
		Name:      "de-DE",
		Countries: []alexa.Country{"US", "CA", "DE"},
		TextSnippets: map[l10n.Key][]string{
			l10n.KeySkillName: []string{"My skill name"},
		},
	}
	err := l10n.Register(deDE)
	assert.NoError(t, err)

	s.AddLocale("de-DE", deDE)
	s.AddCountry("FR")

	for _, l := range s.Locales {
		loc, err := l10n.Resolve(l.Translations.Name)
		assert.NoError(t, err)
		assert.NotEmpty(t, loc.GetSnippet("Test"))
	}
}

func TestValidateTypes(t *testing.T) {
	assert.Error(t, s.ValidateTypes())
	s.AddTypeString("MY_Type")
	assert.NoError(t, s.ValidateTypes())
	s.AddType(gen.NewType("MY_Type2"))
	assert.NoError(t, s.ValidateTypes())
}

func TestBuild(t *testing.T) {
	skill := s.Build()
	res, _ := json.Marshal(skill)
	assert.Contains(t, string(res), skill.Manifest.Publishing.TestingInstructions)
}
