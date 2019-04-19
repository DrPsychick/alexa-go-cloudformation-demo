package gen

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// ModelBuilder builds an alexa.Model instance for a locale
type ModelBuilder struct {
	registry l10n.LocaleRegistry
	intents  []*ModelIntentBuilder
	types    []*ModelTypeBuilder
	//prompts   []*ModelPromptBuilder
}

func NewModelBuilder() *ModelBuilder {
	mb := &ModelBuilder{}
	mb.registry = l10n.NewRegistry()
	return mb
}

func (m *ModelBuilder) AddLocale(locale string, invocation string) *ModelBuilder {
	loc := l10n.NewLocale(locale)
	m.registry.Register(loc)
	loc.TextSnippets[l10n.KeySkillInvocation] = []string{invocation}
	return m
}

func (m *ModelBuilder) WithLocaleRegistry(r l10n.LocaleRegistry) *ModelBuilder {
	m.registry = r
	return m
}

func (m *ModelBuilder) AddIntent(name string) *ModelIntentBuilder {
	i := NewModelIntentBuilder(name).
		WithLocaleRegistry(m.registry) // pass on locales
	m.intents = append(m.intents, i)
	return i
}

func (m *ModelBuilder) AddType(name string) *ModelTypeBuilder {
	t := NewModelTypeBuilder(name).
		WithLocaleRegistry(m.registry)
	m.types = append(m.types, t)
	return t
}

func (m *ModelBuilder) Build() (map[string]*alexa.Model, error) {
	ams := make(map[string]*alexa.Model)

	// build model for each locale registered
	for _, l := range m.registry.GetLocales() {
		if m, err := m.BuildLocale(l.GetName()); err != nil {
			return nil, err
		} else {
			ams[l.GetName()] = m
		}
	}
	return ams, nil
}

func (m *ModelBuilder) BuildLocale(locale string) (*alexa.Model, error) {
	loc, err := m.registry.Resolve(locale)
	if err != nil {
		return &alexa.Model{}, err
	}
	// create basic model
	am := &alexa.Model{
		Model: alexa.InteractionModel{
			Language: alexa.LanguageModel{
				Invocation: loc.Get(l10n.KeySkillInvocation),
			},
		},
	}

	mts := []alexa.ModelType{}
	for _, t := range m.types {
		if mt, err := t.Build(locale); err != nil {
			return &alexa.Model{}, err
		} else {
			mts = append(mts, mt)
		}
	}
	am.Model.Language.Types = mts

	//var prompts = []prompt{}

	// add intents
	for _, i := range m.intents {
		if i, err := i.BuildLanguageIntent(locale); err != nil {
			return &alexa.Model{}, err
		} else {
			am.Model.Language.Intents = append(am.Model.Language.Intents, i)
		}

		//am.Model.Dialog.Intents = append(
		//	am.Model.Dialog.Intents, i.BuildDialogIntent(locale),
		//)
	}
	return am, nil
}

///////////////////////////////////////////////////////

// ModelIntentBuilder
type ModelIntentBuilder struct {
	registry    l10n.LocaleRegistry
	name        string
	samplesName string
	slots       []*ModelSlotBuilder
}

func NewModelIntentBuilder(name string) *ModelIntentBuilder {
	return &ModelIntentBuilder{name: name}
}

func (i *ModelIntentBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelIntentBuilder {
	i.registry = registry
	return i
}

func (i *ModelIntentBuilder) WithSamples(samplesName string) *ModelIntentBuilder {
	i.samplesName = samplesName
	return i
}

func (i *ModelIntentBuilder) WithLocaleSamples(locale string, samples []string) *ModelIntentBuilder {
	loc, err := i.registry.Resolve(locale)
	if err != nil {
		return i
	}
	if i.samplesName == "" {
		i.samplesName = i.name + l10n.KeyPostfixSamples
	}
	loc.Set(i.samplesName, samples)
	return i
}

func (i *ModelIntentBuilder) AddSlot(name string) *ModelSlotBuilder {
	sb := NewModelSlotBuilder(name).
		WithLocaleRegistry(i.registry)
	i.slots = append(i.slots, sb)
	return sb
}

func (i *ModelIntentBuilder) BuildLanguageIntent(locale string) (alexa.ModelIntent, error) {
	loc, err := i.registry.Resolve(locale)
	if err != nil {
		return alexa.ModelIntent{}, err
	}

	mi := alexa.ModelIntent{
		Name:    i.name,
		Samples: loc.GetAll(i.samplesName),
	}

	mss := []alexa.ModelSlot{}
	for _, s := range i.slots {
		if s, err := s.BuildIntentSlot(locale); err != nil {
			return alexa.ModelIntent{}, err
		} else {
			mss = append(mss, s)
		}
	}
	mi.Slots = mss

	return mi, nil
}

func (i *ModelIntentBuilder) BuildDialogIntent(locale string) alexa.DialogIntent {
	di := alexa.DialogIntent{
		Name: i.name,
	}
	dis := []alexa.DialogIntentSlot{}
	for _, s := range i.slots {
		dis = append(dis, s.BuildDialogSlot(locale))
	}
	di.Slots = dis
	return di
}

////////////////////////////////////

// ModelSlotBuilder
type ModelSlotBuilder struct {
	registry         l10n.LocaleRegistry
	name             string
	typeName         string
	samplesName      string
	withConfirmation bool
	withElicitation  bool
}

func NewModelSlotBuilder(name string) *ModelSlotBuilder {
	return &ModelSlotBuilder{name: name}
}

func (s *ModelSlotBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelSlotBuilder {
	s.registry = registry
	return s
}

func (s *ModelSlotBuilder) WithType(name string) *ModelSlotBuilder {
	s.typeName = name
	return s
}

func (s *ModelSlotBuilder) WithSamples(samplesName string) *ModelSlotBuilder {
	s.samplesName = samplesName
	return s
}

func (s *ModelSlotBuilder) WithElicitationPrompt() *ModelSlotBuilder {
	s.withElicitation = true
	return s
}

func (s *ModelSlotBuilder) BuildIntentSlot(locale string) (alexa.ModelSlot, error) {
	l, err := s.registry.Resolve(locale)
	if err != nil {
		return alexa.ModelSlot{}, err
	}
	ms := alexa.ModelSlot{
		Name: s.name,
		Type: s.typeName,
	}
	if s.samplesName != "" {
		ms.Samples = l.GetAll(s.samplesName)
	}
	return ms, nil
}

func (s *ModelSlotBuilder) BuildDialogSlot(locale string) alexa.DialogIntentSlot {
	ds := alexa.DialogIntentSlot{
		Name:         s.name,
		Type:         s.typeName,
		Confirmation: s.withConfirmation,
		Elicitation:  s.withElicitation,
	}
	return ds
}

/////////////////////////////////////////////

// ModelTypeBuilder
type ModelTypeBuilder struct {
	registry   l10n.LocaleRegistry
	name       string
	valuesName string
}

func NewModelTypeBuilder(name string) *ModelTypeBuilder {
	return &ModelTypeBuilder{name: name}
}

func (t *ModelTypeBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelTypeBuilder {
	t.registry = registry
	return t
}

func (t *ModelTypeBuilder) WithValuesName(valuesName string) *ModelTypeBuilder {
	t.valuesName = valuesName
	return t
}

func (t *ModelTypeBuilder) WithLocaleValues(locale string, values []string) *ModelTypeBuilder {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return t
	}
	if t.valuesName == "" {
		t.valuesName = t.name + l10n.KeyPostfixValues
	}
	loc.Set(t.valuesName, values)
	return t
}

func (t *ModelTypeBuilder) Build(locale string) (alexa.ModelType, error) {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return alexa.ModelType{}, err
	}
	var tv = []alexa.TypeValue{}
	for _, v := range loc.GetAll(t.valuesName) {
		tv = append(tv, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	}
	return alexa.ModelType{Name: t.name, Values: tv}, nil

}
