package alexa

const (
	LocaleGermanGerman      = "de-DE"
	LocaleAustralianEnglish = "en-AU"
	LocaleCanadianEnglish   = "en-CA"
	LocaleBritishEnglish    = "en-GB"
	LocaleIndianEnglish     = "en-IN"
	LocaleAmericanEnglish   = "en-US"
	LocaleSpanishSpanish    = "es-ES"
	LocaleMexicanSpanish    = "es-MX"
	LocaleCanadianFrench    = "fr-CA"
	LocaleFrenchFrench      = "fr-FR"
	LocaleItalian           = "it-IT"
	LocaleJapanese          = "ja-JP"
)

// IsEnglish returns true if locale is English
func IsEnglish(locale string) bool {
	switch locale {
	case LocaleAmericanEnglish, LocaleIndianEnglish, LocaleBritishEnglish, LocaleCanadianEnglish, LocaleAustralianEnglish:
		return true
	default:
		return false
	}
}
