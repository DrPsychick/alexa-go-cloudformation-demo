package alexa

// Locale defines the type of the Locale* consts
type Locale string

// locales
const (
	// LocaleAmericanEnglish is the locale for American English
	LocaleAmericanEnglish Locale = "en-US"
	// LocaleAustralianEnglish is the locale for Australian English
	LocaleAustralianEnglish Locale = "en-AU"
	// LocaleBritishEnglish is the locale for UK English
	LocaleBritishEnglish Locale = "en-GB"
	// LocaleCanadianEnglish is the locale for Canadian English
	LocaleCanadianEnglish Locale = "en-CA"
	// LocaleCanadianFrench is the locale for Canadian French
	LocaleCanadianFrench Locale = "fr-CA"
	// LocaleFrenchFrench is the locale for French (France)
	LocaleFrench Locale = "fr-FR"
	// LocaleGerman is the locale for standard dialect German (Germany)
	LocaleGerman Locale = "de-DE"
	//LocaleIndianEnglish is the locale for Indian English
	LocaleIndianEnglish Locale = "en-IN"
	// LocaleItalian is the locale for Italian (Italy)
	LocaleItalian Locale = "it-IT"
	// LocaleJapanese is the locale for Japanese (Japan)
	LocaleJapanese Locale = "ja-JP"
	// LocaleMexicanSpanish is the locale for Mexican Spanish
	LocaleMexicanSpanish Locale = "es-MX"
	// LocaleSpanish is the Locale for Spanish (Spain)
	LocaleSpanish Locale = "es-ES"
)

// IsEnglish returns true if locale is English
func IsEnglish(locale Locale) bool {
	switch locale {
	case LocaleAmericanEnglish, LocaleIndianEnglish, LocaleBritishEnglish, LocaleCanadianEnglish, LocaleAustralianEnglish:
		return true
	default:
		return false
	}
}
