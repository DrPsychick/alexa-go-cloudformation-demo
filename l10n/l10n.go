package l10n

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/l10n"
)

const (
	GreetingTitle     l10n.Key = "greeting_title"
	Greeting          l10n.Key = "greeting"
	GreetingSSML      l10n.Key = "greeting_ssml"
	ByeBye            l10n.Key = "byebye"
	StopTitle         l10n.Key = "stop_title"
	Stop              l10n.Key = "stop"
	SaySomethingTitle l10n.Key = "saysomething_title"
	SaySomething      l10n.Key = "saysomething"
	SaySomethingSSML  l10n.Key = "saysomething_ssml"
)

func init() {
	// ERROR: since l gets overwritten, locale "switches" to the last one
	//var locales = []l10n.Locale{
	//	deDE, enUS, frFR,
	//}
	//for _, l := range locales {
	//	fmt.Printf("Registering locale %s...\n", l.Name)
	//	if err := l10n.Register(&l); err != nil {
	//		panic("registration of locale failed")
	//	}
	//}
	//l, _ := l10n.Resolve("en-US")
	//fmt.Printf("German: %s\n", l.Name)
	if err := l10n.Register(&deDE); err != nil {
		panic("registration of locale failed")
	}
	if err := l10n.Register(&enUS); err != nil {
		panic("registration of locale failed")
	}
	if err := l10n.Register(&frFR); err != nil {
		panic("registration of locale failed")
	}
}
