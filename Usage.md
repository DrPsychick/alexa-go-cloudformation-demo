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