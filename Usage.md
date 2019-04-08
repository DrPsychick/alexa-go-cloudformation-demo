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
	locSkill := skill.AddLocale(l.Name)
	locSkill.Name = l.Name
	locSkill.Description = l.Description
	[...]
	
	locModel := locSkill.AddLanguageModel(l.Invocation)
	
	for i, _ := range l.Intents {
	    locInt := locModel.AddIntent(i.Name)
	    for s, _ := range i.Samples {
	    	locInt.AddSample(s)
	    }
	 }
}
```