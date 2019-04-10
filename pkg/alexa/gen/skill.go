package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// Skill is a logical construct for the skill.
type Skill struct {
	Category            alexa.Category
	DefaultLocale       string
	Locales             map[string]LocaleDef
	Countries           []alexa.Country
	Intents             []Intent
	Models              map[string]Model
	Types               []Type
	TestingInstructions string
	Privacy             Privacy
}

// NewSkill returns a new basic skill
func NewSkill() *Skill {
	s := &Skill{}
	// set sane defaults/allocate space
	s.Locales = make(map[string]LocaleDef)
	s.Countries = []alexa.Country{}
	s.Intents = []Intent{}
	s.Models = make(map[string]Model)
	s.Types = []Type{}

	// add default intents
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

func (s *Skill) SetDefaultLocale(locale string) {
	s.DefaultLocale = locale
}

func (s *Skill) AddLocale(l string, trans *l10n.Locale) {
	if len(s.Locales) == 0 {
		// ensure that a default is set
		if s.DefaultLocale == "" {
			s.DefaultLocale = l
		}
	}
	s.Locales[l] = LocaleDef{Translations: trans}
}

func (s *Skill) AddIntent(intent Intent) {
	s.Intents = append(s.Intents, intent)
}

func (s *Skill) AddIntentString(intent string) {
	s.Intents = append(s.Intents, NewIntent(intent))
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

// Build builds an alexa.Skill object.
func (s *Skill) Build() *alexa.Skill {
	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version: "1.0",
		},
	}

	// default is always set if at least one locale was defined.
	dl, _ := s.Locales[s.DefaultLocale]

	skill.Manifest.Publishing.Category = s.Category
	// TODO: ensure unique occurance
	skill.Manifest.Publishing.Countries = s.Countries

	if s.TestingInstructions != "" {
		skill.Manifest.Publishing.TestingInstructions = s.TestingInstructions
	} else {
		skill.Manifest.Publishing.TestingInstructions = dl.Translations.GetSnippet(l10n.KeySkillTestingInstructions)
	}

	// Permissions are required.
	skill.Manifest.Permissions = &[]alexa.Permission{}

	// TODO can we make this nicer?
	// PrivacyAndCompliance is required.
	skill.Manifest.Privacy = &alexa.Privacy{}
	if s.Privacy.IsExportCompliant {
		skill.Manifest.Privacy.IsExportCompliant = true
	}
	if s.Privacy.ContainsAds {
		skill.Manifest.Privacy.ContainsAds = true
	}

	// Add elements for every locale.
	for _, l := range s.Locales {
		l.BuildLocale(skill)
		l.BuildPrivacyLocale(skill)
	}

	return skill
}

// BuildModels builds an alexa.Model for each locale
func (s *Skill) BuildModels() map[string]*alexa.Model {
	models := make(map[string]*alexa.Model)
	for _, l := range s.Locales {
		models[l.Translations.Name] = s.BuildModel(&l)
	}
	return models
}

// BuildModel builds an alexa.Model for the given locale
func (s *Skill) BuildModel(locale *LocaleDef) *alexa.Model {
	model := &alexa.Model{
		Model: alexa.InteractionModel{
			Language: alexa.LanguageModel{
				Invocation: locale.Translations.GetSnippet(l10n.KeySkillInvocation),
			},
		},
	}

	// add Intents
	for _, i := range s.Intents {
		mi := alexa.ModelIntent{
			Name:    i.Name,
			Samples: []string{},
		}
		if sam := locale.Translations.GetIntent(l10n.Key(i.Name)).Samples; len(sam) > 0 {
			mi.Samples = sam
		}
		model.Model.Language.Intents = append(model.Model.Language.Intents, mi)
	}

	// add Types and Values
	model.Model.Language.Types = []alexa.ModelType{}
	for _, t := range s.Types {
		mt := alexa.ModelType{
			Name: string(t),
		}

		for _, t := range locale.Translations.GetAllSnippets(l10n.Key(t + "Values")) {
			mt.Values = append(mt.Values, alexa.TypeValue{
				Name: alexa.NameValue{Value: t},
			})
		}
		model.Model.Language.Types = append(model.Model.Language.Types, mt)
	}

	s.Models[locale.Translations.Name] = Model{
		Model: model,
	}
	return model

}

// ValidateTypes ensures that Intents only use Types defined in the Skill
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

// TODO: remove this indirection, use l10n.Locale directly?
// LocaleDef links skill locale with l10n.Locale to fetch translations.
type LocaleDef struct {
	Translations *l10n.Locale
}

// BuildLocale adds locale information to the alexa.Skill.
func (l *LocaleDef) BuildLocale(skill *alexa.Skill) {
	if skill.Manifest.Publishing.Locales == nil {
		skill.Manifest.Publishing.Locales = make(map[string]alexa.LocaleDef)
	}
	skill.Manifest.Publishing.Locales[l.Translations.Name] = alexa.LocaleDef{
		Name:         l.Translations.GetSnippet(l10n.KeySkillName),
		Description:  l.Translations.GetSnippet(l10n.KeySkillDescription),
		Summary:      l.Translations.GetSnippet(l10n.KeySkillSummary),
		Examples:     l.Translations.GetAllSnippets(l10n.KeySkillExamplePhrases),
		Keywords:     l.Translations.GetAllSnippets(l10n.KeySkillKeywords),
		SmallIconURI: l.Translations.GetSnippet(l10n.KeySkillSmallIconURI),
		LargeIconURI: l.Translations.GetSnippet(l10n.KeySkillLargeIconURI),
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
	skill.Manifest.Privacy.Locales[l.Translations.Name] = alexa.PrivacyLocaleDef{
		PrivacyPolicyURL: l.Translations.GetSnippet(l10n.KeySkillPrivacyPolicyURL),
		TermsOfUse:       l.Translations.GetSnippet(l10n.KeySkillTermsOfUse),
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
	Name string // specific to intent
	Type Type   // global for skill
}

func NewSlot(name string, t Type) Slot {
	return Slot{Name: name, Type: t}
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

// Privacy
type Privacy struct {
	IsExportCompliant bool
	ContainsAds       bool
	AllowsPurchases   bool
	UsesPersonalInfo  bool
	IsChildDirected   bool
}

func (p *Privacy) SetIsExportCompliant(b bool) {
	p.IsExportCompliant = b
}
func (p *Privacy) SetContainsAds(b bool) {
	p.ContainsAds = b
}

// TODO implement other functions
