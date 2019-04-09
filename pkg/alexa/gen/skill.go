package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

// SkillDefinition is a logical construct for the skill.
type Skill struct {
	Category            alexa.Category
	Locales             map[alexa.Locale]LocaleDef
	Intents             []Intent
	Models              []Model
	Types               []Type
	Countries           []alexa.Country
	TestingInstructions string
	Privacy             []bool // TODO
}

// NewSkill returns a new basic skill
func NewSkill() *Skill {
	s := &Skill{}
	s.AddIntentString(alexa.HelpIntent)
	s.AddIntentString(alexa.CancelIntent)
	s.AddIntentString(alexa.StopIntent)
	return s
}

func (s *Skill) SetCategory(category alexa.Category) {
	s.Category = category
}

func (s *Skill) SetTestingInstructions(instructions string) {
	s.TestingInstructions = instructions
}

func (s *Skill) AddLocale(l string, trans *l10n.Locale) {
	if nil == s.Locales {
		s.Locales = make(map[alexa.Locale]LocaleDef)
	}
	s.Locales[alexa.Locale(l)] = LocaleDef{Translations: trans}
}

func (s *Skill) AddIntent(intent Intent) {
	s.Intents = append(s.Intents, intent)
}

func (s *Skill) AddIntentString(intent string) {
	s.Intents = append(s.Intents, NewIntent(intent))
}

func (s *Skill) AddModel(model Model) {
	s.Models = append(s.Models, model)
}

func (s *Skill) AddType(t Type) {
	// TODO: validate rules!
	s.Types = append(s.Types, t)
}

func (s *Skill) AddTypeString(t string) {
	s.Types = append(s.Types, Type(t))
}

func (s *Skill) AddCountry(c alexa.Country) {
	s.Countries = append(s.Countries, c)
}

// Build renders the skill.json.
func (s *Skill) Build() alexa.Skill {
	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version: "1.0",
		},
	}
	skill.Manifest.Publishing.Category = s.Category

	for _, l := range s.Locales {
		l.Build(skill)
	}

	// TODO: ensure unique occurance
	skill.Manifest.Publishing.Countries = s.Countries
	skill.Manifest.Publishing.TestingInstructions = s.TestingInstructions

	return *skill
}

func (s *Skill) ValidateTypes() error {
	var tm = make(map[string]bool)
	for _, t := range s.Types {
		tm[string(t)] = true
	}
	for _, i := range s.Intents {
		for _, sl := range i.Slots {
			if len(s.Types) == 0 {
				return fmt.Errorf("No types defined in the skill!")
			}
			if _, ok := tm[string(sl.Type)]; !ok {
				return fmt.Errorf("Type validation error: intent slot %s uses type %s which is not defined in the skill", sl.Name, string(sl.Type))
			}
		}
	}
	return nil
}

type LocaleDef struct {
	Translations *l10n.Locale
}

func (l *LocaleDef) Build(s *alexa.Skill) {
	if s.Manifest.Publishing.Locales == nil {
		s.Manifest.Publishing.Locales = make(map[alexa.Locale]alexa.LocaleDef)
	}
	s.Manifest.Publishing.Locales[alexa.Locale(l.Translations.Name)] = alexa.LocaleDef{
		Name:         l.Translations.GetSnippet(l10n.KeySkillName),
		Description:  l.Translations.GetSnippet(l10n.KeySkillDescription),
		Summary:      l.Translations.GetSnippet(l10n.KeySkillSummary),
		Examples:     l.Translations.GetAllSnippets(l10n.KeySkillExamplePhrases),
		Keywords:     l.Translations.GetAllSnippets(l10n.KeySkillKeywords),
		SmallIconURI: l.Translations.GetSnippet(l10n.KeySkillSmallIconURI),
		LargeIconURI: l.Translations.GetSnippet(l10n.KeySkillLargeIconURI),
	}
}

type Intent struct {
	Name  string
	Slots []Slot
}

func NewIntent(name string) Intent {
	return Intent{Name: name}
}

func (i *Intent) AddSlot(slot Slot) {
	i.Slots = append(i.Slots, slot)
}

type Slot struct {
	Name string // specific to intent
	Type Type   // global for skill
}

func NewSlot(name string, t Type) Slot {
	return Slot{Name: name, Type: t}
}

type Type string

func NewType(t string) Type {
	return Type(t)
}

type Model struct {
}
