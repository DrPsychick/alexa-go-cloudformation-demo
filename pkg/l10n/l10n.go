package l10n

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"math/rand"
	"time"
)

// Default keys
const (
	KeySkillName                Key = "SKILL_Name"
	KeySkillDescription         Key = "SKILL_Description"
	KeySkillSummary             Key = "SKILL_Summary"
	KeySkillExamplePhrases      Key = "SKILL_ExamplePhrases"
	KeySkillKeywords            Key = "SKILL_Keywords"
	KeySkillSmallIconURI        Key = "SKILL_SmallIconURI"
	KeySkillLargeIconURI        Key = "SKILL_LargeIconURI"
	KeySkillTestingInstructions Key = "SKILL_TestingInstructions"
	KeySkillInvocation          Key = "SKILL_Invocation"
	KeySkillPrivacyPolicyURL    Key = "SKILL_PrivacyPolicyURL"
	KeySkillTermsOfUse          Key = "SKILL_TermsOfUse"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Registry is the Locale registry
type Registry struct {
	defaultLocale string
	locales       map[string]*Locale
}

// DefaultRegistry is the standard registry used
var DefaultRegistry = &Registry{
	locales: map[string]*Locale{},
}

// Locale is a representation of Keys in a specific language (and can have a fallback Locale)
type Locale struct {
	Name            string  // de-DE, en-US, ...
	Invocation      string  // "my skill"
	Fallback        *Locale // points to fallback (or nil)
	TextSnippets    Snippets
	IntentResponses IntentResponses
	Countries       []alexa.Country // countries associated with this locale
	// Utterances
}

// Key defines the type of a text key
type Key string

// Snippets is the actual representation of key -> array of texts in locale
type Snippets map[Key][]string

// Responses is the representation of a list of IntentResponses
type IntentResponses map[Key]IntentResponse

type IntentResponse struct {
	Samples []string
	Title   []string
	Text    []string
	SSML    []string
}

// RegisterFunc defines the functions to be passed to Register
type RegisterFunc func(cfg *Config)

// Config contains the options for Locale registration
type Config struct {
	DefaultLocale bool
	FallbackFor   string
}

// AsDefault sets the given Locale the default
func AsDefault() RegisterFunc {
	return func(cfg *Config) {
		cfg.DefaultLocale = true
	}
}

// TODO Obsolete:
// AsFallbackFor registers the locale as fallback Locale for the given Locale name
func AsFallbackFor(name string) RegisterFunc {
	return func(cfg *Config) {
		cfg.FallbackFor = name
	}
}

// Register registers a new Locale in the DefaultRegistry
func Register(locale *Locale, opts ...RegisterFunc) error {
	return DefaultRegistry.Register(locale, opts...)
}

// Resolve returns the matching Locale from the DefaultRegistry
func Resolve(name string) (*Locale, error) {
	return DefaultRegistry.Resolve(name)
}

// Register registers a new locale and fails if it already exists
func (r *Registry) Register(l *Locale, opts ...RegisterFunc) error {
	_, ok := r.locales[l.Name]
	if ok {
		return fmt.Errorf("locale %s already registered", l.Name)
	}

	// run all RegisterFuncs
	var cfg Config
	for _, opt := range opts {
		opt(&cfg)
	}

	// order matters, first anything that can fail, then changing data
	if cfg.FallbackFor != "" {
		// fallback locale must be registered
		orig, ok := r.locales[cfg.FallbackFor]
		if !ok {
			return fmt.Errorf("cannot be fallback for locale %s as it is not registered", cfg.FallbackFor)
		}

		// fallback already defined
		if orig.Fallback != nil {
			return fmt.Errorf("fallback for locale %s is already registered", orig.Name)
		}

		// set the fallback
		r.locales[cfg.FallbackFor].Fallback = l
	}

	// set locale as default
	if cfg.DefaultLocale {
		r.defaultLocale = l.Name
	}

	r.locales[l.Name] = l

	return nil
}

func (r *Registry) GetDefaultLocale() string {
	return r.defaultLocale
}
func (r *Registry) GetLocales() map[string]*Locale {
	return r.locales
}

// Resolve returns the Locale matching the given name or an error
func (r Registry) Resolve(name string) (*Locale, error) {
	l, ok := r.locales[name]
	if !ok {
		return nil, fmt.Errorf("locale %s not found", name)
	}
	return l, nil
}

// Get returns the translation for the snippet
func (s Snippets) Get(k Key, args ...interface{}) (string, error) {
	if len(s[k]) < 1 {
		return "", fmt.Errorf("key not defined %s", string(k))
	}
	if len(s[k]) == 1 {
		return fmt.Sprintf(s[k][0], args...), nil
	}
	l := len(s[k])
	r := rand.Intn(l)
	return fmt.Sprintf(s[k][r], args...), nil
}

func (s Snippets) GetAll(k Key, args ...interface{}) ([]string, error) {
	if len(s[k]) < 1 {
		return []string{}, fmt.Errorf("key not defined %s", string(k))
	}
	return s[k], nil
}

// GetSnippet returns the translation and follows fallback chain
func (l Locale) GetSnippet(k Key, args ...interface{}) string {
	if r, err := l.TextSnippets.Get(k, args...); err == nil {
		return r
	}
	if l.Fallback != nil {
		return l.Fallback.GetSnippet(k, args...)
	}

	return string(k)
}

// GetAllSnippets returns all available translations and follows fallback chain
func (l Locale) GetAllSnippets(k Key, args ...interface{}) []string {
	if r, err := l.TextSnippets.GetAll(k); err == nil {
		return r
	}
	if l.Fallback != nil {
		return l.Fallback.GetAllSnippets(k, args...)
	}

	return []string{string(k)}
}

func (l Locale) GetIntent(key Key) IntentResponse {
	return l.IntentResponses[key]
}

func (l Locale) GetSingleIntentResponse(key Key) (string, string, string) {
	r := rand.Intn(len(l.IntentResponses[key].Text))
	return l.IntentResponses[key].Title[0], l.IntentResponses[key].Text[r], l.IntentResponses[key].SSML[r]
}
