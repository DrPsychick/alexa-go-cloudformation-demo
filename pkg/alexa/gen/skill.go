package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

// SkillDefinition is a logical construct for the skill.
type Skill struct {
	Category alexa.Category
	Intents  []Intent
	Models   []Model
	Types    []Type
	Locales  []alexa.Locale
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
