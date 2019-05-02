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

// SkillBuilder helps building the SKILL.json.
type SkillBuilder struct {
	error        error
	registry     l10n.LocaleRegistry
	category     alexa.Category
	countries    []string
	instructions string
	privacyFlags map[string]bool
	locales      map[string]*SkillLocaleBuilder
	model        *modelBuilder
	//permissions2 *SkillPermissionsBuilder
}

// NewSkillBuilder returns a new basic SkillBuilder.
func NewSkillBuilder() *SkillBuilder {
	return &SkillBuilder{
		instructions: l10n.KeySkillTestingInstructions,
		registry:     l10n.NewRegistry(),
		locales:      map[string]*SkillLocaleBuilder{},
		privacyFlags: map[string]bool{},
	}
}

// WithLocaleRegistry passes a LocaleRegistry instance to the builder.
func (s *SkillBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *SkillBuilder {
	s.registry = registry
	return s
}

// WithCategory sets the category of the skill.
func (s *SkillBuilder) WithCategory(category alexa.Category) *SkillBuilder {
	s.category = category
	return s
}

// WithCountries sets the list of countries the skill is available in.
func (s *SkillBuilder) WithCountries(countries []string) *SkillBuilder {
	s.countries = countries
	return s
}

// WithTestingInstructions sets the testing instructions lookup key of the skill.
func (s *SkillBuilder) WithTestingInstructions(instructions string) *SkillBuilder {
	s.instructions = instructions
	return s
}

// WithPrivacyFlag set a specific flag in the "privacyAndCompliance" section.
func (s *SkillBuilder) WithPrivacyFlag(flag string, value bool) *SkillBuilder {
	s.privacyFlags[flag] = value
	return s
}

// AddCountry add a single country to the list of available countries.
func (s *SkillBuilder) AddCountry(country string) *SkillBuilder {
	s.countries = append(s.countries, country)
	return s
}

// AddCountries adds a list of countries to the list of available countries.
func (s *SkillBuilder) AddCountries(cs []string) *SkillBuilder {
	s.countries = append(s.countries, cs...)
	return s
}

// WithLocale creates, registers locale and adds a new locale builder.
func (s *SkillBuilder) AddLocale(locale string, opts ...l10n.RegisterFunc) *SkillBuilder {
	if err := s.registry.Register(l10n.NewLocale(locale), opts...); err != nil {
		s.error = err
		return s
	}
	lb := NewSkillLocaleBuilder(locale).
		WithLocaleRegistry(s.registry)
	s.locales[locale] = lb
	return s
}

// WithDefaultLocale sets the default locale for the skill (used for
func (s *SkillBuilder) WithDefaultLocale(locale string) *SkillBuilder {
	if err := s.registry.SetDefault(locale); err != nil {
		s.error = err
	}
	return s
}

// WithDefaultLocaleTestingInstructions sets the actual testing instructions on the default locale.
func (s *SkillBuilder) WithDefaultLocaleTestingInstructions(instructions string) *SkillBuilder {
	dl := s.registry.GetDefault()
	if dl == nil {
		s.error = fmt.Errorf("no default locale registered")
		return s
	}
	dl.Set(s.instructions, []string{instructions})
	return s
}

// AddModel creates and returns a new modelBuilder attached to the skill.
func (s *SkillBuilder) WithModel() *SkillBuilder {
	s.model = NewModelBuilder().
		WithLocaleRegistry(s.registry)
	return s
}

// Locale returns the corresponding locale builder.
func (s *SkillBuilder) Locale(locale string) *SkillLocaleBuilder {
	if _, ok := s.locales[locale]; !ok {
		s.error = fmt.Errorf("no builder registered for locale '%s'", locale)
		return &SkillLocaleBuilder{}
	}
	return s.locales[locale]
}

// Model returns the corresponding model builder.
func (s *SkillBuilder) Model() *modelBuilder {
	if s.model == nil {
		s.error = fmt.Errorf("no model builder registered")
		return &modelBuilder{}
	}
	return s.model
}

// Build builds an alexa.Skill object.
func (s *SkillBuilder) Build() (*alexa.Skill, error) {
	if s.error != nil {
		return nil, s.error
	}
	if s.registry == nil || len(s.registry.GetLocales()) == 0 {
		return nil, fmt.Errorf("no locales registered to build")
	}
	// create SkillLocaleBuilders from registry
	if s.locales == nil || len(s.locales) == 0 {
		for n := range s.registry.GetLocales() {
			s.locales[n] = NewSkillLocaleBuilder(n).
				WithLocaleRegistry(s.registry)
		}
	}

	// get default locale
	dl := s.registry.GetDefault()
	if dl == nil {
		return nil, fmt.Errorf("no default locale defined")
	}

	skill := &alexa.Skill{
		Manifest: alexa.Manifest{
			Version:    "1.0",
			Publishing: alexa.Publishing{},
		},
	}
	if s.category == "" {
		return nil, fmt.Errorf("skill category is required")
	}
	skill.Manifest.Publishing.Category = s.category
	// TODO: ensure unique occurance?
	if len(s.countries) > 0 {
		skill.Manifest.Publishing.Countries = s.countries
	} else {
		skill.Manifest.Publishing.Worldwide = true
	}
	if dl.Get(s.instructions) == "" {
		return nil, fmt.Errorf("testing instructions are required (%s: %s)", dl.GetName(), s.instructions)
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

// BuildModels builds an alexa.Model for each locale.
func (s *SkillBuilder) BuildModels() (map[string]*alexa.Model, error) {
	if s.error != nil {
		return nil, s.error
	}
	if s.model == nil {
		return nil, fmt.Errorf("no model to build")
	}
	return s.model.Build()
}

// SkillLocaleBuilder represents elements for a specific locale.
type SkillLocaleBuilder struct {
	error            error
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

// NewSkillLocaleBuilder creates a new instance with default locale lookup keys.
func NewSkillLocaleBuilder(locale string) *SkillLocaleBuilder {
	r := l10n.NewRegistry()
	if err := r.Register(l10n.NewLocale(locale)); err != nil {
		return &SkillLocaleBuilder{
			locale: locale,
			error:  err,
		}
	}
	return &SkillLocaleBuilder{
		locale:           locale,
		registry:         r,
		skillName:        l10n.KeySkillName,
		skillSummary:     l10n.KeySkillSummary,
		skillDescription: l10n.KeySkillDescription,
		skillKeywords:    l10n.KeySkillKeywords,
		skillExamples:    l10n.KeySkillExamplePhrases,
		skillSmallIcon:   l10n.KeySkillSmallIconURI,
		skillLargeIcon:   l10n.KeySkillLargeIconURI,
		skillPrivacyURL:  l10n.KeySkillPrivacyPolicyURL,
		skillTermsURL:    l10n.KeySkillTermsOfUseURL,
	}
}

// WithLocaleRegistry passes a LocaleRegistry instance to the builder.
func (l *SkillLocaleBuilder) WithLocaleRegistry(registry l10n.LocaleRegistry) *SkillLocaleBuilder {
	l.registry = registry
	return l
}

// WithName sets the lookup key for the skill name.
func (l *SkillLocaleBuilder) WithName(name string) *SkillLocaleBuilder {
	l.skillName = name
	return l
}

// WithLocaleName sets the name of the skill for the locale.
func (l *SkillLocaleBuilder) WithLocaleName(name string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillName, []string{name})
	return l
}

// WithSummary sets the lookup key for the skill summary (shown with the skill overview).
func (l *SkillLocaleBuilder) WithSummary(summary string) *SkillLocaleBuilder {
	l.skillSummary = summary
	return l
}

// WithLocaleSummary sets the summary of the skill for the locale.
func (l *SkillLocaleBuilder) WithLocaleSummary(summary string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillSummary, []string{summary})
	return l
}

// WithDescription sets the lookup key for the skill description.
func (l *SkillLocaleBuilder) WithDescription(description string) *SkillLocaleBuilder {
	l.skillDescription = description
	return l
}

// WithLocaleDescription sets the description of the skill for the locale.
func (l *SkillLocaleBuilder) WithLocaleDescription(description string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillDescription, []string{description})
	return l
}

// WithExamples sets the lookup key for the skill example phrases.
func (l *SkillLocaleBuilder) WithExamples(examples string) *SkillLocaleBuilder {
	l.skillExamples = examples
	return l
}

// WithLocaleExamples sets the example phrases for the locale (max. 3).
func (l *SkillLocaleBuilder) WithLocaleExamples(examples []string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillExamples, examples)
	return l
}

// WithKeywords sets the lookup key for the skill keywords.
func (l *SkillLocaleBuilder) WithKeywords(keywords string) *SkillLocaleBuilder {
	l.skillKeywords = keywords
	return l
}

// WithLocaleKeywords sets the keywords for the locale (max. 3).
func (l *SkillLocaleBuilder) WithLocaleKeywords(keywords []string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillKeywords, keywords)
	return l
}

// WithSmallIcon sets the lookup key for the skill small icon URL.
func (l *SkillLocaleBuilder) WithSmallIcon(smallicon string) *SkillLocaleBuilder {
	l.skillSmallIcon = smallicon
	return l
}

// WithLocaleSmallIcon sets the small icon URL for the locale.
func (l *SkillLocaleBuilder) WithLocaleSmallIcon(smallicon string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillSmallIcon, []string{smallicon})
	return l
}

// WithLargeIcon sets the lookup key for the skill large icon URL.
func (l *SkillLocaleBuilder) WithLargeIcon(largeicon string) *SkillLocaleBuilder {
	l.skillLargeIcon = largeicon
	return l
}

// WithLocaleLargeIcon sets the large icon URL for the locale.
func (l *SkillLocaleBuilder) WithLocaleLargeIcon(largeicon string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillLargeIcon, []string{largeicon})
	return l
}

// WithPrivacyURL sets the lookup key for the privacy URL.
func (l *SkillLocaleBuilder) WithPrivacyURL(privacy string) *SkillLocaleBuilder {
	l.skillPrivacyURL = privacy
	return l
}

// WithLocalePrivacyURL sets the privacy URL for the locale.
func (l *SkillLocaleBuilder) WithLocalePrivacyURL(privacyURL string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillPrivacyURL, []string{privacyURL})
	return l
}

// WithTermsURL sets the lookup key for the terms URL.
func (l *SkillLocaleBuilder) WithTermsURL(terms string) *SkillLocaleBuilder {
	l.skillTermsURL = terms
	return l
}

// WithLocaleTermsURL sets the terms URL for the locale.
func (l *SkillLocaleBuilder) WithLocaleTermsURL(termsURL string) *SkillLocaleBuilder {
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		l.error = err
		return l
	}
	loc.Set(l.skillTermsURL, []string{termsURL})
	return l
}

// BuildPublishingLocale builds "publishingInformation" entry for the locale.
func (l *SkillLocaleBuilder) BuildPublishingLocale() (alexa.LocaleDef, error) {
	if l.error != nil {
		return alexa.LocaleDef{}, l.error
	}
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return alexa.LocaleDef{}, err
	}
	// sanity checks
	if loc.Get(l.skillName) == "" ||
		loc.Get(l.skillSummary) == "" ||
		loc.Get(l.skillDescription) == "" ||
		loc.Get(l.skillSmallIcon) == "" ||
		loc.Get(l.skillLargeIcon) == "" {
		return alexa.LocaleDef{}, fmt.Errorf(
			"skill requires a name, description, summary, small icon and large icon... but for '%s' at least one was empty",
			l.locale,
		)
	}
	if len(loc.GetAll(l.skillExamples)) > 3 {
		return alexa.LocaleDef{}, fmt.Errorf("only 3 examplePhrases are allowed (%s)", l.locale)
	}
	if len(loc.GetAll(l.skillKeywords)) > 3 {
		return alexa.LocaleDef{}, fmt.Errorf("only 3 keywords are allowed (%s)", l.locale)
	}
	return alexa.LocaleDef{
		Name:         loc.Get(l.skillName),
		Summary:      loc.Get(l.skillSummary),
		Description:  loc.Get(l.skillDescription),
		Keywords:     loc.GetAll(l.skillKeywords),
		Examples:     loc.GetAll(l.skillExamples),
		SmallIconURI: loc.Get(l.skillSmallIcon),
		LargeIconURI: loc.Get(l.skillLargeIcon),
	}, nil
}

// BuildPrivacyLocale builds "privacyAndCompliance" section for the locale.
func (l *SkillLocaleBuilder) BuildPrivacyLocale() (alexa.PrivacyLocaleDef, error) {
	if l.error != nil {
		return alexa.PrivacyLocaleDef{}, l.error
	}
	loc, err := l.registry.Resolve(l.locale)
	if err != nil {
		return alexa.PrivacyLocaleDef{}, err
	}
	p := alexa.PrivacyLocaleDef{
		PrivacyPolicyURL: loc.Get(l.skillPrivacyURL),
	}
	// seems not (yet) supported ?!?
	// Error: privacyAndCompliance.locales.en-US - object instance has properties which are not allowed by the schema: ["termsOfUse"]
	if loc.Get(l.skillTermsURL) != "" {
		return alexa.PrivacyLocaleDef{}, fmt.Errorf("'termsOfUse' makes Skill deployment fail! (%s)", l.locale)
		//p.TermsOfUse = loc.Get(l.skillTermsURL)
	}
	return p, nil
}
