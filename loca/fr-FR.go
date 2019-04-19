package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// just register the locale and fallback to enUS
var frFR = &l10n.Locale{
	Name:         "fr-FR",
	TextSnippets: map[string][]string{},
}
