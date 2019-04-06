package l10n

import (
	"fmt"
	"math/rand"
)

// DefaultRegistry is the standard registry used
var DefaultRegistry = &Registry{
	DefaultLocale: "",
	//fallbacks: map[string]*Locale{},
	locales: map[string]*Locale{},
}

type Locale struct {
	Name         string  // de-DE, en-US, ...
	Fallback     *Locale // points to fallback (or nil)
	TextSnippets Snippets
}

type Key string

type Snippets map[Key][]string

type Registry struct {
	DefaultLocale string
	fallbacks     map[string]*Locale
	locales       map[string]*Locale
}

type RegisterFunc func(cfg *Config)

type Config struct {
	DefaultLocale bool
	FallbackFor   string
}

func AsDefault() RegisterFunc {
	return func(cfg *Config) {
		cfg.DefaultLocale = true
	}
}

func AsFallbackFor(name string) RegisterFunc {
	return func(cfg *Config) {
		cfg.FallbackFor = name
	}
}

func Register(l *Locale, opts ...RegisterFunc) error {
	return DefaultRegistry.Register(l, opts...)
}

func Resolve(n string) (*Locale, error) {
	return DefaultRegistry.Resolve(n)
}

// Register registers a new locale, fails if it already exists
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
	fmt.Printf("%s: Fallback -%s-\n", l.Name, cfg.FallbackFor)
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
		fmt.Printf("Setting fallback for %s to %s\n", orig.Name, l.Name)

		// set the fallback
		r.locales[cfg.FallbackFor].Fallback = l
	}

	// set locale as default
	if cfg.DefaultLocale {
		r.DefaultLocale = l.Name
	}

	r.locales[l.Name] = l

	return nil
}

func (r *Registry) Resolve(name string) (*Locale, error) {
	l, ok := r.locales[name]
	if !ok {
		return nil, fmt.Errorf("locale %s not found", name)
	}
	return l, nil
}

////////////////

// Get returns the translation for the snippet
func (s Snippets) Get(k Key, args ...interface{}) (string, error) {
	if len(s[k]) < 1 {
		return "", fmt.Errorf("key not defined %s", string(k))
	}
	l := len(s[k])
	r := rand.Intn(l)
	//fmt.Printf("length: %d rand: %d", l, r)
	return fmt.Sprintf(s[k][r], args...), nil
}

// GetText returns the translation for the selected language
func (l Locale) GetSnippet(k Key, args ...interface{}) string {
	r, err := l.TextSnippets.Get(k, args...)
	if err == nil {
		return r
	}
	if l.Fallback != nil {
		fmt.Printf("Using fallback: %s\n", l.Fallback.Name)
		r, err = l.Fallback.TextSnippets.Get(k, args...)
		if err == nil {
			return r
		}
	}

	return string(k)
}

//// Register the given locale
//func Register(l *Locale) {
//	locales[l.Name] = l
//}
