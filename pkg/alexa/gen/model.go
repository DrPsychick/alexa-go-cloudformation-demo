package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// ModelBuilder builds an alexa.Model instance for a locale
type ModelBuilder struct {
	registry   l10n.LocaleRegistry
	invocation string
	delegation string
	intents    []*ModelIntentBuilder
	types      []*ModelTypeBuilder
	prompts    []*ModelPromptBuilder
}

func NewModelBuilder() *ModelBuilder {
	return &ModelBuilder{
		registry:   l10n.NewRegistry(),
		invocation: l10n.KeySkillInvocation,
	}
}

func (m *ModelBuilder) WithLocaleRegistry(r l10n.LocaleRegistry) *ModelBuilder {
	m.registry = r
	return m
}

func (m *ModelBuilder) WithInvocation(invocation string) *ModelBuilder {
	m.invocation = invocation
	return m
}

func (m *ModelBuilder) WithDelegationStrategy(strategy string) *ModelBuilder {
	m.delegation = strategy
	return m
}

func (m *ModelBuilder) AddLocale(locale string, invocation string) *ModelBuilder {
	loc := l10n.NewLocale(locale)
	if err := m.registry.Register(loc); err != nil {
		return nil
	}
	loc.Set(m.invocation, []string{invocation})
	return m
}

func (m *ModelBuilder) AddIntent(name string) *ModelIntentBuilder {
	i := NewModelIntentBuilder(name).
		WithLocaleRegistry(m.registry)
	m.intents = append(m.intents, i)
	return i
}

func (m *ModelBuilder) AddType(name string) *ModelTypeBuilder {
	t := NewModelTypeBuilder(name).
		WithLocaleRegistry(m.registry)
	m.types = append(m.types, t)
	return t
}

func (m *ModelBuilder) AddElicitationSlotPrompt(intent string, slot string) *ModelPromptBuilder {
	// intent and slot must exist!
	var sl *ModelSlotBuilder
	for _, i := range m.intents {
		for _, s := range i.slots {
			if i.name == intent && s.name == slot {
				sl = s
				break
			}
		}
	}
	if sl == nil {
		return nil
	}

	p := NewElicitationPromptBuilder(intent, slot).
		WithLocaleRegistry(m.registry)
	m.prompts = append(m.prompts, p)

	// link slot to prompt
	sl.WithElicitationPrompt(p.id)
	return p
}
func (m *ModelBuilder) AddConfirmationSlotPrompt(intent string, slot string) *ModelPromptBuilder {
	// intent and slot must exist!
	var sl *ModelSlotBuilder
	for _, i := range m.intents {
		for _, s := range i.slots {
			if i.name == intent && s.name == slot {
				sl = s
				break
			}
		}
	}
	if sl == nil {
		return nil
	}

	p := NewConfirmationPromptBuilder(intent, slot).
		WithLocaleRegistry(m.registry)
	m.prompts = append(m.prompts, p)

	// link slot to prompt
	sl.WithConfirmationPrompt(p.id)
	return p
}

func (m *ModelBuilder) Build() (map[string]*alexa.Model, error) {
	ams := make(map[string]*alexa.Model)

	// build model for each locale registered
	for _, l := range m.registry.GetLocales() {
		m, err := m.BuildLocale(l.GetName())
		if err != nil {
			return nil, err
		}
		ams[l.GetName()] = m
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
				Invocation: loc.Get(m.invocation),
			},
		},
	}

	var mts []alexa.ModelType
	for _, t := range m.types {
		mt, err := t.Build(locale)
		if err != nil {
			return &alexa.Model{}, err
		}
		mts = append(mts, mt)
	}
	am.Model.Language.Types = mts

	// add prompts - only if we have intents with slots
	// TODO: "Add...Prompt" should not fail, it should fail during build()!
	am.Model.Prompts = []alexa.ModelPrompt{}
	for _, p := range m.prompts {
		mp, err := p.BuildLocale(locale)
		if err != nil {
			return &alexa.Model{}, err
		}
		am.Model.Prompts = append(am.Model.Prompts, mp)
	}

	// add intents
	// TODO: ensure that slot types are defined, if not: fail
	am.Model.Dialog = &alexa.Dialog{}
	if m.delegation != "" {
		am.Model.Dialog.Delegation = m.delegation
	}
	for _, i := range m.intents {
		li, err := i.BuildLanguageIntent(locale)
		if err != nil {
			return &alexa.Model{}, err
		}
		am.Model.Language.Intents = append(am.Model.Language.Intents, li)

		// only needed for intents with slots
		if len(i.slots) > 0 {
			di, err := i.BuildDialogIntent(locale)
			if err != nil {
				return &alexa.Model{}, err
			}
			am.Model.Dialog.Intents = append(am.Model.Dialog.Intents, di)
		}
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
	return &ModelIntentBuilder{
		registry:    l10n.NewRegistry(),
		name:        name,
		samplesName: name + l10n.KeyPostfixSamples,
	}
}

func (i *ModelIntentBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelIntentBuilder {
	i.registry = registry
	return i
}

// WithSamples overwrites the locale lookup key.
func (i *ModelIntentBuilder) WithSamples(samplesName string) *ModelIntentBuilder {
	i.samplesName = samplesName
	return i
}

// WithLocaleSamples sets the lookup key translations for a specific locale.
func (i *ModelIntentBuilder) WithLocaleSamples(locale string, samples []string) *ModelIntentBuilder {
	loc, err := i.registry.Resolve(locale)
	if err != nil {
		return i
	}
	loc.Set(i.samplesName, samples)
	return i
}

func (i *ModelIntentBuilder) AddSlot(name string, typeName string) *ModelSlotBuilder {
	sb := NewModelSlotBuilder(i.name, name, typeName).
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

	var mss []alexa.ModelSlot
	for _, s := range i.slots {
		is, err := s.BuildIntentSlot(locale)
		if err != nil {
			return alexa.ModelIntent{}, err
		}
		mss = append(mss, is)
	}
	mi.Slots = mss

	return mi, nil
}

func (i *ModelIntentBuilder) BuildDialogIntent(locale string) (alexa.DialogIntent, error) {
	di := alexa.DialogIntent{
		Name: i.name,
		// TODO: Confirmation, Delegation, ...
	}
	var dis []alexa.DialogIntentSlot
	for _, s := range i.slots {
		ds, err := s.BuildDialogSlot(locale)
		if err != nil {
			return alexa.DialogIntent{}, err
		}
		dis = append(dis, ds)
	}
	di.Slots = dis
	return di, nil
}

////////////////////////////////////

// ModelSlotBuilder
type ModelSlotBuilder struct {
	registry           l10n.LocaleRegistry
	intent             string
	name               string
	typeName           string
	samplesName        string
	withConfirmation   bool
	withElicitation    bool
	elicitationPrompt  string
	confirmationPrompt string
}

func NewModelSlotBuilder(intent string, name string, typeName string) *ModelSlotBuilder {
	return &ModelSlotBuilder{
		registry:    l10n.NewRegistry(),
		intent:      intent,
		name:        name,
		typeName:    typeName,
		samplesName: intent + "_" + name + l10n.KeyPostfixSamples,
	}
}

func (s *ModelSlotBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelSlotBuilder {
	s.registry = registry
	return s
}

func (s *ModelSlotBuilder) WithSamples(samplesName string) *ModelSlotBuilder {
	s.samplesName = samplesName
	return s
}
func (s *ModelSlotBuilder) WithLocaleSamples(locale string, samples []string) *ModelSlotBuilder {
	loc, err := s.registry.Resolve(locale)
	if err != nil {
		return s
	}
	loc.Set(s.samplesName, samples)
	return s
}

func (s *ModelSlotBuilder) WithConfirmationPrompt(id string) *ModelSlotBuilder {
	s.withConfirmation = true
	s.confirmationPrompt = id
	return s
}

func (s *ModelSlotBuilder) WithElicitationPrompt(id string) *ModelSlotBuilder {
	s.withElicitation = true
	s.elicitationPrompt = id
	return s
}

func (s *ModelSlotBuilder) WithIntentConfirmationPrompt(prompt string) *ModelSlotBuilder {
	// TODO: WithIntentConfirmationPrompt - https://developer.amazon.com/docs/custom-skills/define-the-dialog-to-collect-and-confirm-required-information.html#intent-confirmation
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
	ms.Samples = l.GetAll(s.samplesName)
	return ms, nil
}

func (s *ModelSlotBuilder) BuildDialogSlot(locale string) (alexa.DialogIntentSlot, error) {
	if _, err := s.registry.Resolve(locale); err != nil {
		return alexa.DialogIntentSlot{}, err
	}
	ds := alexa.DialogIntentSlot{
		Name:         s.name,
		Type:         s.typeName,
		Confirmation: s.withConfirmation,
		Elicitation:  s.withElicitation,
	}
	if s.confirmationPrompt != "" {
		ds.Prompts.Confirmation = s.confirmationPrompt
	}
	if s.elicitationPrompt != "" {
		ds.Prompts.Elicitation = s.elicitationPrompt
	}
	return ds, nil
}

/////////////////////////////////////////////

// ModelTypeBuilder
type ModelTypeBuilder struct {
	registry   l10n.LocaleRegistry
	name       string
	valuesName string
}

func NewModelTypeBuilder(name string) *ModelTypeBuilder {
	return &ModelTypeBuilder{
		registry:   l10n.NewRegistry(),
		name:       name,
		valuesName: name + l10n.KeyPostfixValues,
	}
}

func (t *ModelTypeBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelTypeBuilder {
	t.registry = registry
	return t
}

func (t *ModelTypeBuilder) WithValues(valuesName string) *ModelTypeBuilder {
	t.valuesName = valuesName
	return t
}

func (t *ModelTypeBuilder) WithLocaleValues(locale string, values []string) *ModelTypeBuilder {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return t
	}
	loc.Set(t.valuesName, values)
	return t
}

func (t *ModelTypeBuilder) Build(locale string) (alexa.ModelType, error) {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return alexa.ModelType{}, err
	}
	var tv []alexa.TypeValue
	for _, v := range loc.GetAll(t.valuesName) {
		tv = append(tv, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	}
	return alexa.ModelType{Name: t.name, Values: tv}, nil

}

////////////////////////////////////////

type ModelPromptBuilder struct {
	registry   l10n.LocaleRegistry
	intent     string
	slot       string
	promptType string
	id         string
	variations []*PromptVariationsBuilder
}

func NewElicitationPromptBuilder(intent string, slot string) *ModelPromptBuilder {
	return &ModelPromptBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: "Elicit",
		id:         fmt.Sprintf("Elicit.Intent-%s.IntentSlot-%s", intent, slot),
	}
}

func NewConfirmationPromptBuilder(intent string, slot string) *ModelPromptBuilder {
	return &ModelPromptBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: "Confirm",
		id:         fmt.Sprintf("Confirm.Intent-%s.IntentSlot-%s", intent, slot),
	}
}

func (p *ModelPromptBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *ModelPromptBuilder {
	p.registry = registry
	return p
}

func (p *ModelPromptBuilder) AddVariation(varType string) *PromptVariationsBuilder {
	v := NewPromptVariations(p.intent, p.slot, p.promptType, varType).
		WithLocaleRegistry(p.registry)
	p.variations = append(p.variations, v)
	return v
}

func (p *ModelPromptBuilder) BuildLocale(locale string) (alexa.ModelPrompt, error) {
	if len(p.variations) == 0 {
		return alexa.ModelPrompt{}, fmt.Errorf(
			"prompt '%s' requires variations (%s)",
			p.id, locale)
	}
	mp := alexa.ModelPrompt{
		Id:         p.id,
		Variations: []alexa.PromptVariation{},
	}
	for _, v := range p.variations {
		pv, err := v.BuildLocale(locale)
		if err != nil {
			return alexa.ModelPrompt{}, err
		}
		mp.Variations = pv
	}
	return mp, nil
}

///////////////////////////////////////

// PromptVariationsBuilder builds a list of variations for a specific prompt.
type PromptVariationsBuilder struct {
	registry   l10n.LocaleRegistry
	intent     string
	slot       string
	promptType string
	vars       map[string]string
}

func NewPromptVariations(intent string, slot string, promptType string, varType string) *PromptVariationsBuilder {
	t := l10n.KeyPostfixSSML
	if varType == "PlainText" {
		t = l10n.KeyPostfixText
	}
	return &PromptVariationsBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: promptType,
		vars:       map[string]string{varType: fmt.Sprintf("%s_%s_%s%s", intent, slot, promptType, t)},
	}
}

func (v *PromptVariationsBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *PromptVariationsBuilder {
	v.registry = registry
	return v
}

func (v *PromptVariationsBuilder) AddVariation(varType string) *PromptVariationsBuilder {
	t := l10n.KeyPostfixSSML
	if varType == "PlainText" {
		t = l10n.KeyPostfixText
	}
	v.vars[varType] = fmt.Sprintf("%s_%s_%s%s", v.intent, v.slot, v.promptType, t)
	return v
}

func (v *PromptVariationsBuilder) WithTypeValue(varType string, valueName string) *PromptVariationsBuilder {
	v.vars[varType] = valueName
	return v
}
func (v *PromptVariationsBuilder) WithLocaleTypeValue(locale string, varType string, values []string) *PromptVariationsBuilder {
	loc, err := v.registry.Resolve(locale)
	if err != nil {
		return v
	}
	loc.Set(v.vars[varType], values)
	return v
}

func (v *PromptVariationsBuilder) BuildLocale(locale string) ([]alexa.PromptVariation, error) {
	var vs []alexa.PromptVariation
	loc, err := v.registry.Resolve(locale)
	if err != nil {
		return vs, err
	}
	// only useful with content, can never happen if ppl. use NewPromptVariationsBuilder :)
	if len(v.vars) == 0 {
		return []alexa.PromptVariation{}, fmt.Errorf(
			"prompt requires variations (%s: %s-%s-%s)",
			locale, v.promptType, v.intent, v.slot)
	}
	// loop over variation types
	for t, n := range v.vars {
		for _, val := range loc.GetAll(n) {
			vs = append(vs, alexa.PromptVariation{
				Type:  t,
				Value: val,
			})
		}
	}
	if len(vs) == 0 {
		return []alexa.PromptVariation{}, fmt.Errorf(
			"prompt requires variations with values (%s: %s-%s-%s)",
			locale, v.promptType, v.intent, v.slot)
	}
	return vs, nil
}
