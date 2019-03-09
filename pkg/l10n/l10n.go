package l10n

import (
	"fmt"
	"math/rand"
)

type Key string

type Snippets map[Key][]string

// Get returns the translation for the snippet
func (s Snippets) Get(k Key, args ...interface{}) string {
	if len(s[k]) < 1 {
		return string(k)
	}
	l := len(s[k])
	r := rand.Intn(l)
	return fmt.Sprintf(s[k][r], args...)
}

type Locale struct {
	Name         string
	TextSnippets Snippets
}

// GetText returns the translation for the selected language
func (l Locale) GetSnippet(k Key, args ...interface{}) string {
	return l.TextSnippets.Get(k, args...)
}

// locales contains the registered locales
var locales map[string]*Locale

// Register the given locale
func Register(l *Locale) {
	locales[l.Name] = l
}

func Resolve(name string) (*Locale, error) {
	l, ok := locales[name]
	if !ok {
		return nil, fmt.Errorf("locale %s not found", name)
	}
	return l, nil
}
