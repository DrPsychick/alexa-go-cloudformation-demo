package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// l10n.Keys of the project
const (
	GreetingTitle l10n.Key = "greeting_title"
	Greeting      l10n.Key = "greeting"
	GreetingSSML  l10n.Key = "greeting_ssml"
	ByeBye        l10n.Key = "byebye"
	StopTitle     l10n.Key = "stop_title"
	Stop          l10n.Key = "stop"

	// Intents
	SaySomething l10n.Key = "SaySomething"
	DemoIntent   l10n.Key = "DemoIntent"

	// Types
	TypeBeerCountries        l10n.Key = "BEER_Countries"
	TypeBeerCountriesValues  l10n.Key = "BEER_CountriesValues"
	TypePeopleCategory       l10n.Key = "BEER_PeopleCategory"
	TypePeopleCategoryValues l10n.Key = "BEER_PeopleCatgoryValues"
)

func init() {
	var locales = []*l10n.Locale{
		deDE, enUS, frFR,
	}
	for _, l := range locales {
		if err := l10n.Register(l); err != nil {
			panic("registration of locale failed")
		}
	}
}
