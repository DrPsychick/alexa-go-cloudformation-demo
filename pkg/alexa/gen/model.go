package gen

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
)

const (
	KeyInvocation     string = "Invocation"
	KeyPostfixSamples        = "_samples"
	KeyPostfixValues         = "_values"
)

// LocaleRegistry is the interface for an l10n registry.
//type LocaleRegistry interface {
//	Resolve(locale string) (Locale, error)
//	GetDefault() Locale
//	GetLocales() []Locale
//}
//
//// Locale is the interface for a specific locale.
//type Locale interface {
//	GetName() string
//	Get(key string) string
//	GetAny(key string) string
//	GetAll(key string) []string
//}

// ModelBuilder builds an alexa.Model instance for a locale
type ModelBuilder struct {
	//skillBuilder *SkillBuilder
	intents []*ModelIntentBuilder
	types   []*ModelTypeBuilder
	//prompts   []*ModelPromptBuilder
	// with l10n package: pass a registry
	//registry LocaleRegistry

	// without l10n package: store key values
	// locale -> key -> strings
	locales map[string]map[string][]string
}

func NewModelBuilder() *ModelBuilder {
	mb := &ModelBuilder{}
	//mb.registry = l10n.NewRegistry()
	return mb
}

func (m *ModelBuilder) AddLocale(locale string, invocation string) *ModelBuilder {
	if m.locales == nil {
		m.locales = make(map[string]map[string][]string)
	}
	m.locales[locale] = make(map[string][]string)
	m.locales[locale][KeyInvocation] = []string{invocation}
	return m
}

//func (m *ModelBuilder) WithLocaleRegistry(r LocaleRegistry) *ModelBuilder {
//	m.registry = r
//	return m
//}

func (m *ModelBuilder) AddIntent(name string) *ModelIntentBuilder {
	i := NewModelIntentBuilder(name).
		WithLocales(m.locales) // pass on locales
	m.intents = append(m.intents, i)
	return i
}

//func (m *ModelBuilder) AddLocaleIntent(locale string, intent string) *ModelIntentBuilder {
//	m.ensureLocale(locale)
//
//	i := NewModelIntentBuilder(intent).
//		WithLocales(m.locales)
//	m.intents = append(m.intents, i)
//	return i
//}

func (m *ModelBuilder) AddType(name string) *ModelTypeBuilder {
	t := NewModelTypeBuilder(name).
		WithLocales(m.locales)
	m.types = append(m.types, t)
	return t
}

//func (m *ModelBuilder) AddLocaleType(locale string, typeName string) *ModelTypeBuilder {
//	t := NewModelTypeBuilder(typeName)
//	m.types = append(m.types, t)
//	return t
//}

func (m *ModelBuilder) Build() map[string]alexa.Model {
	ams := make(map[string]alexa.Model)

	//if m.registry != nil {
	//	// build model for each locale registered
	//	for _, l := range m.registry.GetLocales() {
	//		ams[l.GetName()] = m.BuildLocale(l.GetName())
	//	}
	//} else {
	for l, _ := range m.locales {
		ams[l] = m.BuildLocale(l)
	}
	//}
	return ams
}

func (m *ModelBuilder) BuildLocale(locale string) alexa.Model {
	// create basic model
	am := alexa.Model{
		Model: alexa.InteractionModel{
			Language: alexa.LanguageModel{
				Invocation: m.locales[locale][KeyInvocation][0],
			},
		},
	}

	mts := []alexa.ModelType{}
	for _, t := range m.types {
		mts = append(mts, t.Build(locale))
	}
	am.Model.Language.Types = mts

	//var prompts = []prompt{}

	// add intents
	for _, i := range m.intents {
		am.Model.Language.Intents = append(
			am.Model.Language.Intents, i.BuildLanguageIntent(locale),
		)

		//am.Model.Dialog.Intents = append(
		//	am.Model.Dialog.Intents, i.BuildDialogIntent(l.Name),
		//)
	}
	return am
}

func (m *ModelBuilder) GetLocales() map[string]map[string][]string {
	return m.locales
}

// ModelIntentBuilder
type ModelIntentBuilder struct {
	name  string
	slots []*ModelSlotBuilder
	// with l10n
	//registry    LocaleRegistry
	//samplesName string
	// without l10n
	locales map[string]map[string][]string
}

func NewModelIntentBuilder(name string) *ModelIntentBuilder {
	return &ModelIntentBuilder{name: name}
}

//func (i *ModelIntentBuilder) WithLocaleRegistry(registry LocaleRegistry) *ModelIntentBuilder {
//	i.registry = registry
//	return i
//}

//func (i *ModelIntentBuilder) AddLocale(locale string) *ModelIntentBuilder {
//	i.locale = locale
//	return i
//}

func (i *ModelIntentBuilder) WithLocales(locales map[string]map[string][]string) *ModelIntentBuilder {
	i.locales = locales
	return i
}

//func (i *ModelIntentBuilder) WithSamples(samplesName string) *ModelIntentBuilder {
//	i.samplesName = samplesName
//	return i
//}

func (i *ModelIntentBuilder) WithLocaleSamples(locale string, samples []string) *ModelIntentBuilder {
	i.locales[locale][i.name+KeyPostfixSamples] = samples
	return i
}

func (i *ModelIntentBuilder) AddSlot(name string) *ModelSlotBuilder {
	sb := NewModelSlotBuilder(name).
		WithLocales(i.locales) // pass on locales
	i.slots = append(i.slots, sb)
	return sb
}

func (i *ModelIntentBuilder) BuildLanguageIntent(locale string) alexa.ModelIntent {
	mi := alexa.ModelIntent{
		Name: i.name,
	}

	//if i.registry != nil && i.samplesName != "" {
	//	l, _ := i.registry.Resolve(locale)
	//	mi.Samples = l.GetAll(i.samplesName)
	//} else {
	mi.Samples = i.locales[locale][i.name+KeyPostfixSamples]
	//}

	mss := []alexa.ModelSlot{}
	for _, s := range i.slots {
		mss = append(mss, s.BuildLocaleIntentSlot(locale))
	}
	mi.Slots = mss

	return mi
}

func (i *ModelIntentBuilder) BuildDialogIntent(locale string) alexa.DialogIntent {
	return alexa.DialogIntent{}
}

// ModelSlotBuilder
type ModelSlotBuilder struct {
	//registry LocaleRegistry
	name     string
	typeName string

	//samplesName     string
	withElicitation bool

	locales map[string]map[string][]string
}

func NewModelSlotBuilder(name string) *ModelSlotBuilder {
	return &ModelSlotBuilder{name: name}
}

//func (s *ModelSlotBuilder) WithLocaleRegistry(registry LocaleRegistry) *ModelSlotBuilder {
//	s.registry = registry
//	return s
//}

func (s *ModelSlotBuilder) WithLocales(locales map[string]map[string][]string) *ModelSlotBuilder {
	s.locales = locales
	return s
}

func (s *ModelSlotBuilder) WithType(name string) *ModelSlotBuilder {
	s.typeName = name
	return s
}

//func (s *ModelSlotBuilder) WithSamples(samplesName string) *ModelSlotBuilder {
//	s.samplesName = samplesName
//	return s
//}

func (s *ModelSlotBuilder) WithElicitationPrompt() *ModelSlotBuilder {
	s.withElicitation = true
	return s
}

//func (s *ModelSlotBuilder) BuildIntentSlot(locale string) alexa.ModelSlot {
//	l, _ := s.registry.Resolve(locale)
//	ms := alexa.ModelSlot{
//		Name: s.name,
//		Type: s.typeName,
//	}
//	if s.samplesName != "" {
//		ms.Samples = l.GetAll(s.samplesName)
//	}
//	return ms
//}

func (s *ModelSlotBuilder) BuildLocaleIntentSlot(locale string) alexa.ModelSlot {
	ms := alexa.ModelSlot{
		Name: s.name,
		Type: s.typeName,
	}
	if len(s.locales[locale][s.name+KeyPostfixSamples]) > 0 {
		ms.Samples = s.locales[locale][s.name+KeyPostfixSamples]
	}
	return ms
}

// ModelTypeBuilder
type ModelTypeBuilder struct {
	name string
	//registry   LocaleRegistry
	//valuesName string
	//locale     string
	//values     []string

	locales map[string]map[string][]string
}

func NewModelTypeBuilder(name string) *ModelTypeBuilder {
	return &ModelTypeBuilder{name: name}
}

func (t *ModelTypeBuilder) WithLocales(locales map[string]map[string][]string) *ModelTypeBuilder {
	t.locales = locales
	return t
}

func (t *ModelTypeBuilder) WithLocaleValues(locale string, values []string) *ModelTypeBuilder {
	t.locales[locale][t.name+KeyPostfixValues] = values
	return t
}

//func (t *ModelTypeBuilder) WithLocaleRegistry(registry LocaleRegistry) *ModelTypeBuilder {
//	t.registry = registry
//	return t
//}
//func (t *ModelTypeBuilder) WithValuesName(valuesName string) *ModelTypeBuilder {
//	t.valuesName = valuesName
//	return t
//}

func (t *ModelTypeBuilder) Build(locale string) alexa.ModelType {
	//l, _ := t.registry.Resolve(locale)
	//var tv = []alexa.TypeValue{}
	//for _, v := range l.GetAll(t.valuesName) {
	//	tv = append(tv, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	//}
	//return alexa.ModelType{Name: t.name, Values: tv}
	tv := t.BuildValues(locale)
	return alexa.ModelType{Name: t.name, Values: tv}
}

func (t *ModelTypeBuilder) BuildValues(locale string) []alexa.TypeValue {
	tv := []alexa.TypeValue{}
	for _, v := range t.locales[locale][t.name+KeyPostfixValues] {
		tv = append(tv, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	}
	return tv
}
