package l10n

import (
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
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
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
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
var DefaultRegistry = &Registry{
	locales: map[string]*Locale{},
}

// Registry is the Locale registry
type Registry struct {
	defaultLocale string
	locales       map[string]*Locale
}

func NewRegistry() *Registry {
	return &Registry{locales: map[string]*Locale{}}
}

// RegisterFunc defines the functions to be passed to Register
type RegisterFunc func(cfg *Config)

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
		fb := r.locales[cfg.FallbackFor]
		fb.Fallback = l
	}

	// set locale as default
	if cfg.DefaultLocale {
		r.defaultLocale = l.Name
	}

	r.locales[l.Name] = l

	return nil
}

func (r *Registry) GetDefault() *Locale {
	return r.locales[r.defaultLocale]
}
func (r *Registry) GetLocales() map[string]*Locale {
	return r.locales
}

// Resolve returns the Locale matching the given name or an error
func (r *Registry) Resolve(locale string) (*Locale, error) {
	l, ok := r.locales[locale]
	if !ok {
		return nil, fmt.Errorf("locale %s not found", locale)
	}
	return l, nil
}

// Snippets is the actual representation of key -> array of texts in locale
type Snippets map[string][]string

func (s Snippets) GetFirst(key string) string {
	return s[key][0]
}

// Get returns the translation for the snippet
func (s Snippets) Get(k string, args ...interface{}) (string, error) {
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

func (s Snippets) GetAll(k string, args ...interface{}) ([]string, error) {
	if len(s[k]) < 1 {
		return []string{}, fmt.Errorf("key not defined %s", string(k))
	}
	return s[k], nil
}

// Locale is a representation of keys in a specific language (and can have a fallback Locale)
type Locale struct {
	Name            string          // de-DE, en-US, ...
	Countries       []alexa.Country // countries associated with this locale
	Invocation      string          // "my skill"
	Fallback        *Locale         // points to fallback (or nil)
	TextSnippets    Snippets
	IntentResponses IntentResponses
}

func (l Locale) GetName() string {
	return l.Name
}

func (l Locale) Get(key string) string {
	return l.TextSnippets.GetFirst(key)
}

func (l Locale) GetAny(key string) string {
	t, _ := l.TextSnippets.Get(key)
	return t
}

func (l Locale) GetAll(key string) []string {
	t, _ := l.TextSnippets.GetAll(key)
	return t
}

// GetSnippet returns the translation and follows fallback chain
func (l Locale) GetSnippet(k string, args ...interface{}) string {
	if r, err := l.TextSnippets.Get(k, args...); err == nil {
		return r
	}
	if l.Fallback != nil {
		return l.Fallback.GetSnippet(k, args...)
	}

	return string(k)
}

// GetAllSnippets returns all available translations and follows fallback chain
func (l Locale) GetAllSnippets(k string, args ...interface{}) []string {
	if r, err := l.TextSnippets.GetAll(k); err == nil {
		return r
	}
	if l.Fallback != nil {
		return l.Fallback.GetAllSnippets(k, args...)
	}

	return []string{string(k)}
}

// GetIntent returns complete localized intent
func (l Locale) GetIntent(key string) (IntentResponse, error) {
	if r, ok := l.IntentResponses[key]; ok {
		return r, nil
	}
	return IntentResponse{}, fmt.Errorf("No %s translations for intent %s", l.Name, key)
}

// TODO: refactor to return a response object usable by ResponseBuilder?
func (l Locale) GetSingleIntentResponse(key string) (string, string, string) {
	ir, _ := l.GetIntent(key)

	if len(ir.Text) > 0 {
		r := rand.Intn(len(ir.Text))
		return ir.Title[0], ir.Text[r], ir.SSML[r]
	}

	return string(key + ".Title"), string(key + ".Text"), string(key + ".SSML")

}

// Responses is the representation of a list of IntentResponses
type IntentResponses map[string]IntentResponse

type IntentResponse struct {
	Samples []string
	Title   []string
	Text    []string
	SSML    []string
	Slots   map[string]Slot
}

type Slot struct {
	Samples             []string
	PromptElicitations  []alexa.PromptVariations
	PromptConfirmations []alexa.PromptVariations // TODO: is this correct?
}

type Prompts map[string][]alexa.PromptVariations
