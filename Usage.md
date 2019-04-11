# Examples of how to use (DRAFT)

## defining the skill
```go
const (
	TypeSlotOne string = "SLOT_one"
	DemoIntent string = "DemoIntent"
	DemoIntentSlotOne string = "DemoIntent_Slot_One"
	DemoIntentSlotOneSamples string = "DemoIntent_Slot_One_Samples"
	MyIntent string = "MyIntent"
)
var deDE = &l10n.Locale{
	Name: "de-DE",
	Countries: []alexa.Country{alexa.CountryGerman},
	TextSnippets: map[string][]string{
		"MyKey": []string{"My Value 1", "My Value 2"},
		DemoIntentTitle: []string{"Demo"},
		DemoIntentSamples: []string{"starte demo"},
		DemoIntentSlotOneSamples: []string{"of {Area}", "in {Area}"},
	},
}
r := NewRegistry().
	WithLocale(deDE)
[...]
// generate new SkillBuilder
skill := gen.NewSkillBuilder().
	WithCategory(alexa.CategoryShopping).
	WithModelDelegation(alexa.DelegationSkillResponse).
	WithPrivacyFlag(gen.FlagIsExportCompliant)
[...]
// add locales
for n, l := range r.GetLocales() {
    skill = skill.WithLocale(gen.NewLocaleBuilder().WithLocale(n, l).
    	WithCountries(l.Countries)
}
[...]
// add types
skill = skill.WithType(loca.SlotOneType)

// add simple intent
skill = skill.WithIntent(gen.NewIntentBuilder(loca.MyIntent))

// add intent with samples and slot
skill = skill.WithIntent(gen.NewIntentBuilder(loca.DemoIntent).
	WithSamples(loca.DemoIntentSamples). // reference key specifically
	WithSlot(gen.NewIntentSlot(
		loca.DemoIntentSlotOne, loca.SlotOneType
    ).WithSlotSamples(local.DemoIntentSlotOneSamples)
))
// OR: WithIntent returns an Intent...
skill = skill.WithIntent(loca.DemoIntent).
	WithSamples(...).WithSlot(...)

s, _ := skill.Build()
res, _ := json.Marshal(s)
[...]

model, _ := skill.BuildModels()
for l, m := range skill.BuildModels() {
	res, _ := json.Marshal(m)
	[...]
}
```

In a loop over locales
```go
type Locale = struct{
	Name: string,
	Invocation: string,
	Description: string,
	IntentResponses: []IntentResponse,
	Prompts: map[Id]Prompt,
	[...]
}

var skill = &Skill{
	Intents: []Intent, // may reference slots
	Types: []Type, // slot types
	Dialog: Dialog, // tied with slot intent prompts
	Prompts: []Prompt, // only list of Ids
}

var enUS = &Locale{
	Name: "en-US",
	Invocation: "my demo",
	Description: "A demo skill",
	
}

var locales = []*Locale{
	enUS, deDE
}
[...]
skill := NewSkill(Type)
skill.AddCategory(...)

int := skill.AddIntent("MyIntent")
int.AddSlot("name", "type")
[...]
skill.AddType("name")
[...]

for l, _ := range locales {
	// 
	locSkill := skill.AddLocale(l.Name)
	locSkill.Name = l.Name
	locSkill.Description = l.Description
	[...]

	// Skill only references Model (separte output JSON files)
	intModel := locSkill.AddInteractionModel()
	
	// InteractionModel -> LanguageModel (with "invocation")
	langModel := intModel.AddLanguageModel(l.Invocation)
	
	// LanguageModel -> []Intent : Intents are defined on the skill
	for _, i := range skill.Intents {
	    locInt := langModel.AddIntent(i.Name)
	    resp := l.IntentResponses[i.Name]
	    
	    // Intent -> []Sample : Samples are defined per locale
	    for _, s := range resp.Samples {
	    	locInt.AddSample(s)
	    }
	    
	    // Intent -> []Slot : Slots are defined on the skill
	    for k, sl := range i.Slots {
	    	slot := locInt.AddSlot(sl.Name, sl.Type)
	    	
	    	// Slot -> []Sample : Slot samples are defined per locale
	    	for _, s = range resp.Slots[k].Samples {
	    		slot.AddSample(s)
	    	}
	    }
	 }
	
	// Model -> []Type : Types are defined on the skill
	for _, t := range skill.Types {
    	type := locSkill.AddType(t.Name)
    	vals := l.Types[t.Name]
    	
    	// Type -> []Value : Values are defined per locale
    	for _, v := range vals {
    		type.AddValue(v)
    	}
	}
	
	// InteractionModel -> Dialog (with "delegationStrategy")
	dia := intModel.AddDialog(skill.Dialog.DelegationStrategy)
	
	// Dialog -> []Intent : Intents are defined on the skill
	// TODO: Dialog has no locale specific information?
	for _, i = range l.Intents {
		// only intents with slots
		if len(i.Slots) == 0 {
			continue
		}
		int := dia.AddIntent(i.Name, i.ConfirmationRequired)
		
		// Intent -> []Prompt
		for p, _ := range i.Prompts {
			int.AddPrompt(p.Name, ...)
		}
		
		// Intent -> []Slot
        for p, _ = range i.Slots {
            int.AddSlot(p.Name, p.Type, ...)
        }
	}
	
	// InteractionModel -> []Prompt : Prompts are defined on the skill
	for _, p = range skill.Prompts {
	    prompt := intModel.AddPrompt(p.Id)
	    
	    // Prompt -> []Variation : Variations are defined per locale
	    for _, v := range l.Prompts[p.Id] {
	    	prompt.AddVariation(v.Type, v.Value)
	    }
	}
}
[...]

skill.RenderSkill()
skill.RenderModel("de-DE")
[...]
```

## Defining a flow -> NO!
```go
var simpleFlow = alexa.DialogFlow{
	Intent: "SaySomething",
	Samples: []string{
		"Say something",
		"Tell me something",
		"What's up",
	},
	Responses: []Response{
		{text: "Yeah, right!", ssml: "<speak>Jop</speak>"},
	}
}
```

### Generic flow, referencing locale -> NO!
```go
var enUS = &alexa.Locale{
    Name: "en-US",
    Intents: []Intent{
        Intent{
            Name: "SaySomething"
            Samples: []string{
                "Say something",
                "Tell me something",
                "What's up",
            },
            Responses: []Response{
                { Text: "Yeah, right!", SSML: "<speak>Jop</speak>" },
            },
        },
    },
}

for _, i := range l.Intents {
    simpleIntent := alexa.SimpleIntent{
    	Intent: i.Name,
    	Samples: i.Samples,
    	Responses: i.Responses,
    }
    app.RegisterIntent(simpleIntent)
}

```

## Make it simple!
```go
// de-DE.go
// only key -> value. Convention defines the structure
var deDE = &l10n.Locale{
	Snippets: [l10n.Key][]string{
		MyIntentTitle: []string{
			"Title",
		},
		MyIntentText: []string{
			"Text",
		},
		MyIntentSSML: []string{
			"<speak>Text</speak>"
		},
	},
}
[...]

// app.go
// Links intent to response (flow)
func (a *Application) handleMyIntent(l *Locale) (string, string, string) {
	return l.GetSnippet(MyIntentTitle), l.GetSnippet(MyIntentText), l.GetSnippet(MyIntentSSML)
}

// More complex function
func (a *Application) handleComplexIntent(l *Locale, s Slots, ...) (string, string, string) {
	// do something based on the slots provided
	// trigger reprompt if unclear, ...
	return title, text, ssml // and more: Media (visual Alexa), Sounds, ...?
}
```


## So, where does it go?
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
