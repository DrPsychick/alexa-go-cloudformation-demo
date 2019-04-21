# Examples of how to use `alexa` package

## Simple case, Skill with only one language
Building the skill
```go
sb := gen.NewSkillBuilder().
    WithCategory(alexa.CategoryCommunication).
    AddCountry("US")

sb.AddLocale("en-US").
    WithLocaleName("my name").
    WithLocaleDescription("my description").
    WithLocaleSummary("my summary").
    WithLocaleKeywords([]string{"word1", "word2"}).
    WithLocaleExamples([]string{"make an example", "give an example"}).
    WithLocaleSmallIcon("https://small.icon").
    WithLocaleLargeIcon("https://large.icon").
    WithLocalePrivacyURL("https://privacy.url/en-US/")

// must be used *after* adding a locale
err = sb.SetDefaultLocaleTestingInstructions("Foo bar")
[...]
// *alexa.Skill
sk, err := sb.Build()
[...]
res, err := json.MarshalIndent(sk, "", "  ")
[...]
```

Building the model
```go
// you can still use l10n.LocaleRegistry to resolve translations if you wish
registry := l10n.NewRegistry()
err := registry.Register(l10n.NewLocale("en-US"))
[...]
loc, err := registry.Resolve("en-US")
loc.Set("MyIntent_Samples", []string{"sample one", "sample two"})
[...]
// or: gen.NewModelBuilder()
mb := sb.AddModel().
    WithDelegationStrategy(alexa.DelegationSkillResponse).
    AddLocale("en-US", "my skill").
    AddLocale("de-DE", "mein skill")

mb.AddType("TypeSlotOne").
    WithLocaleValues("en-US", []string{"One"}).
    WithLocaleValues("de-DE", []string{"Eins"})

mb.AddIntent("MyIntent").
    WithLocaleSamples(loc.GetName(), loc.GetAll("MyIntent_Samples")).
    WithLocaleSamples("de-DE", []string{"sample eins", "sample zwei"}).
    AddSlot("SlotName", "TypeSlotOne").
    WithLocaleSamples(loc.GetName(), []string{"of {Slot}"}).
    WithLocaleSamples("de-DE", []string{"von {Slot}"})

mb.AddElicitationSlotPrompt("MyIntent", "SlotName").
    AddVariation("PlainText").
    WithLocaleValue("de-DE", "PlainText", []string{"Was?", "Wie bitte?"}).
    WithLocaleValue(loc.GetName(), "PlainText", []string{"What?"})

mb.AddConfirmationSlotPrompt("MyIntent", "SlotName").
    AddVariation("PlainText").
    WithLocaleValue(loc.GetName(), "PlainText", []string{"Sure?"}).
    WithLocaleValue("de-DE", "PlainText", []string{"Sicher?"})

// *alexa.Model
m, err := mb.BuildLocale("en_US")
[...]
res, err := json.MarshalIndent(m, "", "  ")
[...]

// map[locale]*alexa.Model of models
ms, err := mb.Build()
```

## International case, multiple languages
Definining locales
```go
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
        l10n.KeySkillTermsOfUse:          []string{"https://toc"},
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
}
//var deDE = &l10n.Locale{...}
```
Building the skill
```go
// register the locales... first one automatically is default
registry := l10n.NewRegistry()
// there are multiple ways to set the default explicitly
registry.Register(enUS, l10n.AsDefault())
registry.SetDefault("en-US")

// pass the registry
sb := gen.NewSkillBuilder().
    WithLocaleRegistry(registry).
    WithCategory(alexa.CategoryFashionAndStyle)
[...]
// *alexa.Skill
s, err := sb.Build()
```
Building the models
```go
// pass the registry
mb := gen.NewModelBuilder().
    WithLocaleRegistry(registry).
    WithDelegationStrategy(alexa.DelegationSkillResponse)
// add intents, types, slots, prompts, ...
mb.AddType("MyType") // looks up "MyType_Values
mb.AddIntent("MyIntent") // looks up "MyIntent_Samples"
mb.AddIntent("SlotIntent"). // looks up "SlotIntent_Samples"
    AddSlot("SlotName", "MyType") // looks up "SlotIntent_SlotName_Samples"

mb.AddElicitationSlotPrompt("SlotIntent", "SlotName").
    AddVariation("PlainText").
    AddVariation("SSML")

ms, err := mb.Build()
```

## Expert case (build your own)
Simply build your own JSON
```go
var skill = &alexa.Skill{...}
var modelEnUs = &alexa.Model{...}
```
As you're an expert, you can easily figure out how to do that in detail by looking at the tests: 
https://github.com/DrPsychick/alexa-go-cloudformation-demo/blob/master/pkg/alexa/skill_test.go


# Make it simple!
Simple `key -> []value` lookups
```go
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
            "<speak>Text</speak>"
        },
        // you can fallback to a different locale
        MyKey:        enUS.GetAll(MyKey),
    },
}
[...]
```

We have to define somewhere in code how it will react, so why not keep the link to loca keys?
```go
// app.go
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
func (a *Application) initialize() { // or in NewApplication
    // define the skill
    skill := alexa.NewSkill(Type)
    // add general elements which are part of the skill
    skill.AddIntents(...)
    skill.AddSlots(...)
    
    // loop over locales and add them
    for _, l := range locales {
        skill.AddLocale(locale) // locale is part of alexa package
        // AddLocale can loop over intents etc. and fetch locale for each element
    }
}
[...]
```
