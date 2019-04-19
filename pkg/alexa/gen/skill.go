package gen

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// Flags for alexa.Privacy
const (
	FlagIsExportCompliant string = "IsExportCompliant"
	FlagContainsAds       string = "ContainsAds"
	FlagAllowsPurchases   string = "AllowsPurchases"
	FlagUsesPersonalInfo  string = "UsesPersonalInfo"
	FlagIsChildDirected   string = "IsChildDirected"
)

// SkillBuilder is a logical construct for the skill.
type SkillBuilder struct {
	registry          l10n.LocaleRegistry
	category          alexa.Category
	countries         []string
	skillInstructions string
	skillCountries    []string
	privacyFlags      map[string]bool
	locales           map[string]*SkillLocaleBuilder
	model             *ModelBuilder
	//permissions2 *SkillPermissionsBuilder

}

// NewSkillBuilder returns a new basic SkillBuilder
func NewSkillBuilder() *SkillBuilder {
	r := l10n.NewRegistry()
	s := &SkillBuilder{
		skillInstructions: l10n.KeySkillTestingInstructions,
		registry:          r,
		locales:           map[string]*SkillLocaleBuilder{},
		privacyFlags:      map[string]bool{},
	}
	return s
}

func (s *SkillBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *SkillBuilder {
	s.registry = registry
	return s
}

func (s *SkillBuilder) WithCategory(category alexa.Category) *SkillBuilder {
	s.category = category
	return s
}

func (s *SkillBuilder) WithCountries(countries []string) *SkillBuilder {
	s.countries = countries
	return s
}

func (s *SkillBuilder) WithTestingInstructions(instructions string) *SkillBuilder {
	s.skillInstructions = instructions
	return s
}

func (s *SkillBuilder) WithLocaleTestingInstructions(locale string, instructions string) *SkillBuilder {
	loc, err := s.registry.Resolve(locale)
	if err != nil {
		return s
	}
	loc.Set(s.skillInstructions, []string{instructions})
	return s
}

func (s *SkillBuilder) WithPrivacyFlag(flag string, value bool) *SkillBuilder {
	s.privacyFlags[flag] = value
	return s
}

func (s *SkillBuilder) AddLocale(locale string) *SkillLocaleBuilder {
	s.registry.Register(l10n.NewLocale(locale))

	lb := NewSkillLocaleBuilder(locale).
		WithLocaleRegistry(s.registry)
	s.locales[locale] = lb
	return lb
}

func (s *SkillBuilder) AddCountry(country string) *SkillBuilder {
	s.countries = append(s.countries, country)
	return s
}

func (s *SkillBuilder) AddCountries(cs []string) *SkillBuilder {
	for _, c := range cs {
		s.countries = append(s.countries, c)
	}
	return s
}

// Build builds an alexa.Skill object.
func (s *SkillBuilder) Build() (*alexa.Skill, error) {
	if s.registry == nil || len(s.registry.GetLocales()) == 0 {
		return nil, fmt.Errorf("No locales registered to build")
	}
	//
	if s.locales == nil || len(s.locales) == 0 {
		for n, _ := range s.registry.GetLocales() {
			s.locales[n] = NewSkillLocaleBuilder(n).
				WithLocaleRegistry(s.registry)

		}
	}

	dl := s.registry.GetDefault()

	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version:    "1.0",
			Publishing: alexa.Publishing{},
		},
	}
	skill.Manifest.Publishing.Category = s.category
	// TODO: ensure unique occurance
	if len(s.countries) > 0 {
		skill.Manifest.Publishing.Countries = s.countries
	} else {
		skill.Manifest.Publishing.Worldwide = true
	}
	skill.Manifest.Publishing.TestingInstructions = dl.Get(s.skillInstructions)

	// TODO: Permissions are required.
	skill.Manifest.Permissions = &[]alexa.Permission{}

	// PrivacyAndCompliance is required.
	skill.Manifest.Privacy = &alexa.Privacy{}
	if s.privacyFlags[FlagIsExportCompliant] {
		skill.Manifest.Privacy.IsExportCompliant = true
	}
	if s.privacyFlags[FlagContainsAds] {
		skill.Manifest.Privacy.ContainsAds = true
	}
	if s.privacyFlags[FlagAllowsPurchases] {
		skill.Manifest.Privacy.AllowsPurchases = true
	}
	if s.privacyFlags[FlagUsesPersonalInfo] {
		skill.Manifest.Privacy.UsesPersonalInfo = true
	}
	if s.privacyFlags[FlagIsChildDirected] {
		skill.Manifest.Privacy.IsChildDirected = true
	}

	// Add elements for every locale.
	skill.Manifest.Publishing.Locales = make(map[string]alexa.LocaleDef)
	skill.Manifest.Privacy.Locales = make(map[string]alexa.PrivacyLocaleDef)
	for n, _ := range s.registry.GetLocales() {
		if l, err := s.locales[n].BuildPublishingLocale(); err != nil {
			return nil, err
		} else {
			skill.Manifest.Publishing.Locales[n] = l
		}

		if l, err := s.locales[n].BuildPrivacyLocale(); err != nil {
			return nil, err
		} else {
			skill.Manifest.Privacy.Locales[n] = l
		}
	}

	return skill, nil
}

// BuildModels builds an alexa.Model for each locale
func (s *SkillBuilder) BuildModels() (map[string]*alexa.Model, error) {
	if s.model == nil {
		return nil, fmt.Errorf("No model to build")
	}
	return s.model.Build()
}

//////////////////////////////////

type SkillLocaleBuilder struct {
	registry         l10n.LocaleRegistry
	locale           string
	skillName        string
	skillDescription string
	skillSummary     string
	skillKeywords    string
	skillExamples    string
	skillSmallIcon   string
	skillLargeIcon   string
	skillPrivacyURL  string
	skillTermsURL    string
}

func NewSkillLocaleBuilder(locale string) *SkillLocaleBuilder {
	return &SkillLocaleBuilder{
		locale:           locale,
		registry:         l10n.NewRegistry(),
		skillName:        l10n.KeySkillName,
		skillDescription: l10n.KeySkillDescription,
		skillSummary:     l10n.KeySkillSummary,
		skillKeywords:    l10n.KeySkillKeywords,
		skillExamples:    l10n.KeySkillExamplePhrases,
		skillSmallIcon:   l10n.KeySkillSmallIconURI,
		skillLargeIcon:   l10n.KeySkillLargeIconURI,
		skillPrivacyURL:  l10n.KeySkillPrivacyPolicyURL,
		skillTermsURL:    l10n.KeySkillTermsOfUse,
	}
}

func (l *SkillLocaleBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *SkillLocaleBuilder {
	l.registry = registry
	return l
}

func (l *SkillLocaleBuilder) WithName(name string) *SkillLocaleBuilder {
	l.skillName = name
	return l
}
func (l *SkillLocaleBuilder) WithLocaleName(name string) *SkillLocaleBuilder {
	if loc, err := l.registry.Resolve(l.locale); err != nil {
		return l
	} else {
		loc.Set(l.skillName, []string{name})
	}
	return l
}
func (l *SkillLocaleBuilder) WithDescription(description string) *SkillLocaleBuilder {
	l.skillDescription = description
	return l
}

// TODO: add WithLocale(Description|Summary|...)
func (l *SkillLocaleBuilder) WithSummary(summary string) *SkillLocaleBuilder {
	l.skillSummary = summary
	return l
}
func (l *SkillLocaleBuilder) WithExamples(examples string) *SkillLocaleBuilder {
	l.skillExamples = examples
	return l
}
func (l *SkillLocaleBuilder) WithKeywords(keywords string) *SkillLocaleBuilder {
	l.skillKeywords = keywords
	return l
}
func (l *SkillLocaleBuilder) WithSmallIcon(smallicon string) *SkillLocaleBuilder {
	l.skillSmallIcon = smallicon
	return l
}
func (l *SkillLocaleBuilder) WithLargeIcon(largeicon string) *SkillLocaleBuilder {
	l.skillLargeIcon = largeicon
	return l
}
func (l *SkillLocaleBuilder) WithPrivacyURL(privacy string) *SkillLocaleBuilder {
	l.skillPrivacyURL = privacy
	return l
}
func (l *SkillLocaleBuilder) WithTermsURL(terms string) *SkillLocaleBuilder {
	l.skillTermsURL = terms
	return l
}

func (l *SkillLocaleBuilder) BuildPublishingLocale() (alexa.LocaleDef, error) {
	if loc, err := l.registry.Resolve(l.locale); err != nil {
		return alexa.LocaleDef{}, err
	} else {
		return alexa.LocaleDef{
			Name:         loc.Get(l.skillName),
			Description:  loc.Get(l.skillDescription),
			Summary:      loc.Get(l.skillSummary),
			Keywords:     loc.GetAll(l.skillKeywords),
			Examples:     loc.GetAll(l.skillExamples),
			SmallIconURI: loc.Get(l.skillSmallIcon),
			LargeIconURI: loc.Get(l.skillLargeIcon),
		}, nil
	}
}

func (l *SkillLocaleBuilder) BuildPrivacyLocale() (alexa.PrivacyLocaleDef, error) {
	if loc, err := l.registry.Resolve(l.locale); err != nil {
		return alexa.PrivacyLocaleDef{}, err
	} else {
		return alexa.PrivacyLocaleDef{
			PrivacyPolicyURL: loc.Get(l.skillPrivacyURL),
			TermsOfUse:       loc.Get(l.skillTermsURL),
		}, nil
	}
}
