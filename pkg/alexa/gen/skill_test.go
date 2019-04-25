package gen_test

import (
	"encoding/json"
	"fmt"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
	"github.com/stretchr/testify/assert"
	"testing"
)

var registry = l10n.NewRegistry()

var enUS = &l10n.Locale{
	Name: "en-US",
	TextSnippets: map[string][]string{
		// Skill
		l10n.KeySkillTestingInstructions: {"Initial instructions"},
		"Skill_Instructions":             {"My instructions"},
		l10n.KeySkillName:                {"SkillName"},
		l10n.KeySkillDescription:         {"SkillDescription"},
		l10n.KeySkillSummary:             {"SkillSummary"},
		l10n.KeySkillKeywords:            {"Keyword1", "Keyword2"},
		l10n.KeySkillExamplePhrases:      {"start me", "boot me up"},
		l10n.KeySkillSmallIconURI:        {"https://small"},
		l10n.KeySkillLargeIconURI:        {"https://large"},
		l10n.KeySkillPrivacyPolicyURL:    {"https://policy"},
		//l10n.KeySkillTermsOfUseURL:       {"https://toc"},
		l10n.KeySkillInvocation: {"call me"},
		"Name":                  {"name"},
		"Description":           {"description"},
		"Summary":               {"summary"},
		"Keywords":              {"key", "words"},
		"Examples":              {"say", "something"},
		"SmallIcon":             {"https://small.icon"},
		"LargeIcon":             {"https://large.icon"},
		"Privacy":               {"https://privacy.url"},
		"Terms":                 {"https://terms.url"},
		// Model
		// Intents
		"MyIntent_Samples":                {"say one", "say two"},
		"MyIntent_Title":                  {"Title"},
		"MyIntent_Text":                   {"Text1", "Text2"},
		"MyIntent_SSML":                   {l10n.Speak("SSML one"), l10n.Speak("SSML two")},
		"SlotIntent_Samples":              {"what about slot {SlotName}"},
		"SlotIntent_Title":                {"Test intent with slot"},
		"SlotIntent_Text":                 {"it seems to work"},
		"SlotIntent_SlotName_Samples":     {"of {SlotName}", "{SlotName}"},
		"SlotIntent_SlotName_Elicit_Text": {"Which slot did you mean?", "I did not understand, which slot?"},
		"SlotIntent_SlotName_Elicit_SSML": {l10n.Speak("I'm sorry, which slot did you mean?")},
		// Types
		"MyType_Values": {"Value 1", "Value 2"},
	},
}

func init() {
	if err := registry.Register(enUS, l10n.AsDefault()); err != nil {
		panic("something went horribly wrong")
	}

}

func TestSetup(t *testing.T) {
	assert.NotEmpty(t, registry.GetLocales())
	assert.NotEmpty(t, registry.GetDefault())
	assert.Equal(t, enUS, registry.GetDefault())
}

// Case 1: input multiple languages directly
func TestSkillBuilder_Build_Locale(t *testing.T) {
	sb1 := gen.NewSkillBuilder().
		WithCategory(alexa.CategoryCommunication).
		AddCountry("US")

	// fails: no default locale
	err := sb1.SetDefaultLocaleTestingInstructions("Foo bar")
	assert.Error(t, err)
	// fails: no locales registered
	err = sb1.SetDefaultLocale("en-US")
	assert.Error(t, err)

	// first locale is automatically the default (TODO: remove this magic)
	sb1.AddLocale("en-US").
		WithLocaleName("my name").
		WithLocaleDescription("my description").
		WithLocaleSummary("my summary").
		WithLocaleKeywords([]string{"word1", "word2"}).
		WithLocaleExamples([]string{"make an example", "give an example"}).
		WithLocaleSmallIcon("https://small.icon").
		WithLocaleLargeIcon("https://large.icon").
		WithLocalePrivacyURL("https://privacy.url/en-US/")
	err = sb1.SetDefaultLocale("en-US")
	assert.NoError(t, err)

	// this can only be called after adding a locale! TODO: make this part of Locale
	err = sb1.SetDefaultLocaleTestingInstructions("Foo bar")
	assert.NoError(t, err)

	de := sb1.AddLocale("de-DE").
		WithLocaleName("mein name").
		WithLocaleDescription("beschreibung").
		WithLocaleSummary("zusammenfassung").
		WithLocalePrivacyURL("https://privacy.url/de-DE/")
	// fails, missing icons
	sk, err := sb1.Build()
	assert.Error(t, err)

	de.WithLocaleSmallIcon("https://small.icon").
		WithLocaleLargeIcon("https://large.icon")
	sk, err = sb1.Build()
	assert.NoError(t, err)

	pl := sk.Manifest.Publishing.Locales
	pr := sk.Manifest.Privacy.Locales
	assert.Equal(t, "my name", pl["en-US"].Name)
	assert.Equal(t, "my description", pl["en-US"].Description)
	assert.Equal(t, "my summary", pl["en-US"].Summary)
	assert.Equal(t, "zusammenfassung", pl["de-DE"].Summary)
	assert.Equal(t, []string{"word1", "word2"}, pl["en-US"].Keywords)
	assert.Equal(t, []string{"make an example", "give an example"}, pl["en-US"].Examples)
	assert.Equal(t, "https://privacy.url/en-US/", pr["en-US"].PrivacyPolicyURL)
	assert.Equal(t, "https://privacy.url/de-DE/", pr["de-DE"].PrivacyPolicyURL)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(res), "my name")

	assert.NoError(t, testBuildImmutability(sb1))

	m := sb1.AddModel()
	assert.IsType(t, &gen.ModelBuilder{}, m)
	ms, err := sb1.BuildModels()
	assert.NoError(t, err)
	assert.IsType(t, map[string]*alexa.Model{}, ms)

	fmt.Printf("%s\n", string(res))
}

// Case 2: input LocaleRegistry with locales that use required keys
func TestSkillBuilder_Build(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry).
		WithCategory(alexa.CategoryFashionAndStyle)
	assert.IsType(t, &gen.SkillBuilder{}, sb)

	sk, err := sb.Build()
	assert.NoError(t, err)
	assert.NotEmpty(t, sk)
	assert.Equal(t, "Initial instructions", sk.Manifest.Publishing.TestingInstructions)

	// use our own l10n key for instructions
	sb.WithTestingInstructions("Skill_Instructions")
	sk, err = sb.Build()
	assert.NoError(t, err)
	assert.NotEmpty(t, sk)
	assert.Equal(t, "My instructions", sk.Manifest.Publishing.TestingInstructions)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	assert.NotContains(t, string(res), "null")

	assert.NoError(t, testBuildImmutability(sb))

	fmt.Printf("%s\n", string(res))
}

// SkillBuilder individual functions
func TestSkillBuilder_WithLocaleTestingInstructionsOverwrite(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry).
		WithCategory(alexa.CategoryDeliveryAndTakeout).
		WithTestingInstructions("Skill_Instructions")
	sk, err := sb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "My instructions", sk.Manifest.Publishing.TestingInstructions)

	// overwrite testing instructions
	err = sb.SetDefaultLocaleTestingInstructions("New instructions")
	assert.NoError(t, err)
	sk, err = sb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "New instructions", sk.Manifest.Publishing.TestingInstructions)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))

	assert.NoError(t, testBuildImmutability(sb))
}

// SkillBuilder cover all functions
func TestSkillBuilder_Build_Full(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry).
		WithCategory(alexa.CategoryAstrology).
		WithTestingInstructions("Skill_Instructions").
		WithPrivacyFlag(gen.FlagIsExportCompliant, true).
		WithPrivacyFlag(gen.FlagUsesPersonalInfo, true).
		WithPrivacyFlag(gen.FlagAllowsPurchases, true).
		WithPrivacyFlag(gen.FlagContainsAds, true).
		WithPrivacyFlag(gen.FlagIsChildDirected, true).
		AddCountry(alexa.CountryFrance).
		WithCountries([]string{alexa.CountryGermany}).
		AddCountries([]string{alexa.CountryCanada, alexa.CountryAustralia})
	assert.NotEmpty(t, registry.GetLocales())

	sk, err := sb.Build()
	assert.NoError(t, err)
	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))

	assert.Contains(t, string(res), alexa.CategoryAstrology)
	assert.NotContains(t, string(res), "FR")
	assert.Contains(t, string(res), "DE")
	assert.Contains(t, string(res), "CA")
	assert.Contains(t, string(res), "AU")

	assert.NoError(t, testBuildImmutability(sb))

	fmt.Printf("Full Skill: %s\n", string(res))

	_, err = sb.BuildModels()
	assert.Error(t, err)
}

// Cover skill validations
func TestSkillRestrictions(t *testing.T) {
	sb := gen.NewSkillBuilder()

	// no locales
	_, err := sb.Build()
	assert.Error(t, err)

	// no category
	en := sb.AddLocale("en-US")
	_, err = sb.Build()
	assert.Error(t, err)

	// no testing instructions (missing translation)
	sb.WithCategory(alexa.CategoryDating)
	_, err = sb.Build()
	assert.Error(t, err)

	// missing skill fields
	err = sb.SetDefaultLocaleTestingInstructions("some text")
	assert.NoError(t, err)
	_, err = sb.Build()
	assert.Error(t, err)

	en.WithLocaleName("Name").
		WithLocaleDescription("Description").
		WithLocaleSummary("Summary").
		WithLocaleKeywords([]string{"keyword"}).
		WithLocaleSmallIcon("https://small").
		WithLocaleLargeIcon("https://large")

	// max is 3 example phrases
	en.WithLocaleExamples([]string{"1", "2", "3", "4"})
	_, err = sb.Build()
	assert.Error(t, err)

	// max is 3 keywords
	en.WithLocaleExamples([]string{"1", "2", "3"})
	en.WithLocaleKeywords([]string{"1", "2", "3", "4"})
	_, err = sb.Build()
	assert.Error(t, err)

	// termsOfUse not allowed (yet)
	en.WithLocaleKeywords([]string{"1", "2", "3"})
	en.WithTermsURL(l10n.KeySkillTermsOfUseURL)
	en.WithLocaleTermsURL("http://terms")
	_, err = sb.Build()
	assert.Error(t, err)

	// now it builds...
	en.WithLocaleTermsURL("")
	_, err = sb.Build()
	assert.NoError(t, err)
}

// Cover error cases
func TestErrors(t *testing.T) {
	sb := gen.NewSkillBuilder()
	// AddLocale
	sb2 := sb.AddLocale("en-US")
	assert.IsType(t, &gen.SkillLocaleBuilder{}, sb2)
	sb3 := sb.AddLocale("en-US")
	assert.Nil(t, sb3)

	// fuck it up hard: no locale set in locale builder
	l := &gen.SkillLocaleBuilder{}
	l.WithLocaleRegistry(l10n.NewRegistry())
	l2 := l.WithLocaleName("foo")
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleSummary("foo")
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleDescription("foo")
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleExamples([]string{"foo", "bar"})
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleKeywords([]string{"foo", "bar"})
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleSmallIcon("https://foo")
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleLargeIcon("https://bar")
	assert.Equal(t, l, l2)
	l2 = l.WithLocalePrivacyURL("https://foo")
	assert.Equal(t, l, l2)
	l2 = l.WithLocaleTermsURL("https://bar")
	assert.Equal(t, l, l2)

	pl, err := l.BuildPublishingLocale()
	assert.Error(t, err)
	assert.Empty(t, pl)

	pl2, err := l.BuildPrivacyLocale()
	assert.Error(t, err)
	assert.Empty(t, pl2)
}

// SkillLocaleBuilder Case 1
func TestSkillLocaleBuilder_WithLocale(t *testing.T) {
	lb := gen.NewSkillLocaleBuilder("en-US").
		WithName("SkillName").
		WithLocaleName("My Skill").
		WithLocaleDescription("description").
		WithLocaleSummary("summary").
		WithLocaleSmallIcon("https://small.icon").
		WithLocaleLargeIcon("https://large.icon")

	l, err := lb.BuildPublishingLocale()
	assert.NoError(t, err)
	assert.Equal(t, "My Skill", l.Name)
	res, err := json.MarshalIndent(l, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
}

// SkillLocaleBuilder Case 2
func TestSkillLocaleBuilder_WithRegistry(t *testing.T) {
	lb := gen.NewSkillLocaleBuilder("en-US").
		WithLocaleRegistry(registry).
		WithName("Name").
		WithDescription("Description").
		WithSummary("Summary").
		WithKeywords("Keywords").
		WithExamples("Examples").
		WithSmallIcon("SmallIcon").
		WithLargeIcon("LargeIcon").
		WithPrivacyURL("Privacy")
		//WithTermsURL("Terms")

	l, err := lb.BuildPublishingLocale()
	assert.NoError(t, err)

	res, err := json.MarshalIndent(l, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(res), ": \"name\"")
	assert.Contains(t, string(res), ": \"description\"")
	assert.Contains(t, string(res), ": \"summary\"")
	assert.Contains(t, string(res), "\"say\",")
	assert.Contains(t, string(res), "\"key\",")
	assert.Contains(t, string(res), "small.icon")
	assert.Contains(t, string(res), "large.icon")

	pl, err := lb.BuildPrivacyLocale()
	assert.NoError(t, err)

	res, err = json.MarshalIndent(pl, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(res), "privacy.url")
	//assert.Contains(t, string(res), "terms.url")
}

// helper to compare two builds
func testBuildImmutability(s *gen.SkillBuilder) error {
	s1, err := s.Build()
	if err != nil {
		return err
	}
	res1, err := json.MarshalIndent(s1, "", "  ")
	if err != nil {
		return err
	}

	s2, err := s.Build()
	if err != nil {
		return err
	}
	res2, err := json.MarshalIndent(s2, "", "  ")
	if err != nil {
		return err
	}
	if string(res1) != string(res2) {
		return fmt.Errorf("Building skill is not immutable!\n%+v\n%+v", string(res1), string(res2))
	}
	return nil
}
