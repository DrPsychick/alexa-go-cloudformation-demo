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
		"MyIntent_Samples":            []string{"say one", "say two"},
		"MY_type_values":              []string{"Value 1", "Value 2"},
		"Skill_Instructions":          []string{"My instructions"},
		l10n.KeySkillName:             []string{"SkillName"},
		l10n.KeySkillDescription:      []string{"SkillDescription"},
		l10n.KeySkillSummary:          []string{"SkillSummary"},
		l10n.KeySkillKeywords:         []string{"Keyword1", "Keyword2"},
		l10n.KeySkillExamplePhrases:   []string{"start me", "boot me up"},
		l10n.KeySkillSmallIconURI:     []string{"https://small"},
		l10n.KeySkillLargeIconURI:     []string{"https://large"},
		l10n.KeySkillPrivacyPolicyURL: []string{"https://policy"},
		l10n.KeySkillTermsOfUse:       []string{"https://toc"},
		"Name":                        []string{"name"},
		"Description":                 []string{"description"},
		"Summary":                     []string{"summary"},
		"Keywords":                    []string{"key", "words"},
		"Examples":                    []string{"say", "something"},
		"SmallIcon":                   []string{"https://small.icon"},
		"LargeIcon":                   []string{"https://large.icon"},
		"Privacy":                     []string{"https://privacy.url"},
		"Terms":                       []string{"https://terms.url"},
		l10n.KeySkillInvocation:       []string{"call me"},
	},
}

func init() {
	registry.Register(enUS, l10n.AsDefault())
}

func TestSetup(t *testing.T) {
	assert.NotEmpty(t, registry.GetLocales())
	assert.NotEmpty(t, registry.GetDefault())
	assert.Equal(t, enUS, registry.GetDefault())
}

func TestSkillBuilder_Build(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry)
	assert.IsType(t, &gen.SkillBuilder{}, sb)

	sk, err := sb.Build()
	assert.NoError(t, err)
	assert.NotEmpty(t, sk)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.NotEmpty(t, string(res))
	assert.NotContains(t, string(res), "null")
	fmt.Printf("Skill: %s\n", string(res))
}

func TestSkillBuilder_WithLocaleTestingInstructionsOverwrite(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry).
		WithTestingInstructions("Skill_Instructions")
	sk, err := sb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "My instructions", sk.Manifest.Publishing.TestingInstructions)

	sb.WithLocaleTestingInstructions("en-US", "New instructions")
	sk, err = sb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "New instructions", sk.Manifest.Publishing.TestingInstructions)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)

	fmt.Printf("Skill: %s\n", string(res))
}

func TestSkillBuilder_Build_Full(t *testing.T) {
	sb := gen.NewSkillBuilder().
		WithLocaleRegistry(registry).
		WithCategory(alexa.CategoryAstrology).
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

	fmt.Printf("Full Skill: %s\n", string(res))

	_, err = sb.BuildModels()
	assert.Error(t, err)
}

func TestSkillBuilder_Build_Locale(t *testing.T) {
	sb := gen.NewSkillBuilder()
	sb.AddLocale("en-US").
		WithLocaleName("my name")
	sb.AddLocale("de-DE").
		WithLocaleName("mein name")
	sk, err := sb.Build()
	assert.NoError(t, err)
	assert.Equal(t, "my name", sk.Manifest.Publishing.Locales["en-US"].Name)

	res, err := json.MarshalIndent(sk, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(res), "my name")

	fmt.Printf("%s\n", string(res))
}

func TestSkillLocaleBuilder_With(t *testing.T) {
	lb := gen.NewSkillLocaleBuilder("en-US").
		WithLocaleRegistry(registry).
		WithName("Name").
		WithDescription("Description").
		WithSummary("Summary").
		WithKeywords("Keywords").
		WithExamples("Examples").
		WithSmallIcon("SmallIcon").
		WithLargeIcon("LargeIcon").
		WithPrivacyURL("Privacy").
		WithTermsURL("Terms")

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

	fmt.Printf("%s\n", string(res))

	pl, err := lb.BuildPrivacyLocale()
	assert.NoError(t, err)

	res, err = json.MarshalIndent(pl, "", "  ")
	assert.NoError(t, err)
	assert.Contains(t, string(res), "privacy.url")
	assert.Contains(t, string(res), "terms.url")

	fmt.Printf("%s\n", string(res))
}

func TestBuildImmutability(t *testing.T) {
	s := gen.NewSkillBuilder()
	s1, _ := s.Build()
	s2, _ := s.Build()

	assert.Equal(t, s1, s2, "Building skill is not immutable!")
}
