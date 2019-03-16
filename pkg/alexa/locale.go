package alexa

type Locale string

// locales
const (
	// LocaleItalian is the locale for Italian
	LocaleItalian Locale = "it-IT"

	// LocaleGerman is the locale for standard dialect German
	LocaleGerman Locale = "de-DE"

	// LocaleAustralianEnglish is the locale for Australian English
	LocaleAustralianEnglish Locale = "en-AU"

	//LocaleCanadianEnglish is the locale for Canadian English
	LocaleCanadianEnglish Locale = "en-CA"

	//LocaleBritishEnglish is the locale for UK English
	LocaleBritishEnglish Locale = "en-GB"

	//LocaleIndianEnglish is the locale for Indian English
	LocaleIndianEnglish Locale = "en-IN"

	//LocaleAmericanEnglish is the locale for American English
	LocaleAmericanEnglish Locale = "en-US"

	// LocaleJapanese is the locale for Japanese
	LocaleJapanese Locale = "ja-JP"
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
