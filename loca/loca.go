package loca

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

const (
	GreetingTitle l10n.Key = "greeting_title"
	Greeting      l10n.Key = "greeting"
	GreetingSSML  l10n.Key = "greeting_ssml"
	ByeBye        l10n.Key = "byebye"
	StopTitle     l10n.Key = "stop_title"
	Stop          l10n.Key = "stop"
	// SaySomething Intent
	SaySomethingTitle  l10n.Key = "saysomething_title"
	SaySomething       l10n.Key = "saysomething"
	SaySomethingSSML   l10n.Key = "saysomething_ssml"
	SaySomethingIntent l10n.Key = "SaySomethingIntent"
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
