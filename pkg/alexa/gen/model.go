package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// modelBuilder builds an alexa.Model instance for a locale
type modelBuilder struct {
	registry   l10n.LocaleRegistry
	invocation string
	delegation string
	intents    map[string]*modelIntentBuilder
	types      map[string]*modelTypeBuilder
	prompts    map[string]*modelPromptBuilder
	error      error
}

// NewModelBuilder returns an initialized modelBuilder.
func NewModelBuilder() *modelBuilder {
	return &modelBuilder{
		registry:   l10n.NewRegistry(),
		invocation: l10n.KeySkillInvocation,
		intents:    map[string]*modelIntentBuilder{},
		types:      map[string]*modelTypeBuilder{},
		prompts:    map[string]*modelPromptBuilder{},
	}
}

// WithLocaleRegistry passes a locale registry.
func (m *modelBuilder) WithLocaleRegistry(r l10n.LocaleRegistry) *modelBuilder {
	m.registry = r
	return m
}

// WithInvocation sets the lookup key for the invocation.
func (m *modelBuilder) WithInvocation(invocation string) *modelBuilder {
	m.invocation = invocation
	return m
}

// WithDelegationStrategy sets the model delegation strategy.
func (m *modelBuilder) WithDelegationStrategy(strategy string) *modelBuilder {
	m.delegation = strategy
	return m
}

// WithLocale creates and sets a new locale.
func (m *modelBuilder) WithLocale(locale string, invocation string) *modelBuilder {
	loc := l10n.NewLocale(locale)
	if err := m.registry.Register(loc); err != nil {
		m.error = err
		return m
	}
	loc.Set(m.invocation, []string{invocation})
	return m
}

// WithIntent creates and sets a new named intent.
func (m *modelBuilder) WithIntent(name string) *modelBuilder {
	i := NewModelIntentBuilder(name).
		WithLocaleRegistry(m.registry)
	m.intents[name] = i
	return m
}

// WithType creates and sets a new named type.
func (m *modelBuilder) WithType(name string) *modelBuilder {
	t := NewModelTypeBuilder(name).
		WithLocaleRegistry(m.registry)
	m.types[name] = t
	return m
}

// WithElicitationSlotPrompt creates and sets an elicitation prompt for the intent-slot.
func (m *modelBuilder) WithElicitationSlotPrompt(intent string, slot string) *modelBuilder {
	// intent and slot must exist!
	var sl *modelSlotBuilder
	for _, i := range m.intents {
		for _, s := range i.slots {
			if i.name == intent && s.name == slot {
				sl = s
				break
			}
		}
	}
	if sl == nil {
		m.error = fmt.Errorf("no matching intent slot: %s-%s", intent, slot)
		return m
	}

	p := NewElicitationPromptBuilder(intent, slot).
		WithLocaleRegistry(m.registry)
	m.prompts[p.id] = p

	// link slot to prompt
	sl.WithElicitationPrompt(p.id)
	return m
}

// WithConfirmationSlotPrompt creates and sets a confirmation prompt for the intent-slot.
func (m *modelBuilder) WithConfirmationSlotPrompt(intent string, slot string) *modelBuilder {
	// intent and slot must exist!
	var sl *modelSlotBuilder
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
	m.prompts[p.id] = p

	// link slot to prompt
	sl.WithConfirmationPrompt(p.id)
	return m
}

// Intent returns the named intent.
func (m *modelBuilder) Intent(name string) *modelIntentBuilder {
	return m.intents[name]
}

// Type returns the named type.
func (m *modelBuilder) Type(name string) *modelTypeBuilder {
	return m.types[name]
}

// ElicitationPrompt returns the elicitation prompt for the intent-slot.
func (m *modelBuilder) ElicitationPrompt(intent string, slot string) *modelPromptBuilder {
	pb := NewElicitationPromptBuilder(intent, slot)
	return m.prompts[pb.id]
}

// ConfirmationPrompt returns the confirmation prompt for the intent-slot.
func (m *modelBuilder) ConfirmationPrompt(intent string, slot string) *modelPromptBuilder {
	pb := NewConfirmationPromptBuilder(intent, slot)
	return m.prompts[pb.id]
}

// Build generates a Model for each locale.
func (m *modelBuilder) Build() (map[string]*alexa.Model, error) {
	if m.error != nil {
		return nil, m.error
	}
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

// BuildLocale generates a Model for the locale.
func (m *modelBuilder) BuildLocale(locale string) (*alexa.Model, error) {
	if m.error != nil {
		return nil, m.error
	}
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

type modelIntentBuilder struct {
	registry    l10n.LocaleRegistry
	name        string
	samplesName string
	slots       map[string]*modelSlotBuilder
	error       error
}

// NewModelIntentBuilder returns an initialized modelIntentBuilder.
func NewModelIntentBuilder(name string) *modelIntentBuilder {
	return &modelIntentBuilder{
		registry:    l10n.NewRegistry(),
		name:        name,
		samplesName: name + l10n.KeyPostfixSamples,
		slots:       map[string]*modelSlotBuilder{},
	}
}

// WithLocaleRegistry passes a locale registry.
func (i *modelIntentBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *modelIntentBuilder {
	i.registry = registry
	return i
}

// WithSamples overwrites the locale lookup key.
func (i *modelIntentBuilder) WithSamples(samplesName string) *modelIntentBuilder {
	i.samplesName = samplesName
	return i
}

// WithLocaleSamples sets the lookup key translations for a specific locale.
func (i *modelIntentBuilder) WithLocaleSamples(locale string, samples []string) *modelIntentBuilder {
	loc, err := i.registry.Resolve(locale)
	if err != nil {
		i.error = err
		return i
	}
	loc.Set(i.samplesName, samples)
	return i
}

// WithSlot creates and sets a named slot for the intent.
func (i *modelIntentBuilder) WithSlot(name string, typeName string) *modelIntentBuilder {
	sb := NewModelSlotBuilder(i.name, name, typeName).
		WithLocaleRegistry(i.registry)
	i.slots[name] = sb
	return i
}

// Slot returns a named slot of the intent.
func (i *modelIntentBuilder) Slot(name string) *modelSlotBuilder {
	return i.slots[name]
}

// BuildLanguageIntent generates a ModelIntent for the locale.
func (i *modelIntentBuilder) BuildLanguageIntent(locale string) (alexa.ModelIntent, error) {
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

// BuildDialogIntent generates a DialogIntent for the locale.
func (i *modelIntentBuilder) BuildDialogIntent(locale string) (alexa.DialogIntent, error) {
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

type modelSlotBuilder struct {
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

// NewModelSlotBuilder returns an initialized modelSlotBuilder.
func NewModelSlotBuilder(intent string, name string, typeName string) *modelSlotBuilder {
	return &modelSlotBuilder{
		registry:    l10n.NewRegistry(),
		intent:      intent,
		name:        name,
		typeName:    typeName,
		samplesName: intent + "_" + name + l10n.KeyPostfixSamples,
	}
}

// WithLocaleRegistry passes a locale registry.
func (s *modelSlotBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *modelSlotBuilder {
	s.registry = registry
	return s
}

// WithSamples set the lookup key for the slot samples.
func (s *modelSlotBuilder) WithSamples(samplesName string) *modelSlotBuilder {
	s.samplesName = samplesName
	return s
}

// WithLocaleSamples sets the translated slot samples for the locale.
func (s *modelSlotBuilder) WithLocaleSamples(locale string, samples []string) *modelSlotBuilder {
	loc, err := s.registry.Resolve(locale)
	if err != nil {
		return s
	}
	loc.Set(s.samplesName, samples)
	return s
}

// WithConfirmationPrompt requires confirmation and links to the prompt id.
func (s *modelSlotBuilder) WithConfirmationPrompt(id string) *modelSlotBuilder {
	s.withConfirmation = true
	s.confirmationPrompt = id
	return s
}

// WithElicitationPrompt requires elicitation and links to the prompt id.
func (s *modelSlotBuilder) WithElicitationPrompt(id string) *modelSlotBuilder {
	s.withElicitation = true
	s.elicitationPrompt = id
	return s
}

// WithIntentConfirmationPrompt does nothing.
func (s *modelSlotBuilder) WithIntentConfirmationPrompt(prompt string) *modelSlotBuilder {
	// TODO: WithIntentConfirmationPrompt - https://developer.amazon.com/docs/custom-skills/define-the-dialog-to-collect-and-confirm-required-information.html#intent-confirmation
	return s
}

// BuildIntentSlot generates a ModelSlot for the locale.
func (s *modelSlotBuilder) BuildIntentSlot(locale string) (alexa.ModelSlot, error) {
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

// BuildDialogSlot generates a DialogIntentSlot for the locale.
func (s *modelSlotBuilder) BuildDialogSlot(locale string) (alexa.DialogIntentSlot, error) {
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

// modelTypeBuilder
type modelTypeBuilder struct {
	registry   l10n.LocaleRegistry
	name       string
	valuesName string
}

// NewModelTypeBuilder returns an initialized modelTypeBuilder.
func NewModelTypeBuilder(name string) *modelTypeBuilder {
	return &modelTypeBuilder{
		registry:   l10n.NewRegistry(),
		name:       name,
		valuesName: name + l10n.KeyPostfixValues,
	}
}

// WithLocaleRegistry passes a locale registry.
func (t *modelTypeBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *modelTypeBuilder {
	t.registry = registry
	return t
}

// WithValues sets the lookup key for the type values.
func (t *modelTypeBuilder) WithValues(valuesName string) *modelTypeBuilder {
	t.valuesName = valuesName
	return t
}

// WithLocaleValues sets the translated values for the type.
func (t *modelTypeBuilder) WithLocaleValues(locale string, values []string) *modelTypeBuilder {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return t
	}
	loc.Set(t.valuesName, values)
	return t
}

// Build generates a ModelType.
func (t *modelTypeBuilder) Build(locale string) (alexa.ModelType, error) {
	loc, err := t.registry.Resolve(locale)
	if err != nil {
		return alexa.ModelType{}, err
	}
	var tvs []alexa.TypeValue
	for _, v := range loc.GetAll(t.valuesName) {
		tvs = append(tvs, alexa.TypeValue{Name: alexa.NameValue{Value: v}})
	}
	return alexa.ModelType{Name: t.name, Values: tvs}, nil

}

////////////////////////////////////////

type modelPromptBuilder struct {
	registry   l10n.LocaleRegistry
	intent     string
	slot       string
	promptType string
	id         string
	variations map[string]*promptVariationsBuilder
}

// NewElicitationPromptBuilder returns an initialized modelPromptBuilder for Elicitation.
func NewElicitationPromptBuilder(intent string, slot string) *modelPromptBuilder {
	return &modelPromptBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: "Elicit",
		id:         fmt.Sprintf("Elicit.Intent-%s.IntentSlot-%s", intent, slot),
		variations: map[string]*promptVariationsBuilder{},
	}
}

// NewConfirmationPromptBuilder returns an initialized modelPromptBuilder for Confirmation.
func NewConfirmationPromptBuilder(intent string, slot string) *modelPromptBuilder {
	return &modelPromptBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: "Confirm",
		id:         fmt.Sprintf("Confirm.Intent-%s.IntentSlot-%s", intent, slot),
		variations: map[string]*promptVariationsBuilder{},
	}
}

// WithLocaleRegistry passes a locale registry.
func (p *modelPromptBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *modelPromptBuilder {
	p.registry = registry
	return p
}

// WithVariation creates and sets variations for the varType.
func (p *modelPromptBuilder) WithVariation(varType string) *modelPromptBuilder {
	v := NewPromptVariations(p.intent, p.slot, p.promptType, varType).
		WithLocaleRegistry(p.registry)
	p.variations[varType] = v
	return p
}

// Variation returns the variations for the varType.
func (p *modelPromptBuilder) Variation(varType string) *promptVariationsBuilder {
	return p.variations[varType]
}

// BuildLocale generates a ModelPrompt for the locale.
func (p *modelPromptBuilder) BuildLocale(locale string) (alexa.ModelPrompt, error) {
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
		mp.Variations = append(mp.Variations, pv...)
	}
	return mp, nil
}

///////////////////////////////////////

type promptVariationsBuilder struct {
	registry   l10n.LocaleRegistry
	intent     string
	slot       string
	promptType string
	vars       map[string]string
	error      error
}

// NewPromptVariations returns an initialized builder with lookup key "$intent_$slot_$promptType_(Text|SSML)".
func NewPromptVariations(intent string, slot string, promptType string, varType string) *promptVariationsBuilder {
	t := l10n.KeyPostfixSSML
	if varType == "PlainText" {
		t = l10n.KeyPostfixText
	}
	return &promptVariationsBuilder{
		registry:   l10n.NewRegistry(),
		intent:     intent,
		slot:       slot,
		promptType: promptType,
		vars:       map[string]string{varType: fmt.Sprintf("%s_%s_%s%s", intent, slot, promptType, t)},
	}
}

// WithLocaleRegistry passes a locale registry.
func (v *promptVariationsBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *promptVariationsBuilder {
	v.registry = registry
	return v
}

// WithVariation sets the lookup key for the varType.
func (v *promptVariationsBuilder) WithVariation(varType string) *promptVariationsBuilder {
	t := l10n.KeyPostfixSSML
	if varType == "PlainText" {
		t = l10n.KeyPostfixText
	}
	v.vars[varType] = fmt.Sprintf("%s_%s_%s%s", v.intent, v.slot, v.promptType, t)
	return v
}

// WithTypeValue sets valueName as the lookup key for the varType.
func (v *promptVariationsBuilder) WithTypeValue(varType string, valueName string) *promptVariationsBuilder {
	v.vars[varType] = valueName
	return v
}

// WithLocaleTypeValue sets the values for the type of the locale.
func (v *promptVariationsBuilder) WithLocaleTypeValue(locale string, varType string, values []string) *promptVariationsBuilder {
	loc, err := v.registry.Resolve(locale)
	if err != nil {
		v.error = err
		return v
	}
	loc.Set(v.vars[varType], values)
	return v
}

// BuildLocale generates a PromptVariation for the locale.
func (v *promptVariationsBuilder) BuildLocale(locale string) ([]alexa.PromptVariation, error) {
	var vs []alexa.PromptVariation
	if v.error != nil {
		return vs, v.error
	}
	loc, err := v.registry.Resolve(locale)
	if err != nil {
		return vs, err
	}
	// only useful with content, can never happen as you must use NewPromptVariationsBuilder.
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
