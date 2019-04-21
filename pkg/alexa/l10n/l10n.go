package l10n

import (
	"fmt"
	"math/rand"
	"time"
)

// Default keys
const (
	KeySkillName                string = "SKILL_Name"
	KeySkillDescription         string = "SKILL_Description"
	KeySkillSummary             string = "SKILL_Summary"
	KeySkillExamplePhrases      string = "SKILL_ExamplePhrases"
	KeySkillKeywords            string = "SKILL_Keywords"
	KeySkillSmallIconURI        string = "SKILL_SmallIconURI"
	KeySkillLargeIconURI        string = "SKILL_LargeIconURI"
	KeySkillTestingInstructions string = "SKILL_TestingInstructions"
	KeySkillInvocation          string = "SKILL_Invocation"
	KeySkillPrivacyPolicyURL    string = "SKILL_PrivacyPolicyURL"
	KeySkillTermsOfUse          string = "SKILL_TermsOfUse"
	KeyPostfixSamples           string = "_Samples"
	KeyPostfixValues            string = "_Values"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// LocaleRegistry is the interface for an l10n registry.
type LocaleRegistry interface {
	Register(locale LocaleInstance, opts ...RegisterFunc) error
	Resolve(locale string) (LocaleInstance, error)
	GetDefault() LocaleInstance
	SetDefault(locale string) error
	GetLocales() map[string]LocaleInstance
}

// LocaleInstance is the interface for a specific locale.
type LocaleInstance interface {
	GetName() string
	Set(key string, values []string)
	Get(key string, args ...interface{}) string
	GetAny(key string, args ...interface{}) string
	GetAll(key string, args ...interface{}) []string
}

// Config contains the options for Locale registration
type Config struct {
	DefaultLocale bool
	FallbackFor   string
}

// TODO: move to `ssml` package
func Speak(text string) string {
	return "<speak>" + text + "</speak>"
}
func UseVoice(voice string, text string) string {
	return `<voice name="` + voice + `">` + text + `</voice>`
}

func UseVoiceLang(voice string, language string, text string) string {
	return `<voice name="` + voice + `"><lang xml:lang="` + language + `">` + text + `</lang></voice>`
}

// DefaultRegistry is the standard registry used
var DefaultRegistry = NewRegistry()

// Registry is the Locale registry
type Registry struct {
	defaultLocale string
	locales       map[string]LocaleInstance
}

func NewRegistry() LocaleRegistry {
	return &Registry{locales: map[string]LocaleInstance{}}
}

// RegisterFunc defines the functions to be passed to Register
type RegisterFunc func(cfg *Config)

// AsDefault sets the given Locale the default
func AsDefault() RegisterFunc {
	return func(cfg *Config) {
		cfg.DefaultLocale = true
	}
}

// Register registers a new Locale in the DefaultRegistry
func Register(locale LocaleInstance, opts ...RegisterFunc) error {
	return DefaultRegistry.Register(locale, opts...)
}

// Resolve returns the matching Locale from the DefaultRegistry
func Resolve(name string) (LocaleInstance, error) {
	return DefaultRegistry.Resolve(name)
}

// Register registers a new locale and fails if it already exists
func (r *Registry) Register(l LocaleInstance, opts ...RegisterFunc) error {
	_, ok := r.locales[l.GetName()]
	if ok {
		return fmt.Errorf("locale %s already registered", l.GetName())
	}

	// run all RegisterFuncs
	var cfg Config
	for _, opt := range opts {
		opt(&cfg)
	}

	// set locale as default
	if cfg.DefaultLocale || r.defaultLocale == "" {
		r.defaultLocale = l.GetName()
	}

	r.locales[l.GetName()] = l

	return nil
}

func (r *Registry) GetDefault() LocaleInstance {
	return r.locales[r.defaultLocale]
}

func (r *Registry) SetDefault(locale string) error {
	_, ok := r.locales[locale]
	if !ok {
		return fmt.Errorf("Locale '%s' not registered, cannot make it the default!", locale)
	}
	r.defaultLocale = locale
	return nil
}

func (r *Registry) GetLocales() map[string]LocaleInstance {
	return r.locales
}

// Resolve returns the Locale matching the given name or an error
func (r *Registry) Resolve(locale string) (LocaleInstance, error) {
	l, ok := r.locales[locale]
	if !ok {
		return nil, fmt.Errorf("locale %s not found", locale)
	}
	return l, nil
}

///////////////////////////////////////////

// Locale is a representation of keys in a specific language (and can have a fallback Locale)
type Locale struct {
	Name         string // de-DE, en-US, ...
	TextSnippets Snippets
	foo          Snippets
}

func NewLocale(locale string) *Locale {
	return &Locale{
		Name:         locale,
		TextSnippets: Snippets{},
	}
}

func (l *Locale) GetName() string {
	return l.Name
}

func (l *Locale) Set(key string, values []string) {
	l.TextSnippets[key] = values
}

func (l *Locale) Get(key string, args ...interface{}) string {
	return l.TextSnippets.GetFirst(key, args...)
}

func (l *Locale) GetAny(key string, args ...interface{}) string {
	t, _ := l.TextSnippets.GetAny(key, args...)
	return t
}

func (l *Locale) GetAll(key string, args ...interface{}) []string {
	t, _ := l.TextSnippets.GetAll(key, args...)
	return t
}

////////////////////////////////////

// Snippets is the actual representation of key -> array of texts in locale
type Snippets map[string][]string

func (s Snippets) GetFirst(key string, args ...interface{}) string {
	_, ok := s[key]
	if !ok || len(s[key]) == 0 {
		return ""
	}
	return fmt.Sprintf(s[key][0], args...)
}

// Get returns the translation for the snippet
func (s Snippets) GetAny(key string, args ...interface{}) (string, error) {
	_, ok := s[key]
	if !ok {
		return "", fmt.Errorf("key not defined %s", key)
	}
	if len(s[key]) == 0 {
		return "", fmt.Errorf("key not defined %s", key)
	}
	if len(s[key]) == 1 {
		return fmt.Sprintf(s[key][0], args...), nil
	}
	l := len(s[key])
	r := rand.Intn(l)
	return fmt.Sprintf(s[key][r], args...), nil
}

func (s Snippets) GetAll(k string, args ...interface{}) ([]string, error) {
	if len(s[k]) == 0 {
		return []string{}, fmt.Errorf("key not defined %s", k)
	}
	r := []string{}
	for _, v := range s[k] {
		r = append(r, fmt.Sprintf(v, args...))
	}
	return r, nil
}
