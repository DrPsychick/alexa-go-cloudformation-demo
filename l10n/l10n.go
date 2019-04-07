package l10n

import "github.com/DrPsychick/alexa-go-cloudformation-demo/pkg/l10n"

const (
	GreetingTitle     l10n.Key = "greeting_title"
	Greeting          l10n.Key = "greeting"
	GreetingSSML      l10n.Key = "greeting_ssml"
	ByeBye            l10n.Key = "byebye"
	StopTitle         l10n.Key = "stop_title"
	Stop              l10n.Key = "stop"
	SaySomethingTitle l10n.Key = "saysomething_title"
	SaySomething      l10n.Key = "saysomething"
)

func init() {
	for _, l := range []l10n.Locale{deDE, enUS, frFR} {
		if err := l10n.Register(&l); err != nil {
			panic("registration of locale failed")
		}
	}
}
