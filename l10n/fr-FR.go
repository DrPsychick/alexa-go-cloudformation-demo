package l10n

import (
	"github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

// just register the locale and fallback to enUS
var frFR = l10n.Locale{
	Name:         "fr-FR",
	Fallback:     &enUS,
	TextSnippets: map[l10n.Key][]string{},
}
