package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// SkillBuilder is a logical construct for the skill.
type SkillBuilder struct {
	category            alexa.Category
	defaultLocale       string
	locales             map[string]LocaleDef
	countries           []alexa.Country
	intents             []Intent
	models              map[string]Model
	types               []Type
	testingInstructions string
	privacy             privacy
	modelDelegation     alexa.DialogDelegation
}

// NewSkillBuilder returns a new basic SkillBuilder
func NewSkillBuilder() *SkillBuilder {
	s := &SkillBuilder{}
	// set sane defaults/allocate space
	s.locales = make(map[string]LocaleDef)
	s.models = make(map[string]Model)

	// add default intents
	//s.AddIntentString(alexa.HelpIntent)
	//s.AddIntentString(alexa.CancelIntent)
	//s.AddIntentString(alexa.StopIntent)
	return s
}

func (s *SkillBuilder) WithCategory(category alexa.Category) *SkillBuilder {
	s.category = category
	return s
}

func (s *SkillBuilder) SetCategory(category alexa.Category) {
	s.category = category
}

func (s *SkillBuilder) SetTestingInstructions(instructions string) {
	s.testingInstructions = instructions
}

func (s *SkillBuilder) SetDefaultLocale(locale string) {
	s.defaultLocale = locale
}

func (s *SkillBuilder) SetModelDelegation(delegation alexa.DialogDelegation) {
	s.modelDelegation = delegation
}

func (s *SkillBuilder) AddLocale(l string, trans l10n.LocaleInstance) {
	if len(s.locales) == 0 {
		// ensure that a default is set
		if s.defaultLocale == "" {
			s.defaultLocale = l
		}
	}
	s.locales[l] = LocaleDef{Translations: trans}
}

func (s *SkillBuilder) AddIntent(intent Intent) {
	s.intents = append(s.intents, intent)
}

// AddIntentString creates an Intent from the string, adds it and returns a reference.
func (s *SkillBuilder) AddIntentString(intent string) *Intent {
	i := NewIntent(intent)
	s.intents = append(s.intents, i)
	return &i
}

func (s *SkillBuilder) AddType(t Type) {
	// TODO: validate rules!
	s.types = append(s.types, t)
}

func (s *SkillBuilder) AddTypeString(t string) {
	s.types = append(s.types, Type(t))
}

func (s *SkillBuilder) AddCountry(c alexa.Country) {
	s.countries = append(s.countries, c)
}

func (s *SkillBuilder) AddCountries(cs []alexa.Country) {
	for _, c := range cs {
		s.countries = append(s.countries, c)
	}
}

// Build builds an alexa.Skill object.
// TODO: return errors!
func (s *SkillBuilder) Build() (*alexa.Skill, error) {
	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version: "1.0",
		},
	}

	// default is always set if at least one locale was defined.
	dl, _ := s.locales[s.defaultLocale]

	skill.Manifest.Publishing.Category = s.category
	// TODO: ensure unique occurance
	skill.Manifest.Publishing.Countries = s.countries

	if s.testingInstructions != "" {
		skill.Manifest.Publishing.TestingInstructions = s.testingInstructions
	} else {
		skill.Manifest.Publishing.TestingInstructions = dl.Translations.Get(l10n.KeySkillTestingInstructions)
	}

	// Permissions are required.
	skill.Manifest.Permissions = &[]alexa.Permission{}

	// PrivacyAndCompliance is required.
	skill.Manifest.Privacy = &alexa.Privacy{}
	if s.privacy.flags[FlagIsExportCompliant] {
		skill.Manifest.Privacy.IsExportCompliant = true
	}
	if s.privacy.flags[FlagContainsAds] {
		skill.Manifest.Privacy.ContainsAds = true
	}
	if s.privacy.flags[FlagAllowsPurchases] {
		skill.Manifest.Privacy.AllowsPurchases = true
	}
	if s.privacy.flags[FlagUsesPersonalInfo] {
		skill.Manifest.Privacy.UsesPersonalInfo = true
	}
	if s.privacy.flags[FlagIsChildDirectred] {
		skill.Manifest.Privacy.IsChildDirected = true
	}

	// Add elements for every locale.
	for _, l := range s.locales {
		l.BuildLocale(skill)
		l.BuildPrivacyLocale(skill)
	}

	return skill, nil
}

// BuildModels builds an alexa.Model for each locale
func (s *SkillBuilder) BuildModels() (map[string]*alexa.Model, error) {
	var err error
	models := make(map[string]*alexa.Model)
	for _, l := range s.locales {
		models[l.Translations.GetName()], err = s.BuildModel(&l)
		if err != nil {
			return nil, fmt.Errorf("Could not build model for locale %s: %s", l.Translations.GetName(), err)
		}
	}
	return models, nil
}

// BuildModel builds an alexa.Model for the given locale
func (s *SkillBuilder) BuildModel(locale *LocaleDef) (*alexa.Model, error) {
	model := &alexa.Model{
		Model: alexa.InteractionModel{
			Language: alexa.LanguageModel{
				Invocation: locale.Translations.Get(l10n.KeySkillInvocation),
			},
		},
	}

	var prompts = []prompt{}

	// add Intents
	for _, i := range s.intents {
		samples := locale.Translations.GetAll(i.Name + l10n.KeyPostfixSamples)

		// create LanguageModel.Intent
		mi := alexa.ModelIntent{
			Name:    i.Name,
			Samples: []string{},
		}
		if len(samples) > 0 {
			mi.Samples = samples
		}

		//// loop over slots
		//var di_slots = []alexa.DialogIntentSlot{}
		//for _, sl := range i.Slots {
		//	ls := li.Slots[sl.Name]
		//
		//	if mi.Slots == nil {
		//		mi.Slots = []alexa.ModelSlot{}
		//	}
		//
		//	// create ModelSlot
		//	slot := alexa.ModelSlot{
		//		Name:    sl.Name,
		//		Type:    string(sl.Type),
		//		Samples: []string{},
		//	}
		//	if len(ls.Samples) > 0 {
		//		slot.Samples = ls.Samples
		//	}
		//	// add slot to ModelIntent
		//	mi.Slots = append(mi.Slots, slot)
		//
		//	// add slot DialogIntent - Elicitations
		//	pe := ls.PromptElicitations
		//	if len(pe) > 0 {
		//		p := prompt{Id: "Elicit.Intent-" + i.Name + ".IntentSlot-" + sl.Name}
		//		dis := alexa.DialogIntentSlot{
		//			Name: sl.Name,
		//			Type: string(sl.Type),
		//		}
		//		dis.Elicitation = true
		//		dis.Prompts = alexa.SlotPrompts{
		//			Elicitation: p.Id,
		//		}
		//		p.Variations = pe
		//		prompts = append(prompts, p)
		//		di_slots = append(di_slots, dis)
		//	}
		//
		//	// add slot DialogIntent - Confirmations
		//	pc := ls.PromptConfirmations
		//	if len(pc) > 0 {
		//		p := prompt{Id: "Confirm.Intent-" + i.Name + ".IntentSlot-" + sl.Name}
		//		dis := alexa.DialogIntentSlot{
		//			Name: sl.Name,
		//			Type: string(sl.Type),
		//		}
		//		dis.Confirmation = true
		//		dis.Prompts = alexa.SlotPrompts{
		//			Confirmation: p.Id,
		//		}
		//		p.Confirmations = pc
		//		prompts = append(prompts, p)
		//		di_slots = append(di_slots, dis)
		//	}
		//}
		// add LanguageModel.Intents
		model.Model.Language.Intents = append(model.Model.Language.Intents, mi)

		//// add Dialog.Intents
		//if len(di_slots) > 0 {
		//	// add Dialog
		//	if model.Model.Dialog == nil {
		//		model.Model.Dialog = &alexa.Dialog{
		//			Delegation: alexa.DialogDelegation(s.modelDelegation),
		//		}
		//	}
		//
		//	// create Dialog.Intent
		//	di := alexa.DialogIntent{
		//		Name:  i.Name,
		//		Slots: di_slots,
		//	}
		//	model.Model.Dialog.Intents = append(model.Model.Dialog.Intents, di)
		//}
	}

	// add Types and Values
	model.Model.Language.Types = []alexa.ModelType{}
	for _, t := range s.types {
		mt := alexa.ModelType{
			Name: string(t),
		}

		for _, t := range locale.Translations.GetAll(string(t + "Values")) {
			mt.Values = append(mt.Values, alexa.TypeValue{
				Name: alexa.NameValue{Value: t},
			})
		}
		model.Model.Language.Types = append(model.Model.Language.Types, mt)
	}

	// add Prompts
	if len(prompts) > 0 {
		model.Model.Prompts = &[]alexa.ModelPrompt{}
	}
	for _, p := range prompts {
		mp := alexa.ModelPrompt{
			Id:         p.Id,
			Variations: p.Variations,
		}
		*model.Model.Prompts = append(*model.Model.Prompts, mp)
	}

	// store reference to the model
	s.models[locale.Translations.GetName()] = Model{
		Model: model,
	}
	return model, nil

}

// ValidateTypes ensures that Intents only use Types defined in the Skill
func (s *SkillBuilder) ValidateTypes() error {
	var tm = make(map[string]bool)
	for _, t := range s.types {
		tm[string(t)] = true
	}
	for _, i := range s.intents {
		for _, sl := range i.Slots {
			if len(s.types) == 0 {
				return fmt.Errorf("No types defined in the skill!")
			}
			if _, ok := tm[string(sl.Type)]; !ok {
				return fmt.Errorf("Type validation error: intent slot %s uses type %s which is not defined in the skill", sl.Name, string(sl.Type))
			}
		}
	}
	return nil
}

// TODO: remove this indirection, use l10n.Locale directly?
// LocaleDef links skill locale with l10n.Locale to fetch translations.
type LocaleDef struct {
	Translations l10n.LocaleInstance
}

// BuildLocale adds locale information to the alexa.Skill.
func (l *LocaleDef) BuildLocale(skill *alexa.Skill) {
	if skill.Manifest.Publishing.Locales == nil {
		skill.Manifest.Publishing.Locales = make(map[string]alexa.LocaleDef)
	}
	skill.Manifest.Publishing.Locales[l.Translations.GetName()] = alexa.LocaleDef{
		Name:         l.Translations.Get(l10n.KeySkillName),
		Description:  l.Translations.Get(l10n.KeySkillDescription),
		Summary:      l.Translations.Get(l10n.KeySkillSummary),
		Examples:     l.Translations.GetAll(l10n.KeySkillExamplePhrases),
		Keywords:     l.Translations.GetAll(l10n.KeySkillKeywords),
		SmallIconURI: l.Translations.Get(l10n.KeySkillSmallIconURI),
		LargeIconURI: l.Translations.Get(l10n.KeySkillLargeIconURI),
	}
}

// BuildPrivacyLocale adds PrivacyAndCompliance locale information to the alexa.Skill
func (l *LocaleDef) BuildPrivacyLocale(skill *alexa.Skill) {
	if skill.Manifest.Privacy == nil { // TODO not needed, can we rely on it? (see above)
		skill.Manifest.Privacy = &alexa.Privacy{}
	}
	if skill.Manifest.Privacy.Locales == nil {
		skill.Manifest.Privacy.Locales = make(map[string]alexa.PrivacyLocaleDef)
	}
	skill.Manifest.Privacy.Locales[l.Translations.GetName()] = alexa.PrivacyLocaleDef{
		PrivacyPolicyURL: l.Translations.Get(l10n.KeySkillPrivacyPolicyURL),
		TermsOfUse:       l.Translations.Get(l10n.KeySkillTermsOfUse),
	}
}

// Intent
type Intent struct {
	Name  string
	Slots []Slot
}

// NewIntent returns a new intent with the given name
func NewIntent(name string) Intent {
	return Intent{Name: name}
}

func (i *Intent) AddSlot(slot Slot) {
	i.Slots = append(i.Slots, slot)
}

// Slot
type Slot struct {
	Name         string // specific to intent
	Type         Type   // global for skill
	Confirmation bool
	Elicitation  bool
}

func NewSlot(name string, t Type) Slot {
	return Slot{Name: name, Type: t, Confirmation: false, Elicitation: false}
}

// TODO: overkill? just use 'string'?
// Type
type Type string

func NewType(t string) Type {
	return Type(t)
}

// TODO: what for?
// Model
type Model struct {
	Model *alexa.Model
}

// Flags for alexa.Privacy
const (
	FlagIsExportCompliant string = "IsExportCompliant"
	FlagContainsAds       string = "ContainsAds"
	FlagAllowsPurchases   string = "AllowsPurchases"
	FlagUsesPersonalInfo  string = "UsesPersonalInfo"
	FlagIsChildDirectred  string = "IsChildDirected"
)

// privacy
type privacy struct {
	flags map[string]bool
}

func (s *SkillBuilder) SetPrivacyFlag(name string, b bool) {
	s.privacy.flags[name] = b
}

// TODO implement other functions

// prompt is used internally to build the models
type prompt struct {
	Id            string
	Variations    []alexa.PromptVariations
	Confirmations []alexa.PromptVariations
}
