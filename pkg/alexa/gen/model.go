package gen

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// ModelBuilder builds an alexa.Model instance for a locale
type ModelBuilder struct {
	//skillBuilder *SkillBuilder
	registry *l10n.Registry
	intents  []*ModelIntentBuilder
	types    []*ModelTypeBuilder
	//prompts   []*ModelPromptBuilder
}

func NewModelBuilder() *ModelBuilder {
	return &ModelBuilder{}
}

func (m *ModelBuilder) WithL10NRegistr(r *l10n.Registry) *ModelBuilder {
	m.registry = r
	return m
}

func (m *ModelBuilder) WithIntent(intent string) *ModelIntentBuilder {
	i := NewModelIntentBuilder(intent).
		WithRegistry(m.registry)
	m.intents = append(m.intents, i)
	return i
}

func (m *ModelBuilder) WithType(typeName string) *ModelTypeBuilder {
	t := NewModelTypeBuilder(typeName).
		WithRegistry(m.registry)
	m.types = append(m.types, t)
	return t
}

func (m *ModelBuilder) Build() map[string]alexa.Model {
	ams := make(map[string]alexa.Model)

	// build model for each locale registered
	for _, l := range m.registry.GetLocales() {
		// create basic model
		am := alexa.Model{
			Model: alexa.InteractionModel{
				Language: alexa.LanguageModel{
					Invocation: l.GetSnippet(l10n.KeySkillInvocation),
				},
			},
		}

		//var prompts = []prompt{}

		// add intents
		for _, i := range m.intents {
			am.Model.Language.Intents = append(
				am.Model.Language.Intents, i.BuildLanguageIntent(l.Name),
			)

			am.Model.Dialog.Intents = append(
				am.Model.Dialog.Intents, i.BuildDialogIntent(l.Name),
			)
		}
		ams[l.Name] = am
	}
	return ams
}

// ModelIntentBuilder
type ModelIntentBuilder struct {
	name        string
	registry    *l10n.Registry
	samplesName string
	slots       []*ModelSlotBuilder
}

func NewModelIntentBuilder(intent string) *ModelIntentBuilder {
	return &ModelIntentBuilder{name: intent}
}

func (i *ModelIntentBuilder) WithRegistry(registry *l10n.Registry) *ModelIntentBuilder {
	i.registry = registry
	return i
}

func (i *ModelIntentBuilder) WithSamples(samplesName string) *ModelIntentBuilder {
	i.samplesName = samplesName
	return i
}

func (i *ModelIntentBuilder) WithSlot(slotName string) *ModelSlotBuilder {
	sb := NewModelSlotBuilder(slotName).
		WithRegistry(i.registry)
	i.slots = append(i.slots, sb)
	return sb
}

func (i *ModelIntentBuilder) BuildLanguageIntent(locale string) alexa.ModelIntent {
	l, _ := i.registry.Resolve(locale)
	mi := alexa.ModelIntent{
		Name: i.name,
	}
	if i.samplesName != "" {
		mi.Samples = l.GetAllSnippets(l10n.Key(i.samplesName))
	}
	mss := &[]alexa.ModelSlot{}
	for _, s := range i.slots {
		*mss = append(*mss, s.Build(locale))
	}
	return mi
}

func (i *ModelIntentBuilder) BuildDialogIntent(locale string) alexa.DialogIntent {
	return alexa.DialogIntent{}
}

// ModelSlotBuilder
type ModelSlotBuilder struct {
	registry    *l10n.Registry
	slotName    string
	typeName    string
	samplesName string
}

func NewModelSlotBuilder(slotName string) *ModelSlotBuilder {
	return &ModelSlotBuilder{slotName: slotName}
}

func (s *ModelSlotBuilder) WithRegistry(registry *l10n.Registry) *ModelSlotBuilder {
	s.registry = registry
	return s
}

func (s *ModelSlotBuilder) WithType(typeName string) *ModelSlotBuilder {
	s.typeName = typeName
	return s
}

func (s *ModelSlotBuilder) WithSamples(samplesName string) *ModelSlotBuilder {
	s.samplesName = samplesName
	return s
}

func (s *ModelSlotBuilder) Build(locale string) alexa.ModelSlot {
	l, _ := s.registry.Resolve(locale)
	ms := alexa.ModelSlot{
		Name: s.slotName,
		Type: s.typeName,
	}
	if s.samplesName != "" {
		ms.Samples = l.GetAllSnippets(l10n.Key(s.samplesName))
	}
	return ms
}

// ModelTypeBuilder
type ModelTypeBuilder struct {
	name       string
	registry   *l10n.Registry
	valuesName string
}

func NewModelTypeBuilder(name string) *ModelTypeBuilder {
	return &ModelTypeBuilder{name: name}
}

func (t *ModelTypeBuilder) WithRegistry(registry *l10n.Registry) *ModelTypeBuilder {
	t.registry = registry
	return t
}
func (t *ModelTypeBuilder) WithValuesName(valuesName string) *ModelTypeBuilder {
	t.valuesName = valuesName
	return t
}

func (t *ModelTypeBuilder) Build(locale string) alexa.ModelType {
	l, _ := t.registry.Resolve(locale)
	var tv = []alexa.TypeValue{}
	for _, v := range l.GetAllSnippets(l10n.Key(t.valuesName)) {
		tv = append(tv, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	}
	return alexa.ModelType{Name: t.name, Values: tv}
}
