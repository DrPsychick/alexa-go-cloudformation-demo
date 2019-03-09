package alexa

// locales
const (
	// LocaleItalian is the locale for Italian
	LocaleItalian = "it-IT"

	// LocaleGerman is the locale for standard dialect German
	LocaleGerman = "de-DE"

	// LocaleAustralianEnglish is the locale for Australian English
	LocaleAustralianEnglish = "en-AU"

	//LocaleCanadianEnglish is the locale for Canadian English
	LocaleCanadianEnglish = "en-CA"

	//LocaleBritishEnglish is the locale for UK English
	LocaleBritishEnglish = "en-GB"

	//LocaleIndianEnglish is the locale for Indian English
	LocaleIndianEnglish = "en-IN"

	//LocaleAmericanEnglish is the locale for American English
	LocaleAmericanEnglish = "en-US"

	// LocaleJapanese is the locale for Japanese
	LocaleJapanese = "ja-JP"
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
