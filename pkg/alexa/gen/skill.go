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
	registry     l10n.LocaleRegistry
	category     alexa.Category
	countries    []string
	instructions string
	privacyFlags map[string]bool
	locales      map[string]*SkillLocaleBuilder
	model        *ModelBuilder
	//permissions2 *SkillPermissionsBuilder
}

// NewSkillBuilder returns a new basic SkillBuilder
func NewSkillBuilder() *SkillBuilder {
	return &SkillBuilder{
		instructions: l10n.KeySkillTestingInstructions,
		registry:     l10n.NewRegistry(),
		locales:      map[string]*SkillLocaleBuilder{},
		privacyFlags: map[string]bool{},
	}
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
	s.instructions = instructions
	return s
}

func (s *SkillBuilder) WithPrivacyFlag(flag string, value bool) *SkillBuilder {
	s.privacyFlags[flag] = value
	return s
}

func (s *SkillBuilder) AddLocale(locale string) *SkillLocaleBuilder {
	if err := s.registry.Register(l10n.NewLocale(locale)); err != nil {
		return nil
	}

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
	s.countries = append(s.countries, cs...)
	return s
}

func (s *SkillBuilder) AddModel() *ModelBuilder {
	s.model = NewModelBuilder().
		WithLocaleRegistry(s.registry)
	return s.model
}

func (s *SkillBuilder) SetDefaultLocale(locale string) error {
	if err := s.registry.SetDefault(locale); err != nil {
		return err
	}
	return nil
}

func (s *SkillBuilder) SetDefaultLocaleTestingInstructions(instructions string) error {
	dl := s.registry.GetDefault()
	if dl == nil {
		return fmt.Errorf("No default locale registered!")
	}
	dl.Set(s.instructions, []string{instructions})
	return nil
}

// Build builds an alexa.Skill object.
func (s *SkillBuilder) Build() (*alexa.Skill, error) {
	if s.registry == nil || len(s.registry.GetLocales()) == 0 {
		return nil, fmt.Errorf("No locales registered to build")
	}
	// create SkillLocaleBuilders from registry
	if s.locales == nil || len(s.locales) == 0 {
		for n, _ := range s.registry.GetLocales() {
			s.locales[n] = NewSkillLocaleBuilder(n).
				WithLocaleRegistry(s.registry)
		}
	}

	// get default locale
	dl := s.registry.GetDefault()
	if dl == nil {
		return nil, fmt.Errorf("No default locale defined!")
	}

	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version:    "1.0",
			Publishing: alexa.Publishing{},
		},
	}
	if s.category == "" {
		return nil, fmt.Errorf("Skill category is required!")
	}
	skill.Manifest.Publishing.Category = s.category
	// TODO: ensure unique occurance?
	if len(s.countries) > 0 {
		skill.Manifest.Publishing.Countries = s.countries
	} else {
		skill.Manifest.Publishing.Worldwide = true
	}
	if dl.Get(s.instructions) == "" {
		return nil, fmt.Errorf("Testing instructions are required! (%s: %s)", dl.GetName(), s.instructions)
	}
	skill.Manifest.Publishing.TestingInstructions = dl.Get(s.instructions)

	// TODO: Permissions are required.
	skill.Manifest.Permissions = []alexa.Permission{}

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
	skill.Manifest.Publishing.Locales = map[string]alexa.LocaleDef{}
	skill.Manifest.Privacy.Locales = map[string]alexa.PrivacyLocaleDef{}
	for n, lb := range s.locales {
		l1, err := lb.BuildPublishingLocale()
		if err != nil {
			return nil, err
		}
		skill.Manifest.Publishing.Locales[n] = l1

		l2, err := lb.BuildPrivacyLocale()
		if err != nil {
			return nil, err
		}
		skill.Manifest.Privacy.Locales[n] = l2
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

// SkillLocaleBuilder builds elements for a specific locale.
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
	r := l10n.NewRegistry()
	if err := r.Register(l10n.NewLocale(locale)); err != nil {
		return nil
	}
	return &SkillLocaleBuilder{
		locale:           locale,
		registry:         r,
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
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillName, []string{name})
	return l
}

func (l *SkillLocaleBuilder) WithDescription(description string) *SkillLocaleBuilder {
	l.skillDescription = description
	return l
}
func (l *SkillLocaleBuilder) WithLocaleDescription(description string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillDescription, []string{description})
	return l
}

func (l *SkillLocaleBuilder) WithSummary(summary string) *SkillLocaleBuilder {
	l.skillSummary = summary
	return l
}
func (l *SkillLocaleBuilder) WithLocaleSummary(summary string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillSummary, []string{summary})
	return l
}

func (l *SkillLocaleBuilder) WithExamples(examples string) *SkillLocaleBuilder {
	l.skillExamples = examples
	return l
}
func (l *SkillLocaleBuilder) WithLocaleExamples(examples []string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillExamples, examples)
	return l
}

func (l *SkillLocaleBuilder) WithKeywords(keywords string) *SkillLocaleBuilder {
	l.skillKeywords = keywords
	return l
}
func (l *SkillLocaleBuilder) WithLocaleKeywords(keywords []string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillKeywords, keywords)
	return l
}

func (l *SkillLocaleBuilder) WithSmallIcon(smallicon string) *SkillLocaleBuilder {
	l.skillSmallIcon = smallicon
	return l
}
func (l *SkillLocaleBuilder) WithLocaleSmallIcon(smallicon string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillSmallIcon, []string{smallicon})
	return l
}

func (l *SkillLocaleBuilder) WithLargeIcon(largeicon string) *SkillLocaleBuilder {
	l.skillLargeIcon = largeicon
	return l
}
func (l *SkillLocaleBuilder) WithLocaleLargeIcon(largeicon string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillLargeIcon, []string{largeicon})
	return l
}

func (l *SkillLocaleBuilder) WithPrivacyURL(privacy string) *SkillLocaleBuilder {
	l.skillPrivacyURL = privacy
	return l
}
func (l *SkillLocaleBuilder) WithLocalePrivacyURL(privacyURL string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillPrivacyURL, []string{privacyURL})
	return l
}

func (l *SkillLocaleBuilder) WithTermsURL(terms string) *SkillLocaleBuilder {
	l.skillTermsURL = terms
	return l
}
func (l *SkillLocaleBuilder) WithLocaleTermsURL(termsURL string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return l
	}
	loc.Set(l.skillTermsURL, []string{termsURL})
	return l
}

func (l *SkillLocaleBuilder) BuildPublishingLocale() (alexa.LocaleDef, error) {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return alexa.LocaleDef{}, err
	}
	// sanity checks
	if loc.Get(l.skillName) == "" ||
		loc.Get(l.skillDescription) == "" ||
		loc.Get(l.skillSummary) == "" ||
		loc.Get(l.skillSmallIcon) == "" ||
		loc.Get(l.skillLargeIcon) == "" {
		return alexa.LocaleDef{}, fmt.Errorf(
			"Skill requires a name, description, summary, small icon and large icon... but for '%s' at least one was empty",
			l.locale,
		)
	}
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

func (l *SkillLocaleBuilder) BuildPrivacyLocale() (alexa.PrivacyLocaleDef, error) {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return alexa.PrivacyLocaleDef{}, err
	}
	return alexa.PrivacyLocaleDef{
		PrivacyPolicyURL: loc.Get(l.skillPrivacyURL),
		TermsOfUse:       loc.Get(l.skillTermsURL),
	}, nil
}
