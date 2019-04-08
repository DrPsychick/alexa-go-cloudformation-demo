# Examples of how to use (DRAFT)

## defining the skill
```go
skill := NewSkill(alexa.SkillType)
skill.AddCategory(alexa.CategoryShopping)
[...]
deSkill := skill.AddLocale("de-DE")
deSkill.Name = "Foo"
deSkill.Description = "Bar"
[...]

deModel := deSkill.AddLanguageModel("my skill") // invocation name needed only once per locale

int := deModel.AddIntent("DemoIntent")
int.AddSample("Erzähl mir was")
[...]

```

In a loop over locales
```go
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
	
	// LanguageModel -> []Intent
	for i, _ := range l.Intents {
	    locInt := langModel.AddIntent(i.Name)
	    
	    // Intent -> []Sample
	    for s, _ := range i.Samples {
	    	locInt.AddSample(s)
	    }
	    
	    // Intent -> []Slot -> []Sample
	    for sl, _ := range i.Slots {
	    	slot := locInt.AddSlot(sl)
	    	
	    	for s, _ = range sl.Samples {
	    		slot.AddSample(s)
	    	}
	    }
	 }
	
	// Model -> []Type
	for t, _ := range l.Types {
    	type := locSkill.AddType(t.Name)
    	
    	for v, _ := range t.Values {
    		type.AddValue(v.Value)
    	}
	}
	
	// InteractionModel -> Dialog (with "delegationStrategy")
	dia := intModel.AddDialog(l.DelegationStrategy)
	
	// Dialog -> []Intent
	for i, _ = range l.Intents {
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
	
	// InteractionModel -> []Prompt
	for p, _ = range l.Prompts {
	    prompt := intModel.AddPrompt(p.Id)
	    
	    // Prompt -> []Variation
	    for v, _ := range p.Variations {
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
