# how to use `alexa` package to build skill and models

## Simple case, Skill with only one language
Building the skill
```go
package demo
import (
    "github.com/drpsychick/go-alexa-lambda"
    "github.com/drpsychick/go-alexa-lambda/gen"
    "encoding/json"
    "fmt"
)

func demo() {
	var sb *gen.SkillBuilder
    sb = gen.NewSkillBuilder().
        WithCategory(alexa.CategoryCommunication).
        AddCountry("US")
    
    sb.AddLocale("en-US").Locale("en-US").
        WithLocaleName("my name").
        WithLocaleDescription("my description").
        WithLocaleSummary("my summary").
        WithLocaleKeywords([]string{"word1", "word2"}).
        WithLocaleExamples([]string{"make an example", "give an example"}).
        WithLocaleSmallIcon("https://small.icon").
        WithLocaleLargeIcon("https://large.icon").
        WithLocalePrivacyURL("https://privacy.url/en-US/")

    // must be used *after* adding a locale
    sb.WithDefaultLocaleTestingInstructions("Foo bar")
    
    sk, err := sb.Build()
    if err != nil { return }

    res, err := json.MarshalIndent(sk, "", "  ")
    if err != nil { return }
    fmt.Printf("%s\n", string(res))
}
```

Building the model
```go
package demo
import (
    "github.com/drpsychick/go-alexa-lambda"
    "github.com/drpsychick/go-alexa-lambda/l10n"
    "github.com/drpsychick/go-alexa-lambda/gen"
)
func demo() {
    // you can still use l10n.LocaleRegistry to resolve translations if you wish
    registry := l10n.NewRegistry()
    err := registry.Register(l10n.NewLocale("en-US"))
    if err != nil { return }
    
    loc, err := registry.Resolve("en-US")
    loc.Set("MyIntent_Samples", []string{"sample one", "sample two"})
    
    // or: gen.NewModelBuilder()
    sb := gen.NewSkillBuilder()
    mb := sb.WithModel().Model().
        WithDelegationStrategy(alexa.DelegationSkillResponse).
        WithLocale("en-US", "my skill").
        WithLocale("de-DE", "mein skill")
    
    mb.WithType("TypeSlotOne").Type("TypeSlotOne").
        WithLocaleValues("en-US", []string{"One"}).
        WithLocaleValues("de-DE", []string{"Eins"})
    
    mb.WithIntent("MyIntent").Intent("MyIntent").
        WithLocaleSamples(loc.GetName(), loc.GetAll("MyIntent_Samples")).
        WithLocaleSamples("de-DE", []string{"sample eins", "sample zwei"}).
        WithSlot("SlotName", "TypeSlotOne").Slot("SlotName").
            WithLocaleSamples(loc.GetName(), []string{"of {Slot}"}).
            WithLocaleSamples("de-DE", []string{"von {Slot}"})
    
    mb.WithElicitationSlotPrompt("MyIntent", "SlotName").
        ElicitationPrompt("MyIntent", "SlotName").
            WithVariation("PlainText").Variation("PlainText").
                WithLocaleTypeValue("de-DE", "PlainText", []string{"Was?", "Wie bitte?"}).
                WithLocaleTypeValue(loc.GetName(), "PlainText", []string{"What?"})
        
    mb.WithConfirmationSlotPrompt("MyIntent", "SlotName").
    	ConfirmationPrompt("MyIntent", "SlotName").
            WithVariation("PlainText").Variation("PlainText").
                WithLocaleTypeValue(loc.GetName(), "PlainText", []string{"Sure?"}).
                WithLocaleTypeValue("de-DE", "PlainText", []string{"Sicher?"})
    
    // *alexa.Model
    m, err := mb.BuildLocale("en_US")
    
    // map[locale]*alexa.Model of models
    ms, err := mb.Build()
}
```

## International case, multiple languages
Definining locales
```go
package demo
import (
    "github.com/drpsychick/go-alexa-lambda/l10n"	
)
var enUS = &l10n.Locale{
    Name: "en-US",
    TextSnippets: map[string][]string{
        // Skill
        l10n.KeySkillName:                []string{"SkillName"},
        l10n.KeySkillDescription:         []string{"SkillDescription"},
        l10n.KeySkillSummary:             []string{"SkillSummary"},
        l10n.KeySkillKeywords:            []string{"Keyword1", "Keyword2"},
        l10n.KeySkillExamplePhrases:      []string{"start me", "boot me up"},
        l10n.KeySkillSmallIconURI:        []string{"https://small"},
        l10n.KeySkillLargeIconURI:        []string{"https://large"},
        l10n.KeySkillPrivacyPolicyURL:    []string{"https://policy"},
        l10n.KeySkillTermsOfUseURL:       []string{"https://toc"},
        l10n.KeySkillInvocation:          []string{"call me"},
        l10n.KeySkillTestingInstructions: []string{"Initial instructions"},
        // My Intent
        "MyIntent_Samples":                []string{"say one", "say two"},
        "MyIntent_Title":                  []string{"Title"},
        "MyIntent_Text":                   []string{"Text1", "Text2"},
        "MyIntent_SSML":                   []string{l10n.Speak("SSML one"), l10n.Speak("SSML two")},
        // Slot Intent
        "SlotIntent_Samples":              []string{"what about slot {SlotName}"},
        "SlotIntent_Title":                []string{"Test intent with slot"},
        "SlotIntent_Text":                 []string{"it seems to work"},
        "SlotIntent_SlotName_Samples":     []string{"of {SlotName}", "{SlotName}"},
        "SlotIntent_SlotName_Elicit_Text": []string{"Which slot did you mean?", "I did not understand, which slot?"},
        "SlotIntent_SlotName_Elicit_SSML": []string{l10n.Speak("I'm sorry, which slot did you mean?")},
        // Types
        "MyType_Values":                   []string{"Value 1", "Value 2"},
    },
}
```
Building the skill
```go
package demo
import (
    "github.com/drpsychick/go-alexa-lambda"	
    "github.com/drpsychick/go-alexa-lambda/l10n"	
    "github.com/drpsychick/go-alexa-lambda/gen"	
)
var enUS = &l10n.Locale{}
func demo() {
    // register the locales... first one automatically is default
    registry := l10n.NewRegistry()
    // there are multiple ways to set the default explicitly
    registry.Register(enUS, l10n.AsDefault())
    registry.SetDefault("en-US")
    
    // pass the registry
    sb := gen.NewSkillBuilder().
        WithLocaleRegistry(registry).
        WithCategory(alexa.CategoryFashionAndStyle)

    // *alexa.Skill
    s, err := sb.Build()
}   
```
Building the models
```go
package demo
import (
    "github.com/drpsychick/go-alexa-lambda"	
    "github.com/drpsychick/go-alexa-lambda/l10n"	
    "github.com/drpsychick/go-alexa-lambda/gen"	
)
var registry = &l10n.Registry{}
func demo() {
    // pass the registry
    mb := gen.NewModelBuilder().
        WithLocaleRegistry(registry).
        WithDelegationStrategy(alexa.DelegationSkillResponse)
    // add intents, types, slots, prompts, ...
    mb.WithType("MyType") // looks up "MyType_Values
    mb.WithIntent("MyIntent") // looks up "MyIntent_Samples"
    mb.Intent("SlotIntent"). // looks up "SlotIntent_Samples"
        WithSlot("SlotName", "MyType") // looks up "SlotIntent_SlotName_Samples"
    
    mb.WithElicitationSlotPrompt("SlotIntent", "SlotName")
    mb.ElicitationPrompt("SlotIntent", "SlotName").
        WithVariation("PlainText").
        WithVariation("SSML")
    
    ms, err := mb.Build()
}
```

## Expert case (build your own)
Simply build your own JSON
```go
package demo
import ("github.com/drpsychick/go-alexa-lambda")
var skill = &alexa.Skill{}
var modelEnUs = &alexa.Model{}
```
As you're an expert, you can easily figure out how to do that in detail by looking at the tests: 
https://github.com/DrPsychick/alexa-go-cloudformation-demo/blob/master/pkg/alexa/skill_test.go

# how to use `alexa` package to build lambda

Application: defines the flow
```go
package alfalfa

import (
    "github.com/drpsychick/alexa-go-cloudformation-demo/loca"
    "github.com/drpsychick/go-alexa-lambda/l10n"
    "github.com/hamba/pkg/log"
    "github.com/hamba/pkg/stats"
    "log"
    "fmt"
)

// The Applications reponsibility is to execute different functions for the skill,
// but it does not know about intents/the skill, it could be used as well without alexa.
// Our application purpose will be to return simple responses, playing with Alexas voices and languages.
type Application struct {
    logger  log.Logger
    statter stats.Statter
}
func NewApplication(l log.Logger, s stats.Statter) *Application {
    return &Application{
        logger:  l,
        statter: s,
    }
}

// ApplicationResponse 
type ApplicationResponse struct {
    Title  string
    Text   string
    Speech string
    Image  string
    End    bool
}
func NewApplicationResponse() *ApplicationResponse {
	return &ApplicationResponse{}
}

type Config struct {
	context string
	user    string
	session string
	time    string
}

type ResponseFunc func(cfg *Config)

// WithContext|User|Session|Time|...
func WithUser(user string) ResponseFunc {
	return func(cfg *Config) {
		cfg.user = user
	}
}

type AppResponse func(locale l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error)

// SaySomething responds with a random sentence.
// should this use translations? yes!
// the application binds function with l10n key and handles errors/default
func (a *Application) SaySomething() AppResponse {
	return AppResponse(func(loc l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error) {
        // run all ResponseFuncs
        cfg := &Config{}
        for _, opt := range opts {
            opt(cfg)
        }
        
        tit := ""
        msg := ""
        msgSSML := ""
        if cfg.user != "" {
        	// personalized response
        	tit = loc.GetAny("SaySomething_UserTitle", cfg.user)
        	msg = loc.GetAny("SaySomething_UserResponse", cfg.user)
            msgSSML = loc.GetAny("SaySomething_UserResponse"+l10n.KeyPostfixSSML, cfg.user)        	
        } else {
        	tit = loc.GetAny("SaySomething_Title")
            msg = loc.GetAny("SaySomething_Response")
            msgSSML = loc.GetAny("SaySomething_Response"+l10n.KeyPostfixSSML)
    	}
        
    	if msg == "" {
    		return ApplicationResponse{}, fmt.Errorf("Error_NoTranslation")
    	}
	
		return ApplicationResponse{
			Title: tit,
			Text: msg,
			Speech: msgSSML,
			End: true,
		}, nil
    })
}
// ShowSomething responds with an image and a sentence.
func (a *Application) ShowSomething() AppResponse {
    return AppResponse(func(loc l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error) {
	    // similar to above, but returns an image URL which lambda can render into a card.
	    return ApplicationResponse{
	    	Text: "text", 
	    	Speech: "<speak>text</speak", 
	    	Image: "https://myimage.url/img_%s.png",
	    	End: true,
	    }, nil
	})
}

func (a *Application) StopFunc() AppResponse {
    return AppResponse(func(locale l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error) {
        return ApplicationResponse{
            Title:  locale.GetAny(l10n.KeyStopTitle),
            Text:   locale.GetAny(l10n.KeyStopText),
            Speech: locale.GetAny(l10n.KeyStopSSML),
            End:    true,
        }, nil
    })
}

func (a *Application) ErrorFunc(err error) AppResponse {
	return AppResponse(func(locale l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error) {
		return ApplicationResponse{
			Title:  locale.GetAny(l10n.KeyErrorTitle),
			Text:   locale.GetAny(l10n.KeyErrorText, err),
			Speech: locale.GetAny(l10n.KeyErrorSSML),
			End:    true,
		}, nil
	})
}

func (a *Application) ConvertTimeFunc() AppResponse {
	return AppResponse(func(locale l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error) {
        // run all ResponseFuncs
        cfg := &Config{}
        for _, opt := range opts {
            opt(cfg)
        }
        
        if cfg.time == "" {
        	return ApplicationResponse{}, fmt.Errorf("Error_NoTimeProvided")
        }
        // TODO: implement time conversion
    
        // trigger actions
        result, err := call.MyAPI(a)
        if err != nil {
            return ApplicationResponse{}, fmt.Errorf("Error_APICallFailed")
        }
        img := redis.GetKey() // fetch images
        
        // define standard response
        r := ApplicationResponse{
            Title:  locale.GetAny(loca.GenericTitle),
            Text:   locale.GetAny(loca.ConvertTimeText),
        }
        
        // api call was fine
        if result == "fine" {
            r.Title = locale.GetAny(loca.ConvertTimeFineTitle)
            r.Text = locale.GetAny(loca.ConvertTimeFineText)
            if img != "" {
                r.Image = img
            }
            r.End = true
        }
        
        return r, nil
    })
}
```

Lambda: registers handlers and builds the response
```go
package lambda

import (
    "github.com/drpsychick/alexa-go-cloudformation-demo/loca"
    "github.com/drpsychick/go-alexa-lambda"
    "github.com/drpsychick/go-alexa-lambda/l10n"
    "github.com/drpsychick/go-alexa-lambda/gen"
    "strings"
    "fmt"
)

type Application interface {
	Stop()          AppResponse
	SaySomething()  AppResponse
	ShowSomething() AppResponse
	WithUser(user string)
}
// interface to the application response
type ApplicationResponse interface {
	Title()  string
	Text()   string
	Speech() string
	Image()  string
	End()    bool
}

type Config struct {
	context string
	user    string
	session string
	time    string
}

type ResponseFunc func(cfg *Config)

// WithContext|User|Session|Time|...
func WithUser(user string) ResponseFunc {
	return func(cfg *Config) {
		cfg.user = user
	}
}

type AppResponse func(locale l10n.LocaleInstance, opts... ResponseFunc) (ApplicationResponse, error)

// requires loca.Registry and a gen.SkillBuilder
func NewMux(app Application, sk gen.SkillBuilder) alexa.Handler {
	sk.WithModel()
		
    mux := alexa.NewServerMux()

    // special requests that always exist
    mux.HandleRequestTypeFunc(alexa.TypeLaunchRequest, handleLaunch(app))
    mux.HandleRequestTypeFunc(alexa.TypeCanFulfillIntentRequest, handleCanFulfillIntent)
    
    // needs: mux, intent name, app function, skill/model
    registerSimpleIntent(mux, alexa.StopIntent, app.Stop(), sk)
    registerSimpleIntent(mux, "SaySomething", app.SaySomething(), sk)
    registerStandardIntent(mux, "ShowSomething", app.ShowSomething(), sk)

    return mux
}

func registerSimpleIntent(mux *alexa.ServeMux, intent string, f AppResponse, sk gen.SkillBuilder) {
	// register intent
	mux.HandleIntent(intent, handleSimpleIntent(f))
	sk.Model().WithIntent(intent)
}

func registerStandardIntent(mux *alexa.ServeMux, intent string, f AppResponse, sk gen.SkillBuilder) {
	mux.HandleIntent(intent, handleStandardIntent(f))
	sk.Model().WithIntent(intent)
}

func handleSimpleIntent(f AppResponse) alexa.Handler {
	return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
        // resolve locale
        def := loca.Registry.GetDefault()
        if def == nil {
        	handleError(b, r, fmt.Errorf("no default locale registered"))
        	return
        }
        loc, err := loca.Registry.Resolve(r.Locale)
        if err != nil {
        	handleError(b, r, err)
        	return
        }
		
		// run func
		resp, err := f(loc, WithUser(r.Request.Session.User.UserID))
		
		if err != nil {
			title := loc.Get(l10n.KeyMissingTranslation_Title)
			text := loc.Get(l10n.KeyMissingTranslation, err)
			if text == "" {
				title = reg.GetDefault().Get(l10n.KeyMissingTranslation_Title)
				text = reg.GetDefault().Get(l10n.KeyMissingTranslation, err)
			}
			b.WithSimpleCard(title, text)
		}
		
		// configure response
		if resp.Speech() != "" {
			b.WithSpeech(resp.Speech())
		}
		b.WithSimpleCard(resp.Title(), resp.Text())
		if resp.End() {
			b.WithShouldEndSession(true)
		}
	})
}

func handleStandardIntent(f AppResponse) alexa.Handler {
    return alexa.HandlerFunc(func(b *alexa.ResponseBuilder, r *alexa.Request) {
        // resolve locale
        def := loca.Registry.GetDefault()
        if def == nil {
        	handleError(b, r, fmt.Errorf("no default locale registered"))
        	return
        }
        loc, err := loca.Registry.Resolve(r.Locale)
        if err != nil {
        	handleError(b, r, err)
        	return
        }
        
        // run func
        resp, err := f(loc, WithUser(r.Request.Session.User.UserID))
        
        if err != nil {
            title := loc.Get(l10n.KeyMissingTranslation_Title)
            text := loc.Get(l10n.KeyMissingTranslation, err)
            img := loc.Get(l10n.KeyMissingTranslationImage)
            if text == "" {
                title = def.Get(l10n.KeyMissingTranslation_Title)
                text = def.Get(l10n.KeyMissingTranslation, err)
                img = def.Get(l10n.KeyMissingTranslationImage)
            }
            b.WithStandardCard(title, text, &alexa.Image{
                SmallImageURL: fmt.Sprintf(img, "small"),
                LargeImageURL: fmt.Sprintf(img, "large"),
            })
            b.WithShouldEndSession(true)
            return
        }
        
        // configure response
        if resp.Speech() != "" {
            b.WithSpeech(resp.Speech())
        }
        b.WithStandardCard(resp.Title(), resp.Text(), &alexa.Image{
            SmallImageURL: fmt.Sprintf(resp.Image(), "small"),
            LargeImageURL: fmt.Sprintf(resp.Image(), "large"),
        })
        if resp.End() {
            b.WithShouldEndSession(true)
        }
    })
}

func localeDefaults(locale string) l10n.LocaleInstance {
	l, err := loca.Registry.Resolve(locale)
	if err != nil {
		l = l10n.NewLocale(locale)
	    loca.Registry.Register(l)
	}
    if l.Get(l10n.KeyErrorTitle) == "" { l.Set(l10n.KeyErrorTitle, []string{"Error"})}
    if l.Get(l10n.KeyErrorText) == "" { l.Set(l10n.KeyErrorText, []string{"The app returned an error:\n%s"})}
    if l.Get(l10n.KeyErrorMissingPlaceholder) == "" { l.Set(l10n.KeyErrorMissingPlaceholder, []string{"the string is missing a placeholder %%s: '%s'"})}
	return l
}
func handleError(b *alexa.ResponseBuilder, r *alexa.Request, err error) {
    l := localeDefaults(r.Locale)
    b.WithSimpleCard(l.GetAny(l10n.KeyErrorTitle), l.GetAny(l10n.KeyErrorText, err))
}
```



# Make it simple!
Simple `key -> []value` lookups
```go
package demo
import (    
    "github.com/drpsychick/go-alexa-lambda/l10n"
)
// de-DE.go
// only key -> value. Convention defines the structure
var deDE = &l10n.Locale{
    TextSnippets: map[string][]string{
        MyIntentTitle: []string{
            "Title",
        },
        MyIntentText: []string{
            "Text",
        },
        MyIntentSSML: []string{
            "<speak>Text</speak>",
        },
        // you can fallback to a different locale
        MyKey:        enUS.GetAll(MyKey),
    },
}
```

We have to define somewhere in code how it will react, so why not keep the link to loca keys?
```go
// app.go
package demo
// Links intent to response (flow)
func (a *Application) handleMyIntent(l l10n.LocaleInstance) (string, string, string) {
    return l.GetSnippet(MyIntentTitle), l.GetSnippet(MyIntentText), l.GetSnippet(MyIntentSSML)
}

// More complex function
func (a *Application) handleComplexIntent(l l10n.LocaleInstance, s Slots, ...) (string, string, string) {
    // do something based on the slots provided
    // trigger reprompt if unclear, ...
    return title, text, ssml // and more: Media (visual Alexa), Sounds, ...?
}
```

## So, where does it go?
Still TODO: we define Intents and Slots for the model and I would like to use this definition for lambda
```go
// app.go
package demo
func (a *Application) initialize() { // or in NewApplication
    // define the skill
    skill := alexa.NewSkill(Type)
    // add general elements which are part of the skill
    skill.AddIntents()
    skill.AddSlots()
    
    // loop over locales and add them
    for _, l := range locales {
        skill.AddLocale(locale) // locale is part of alexa package
        // AddLocale can loop over intents etc. and fetch locale for each element
    }
}
```
