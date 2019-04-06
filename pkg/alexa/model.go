package alexa

// Model is the root of an interactionModel
type Model struct {
	Model InteractionModel `json:"interactionModel"`
}

// InteractionModel defines the base model structure
type InteractionModel struct {
	Language LanguageModel  `json:"languageModel"`
	Dialog   *Dialog        `json:"dialog"`
	Prompts  *[]ModelPrompt `json:"prompts"`
}

// LanguageModel
type LanguageModel struct {
	Invocation string        `json:"invocationName"`
	Intents    []ModelIntent `json:"intents"`
	Types      []ModelType   `json:"types,omitempty"`
}

type ModelIntent struct {
	Name    string       `json:"name"`
	Samples []string     `json:"samples"`
	Slots   *[]ModelSlot `json:"slots,omitempty"`
}

type ModelSlot struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Samples []string `json:"samples"`
}

type ModelType struct {
	Name   string      `json:"name"`
	Values []TypeValue `json:"values"`
}

type TypeValue struct {
	Id   string    `json:"id,omitempty"`
	Name NameValue `json:"name"`
}

type NameValue struct {
	Value    string   `json:"value"`
	Synonyms []string `json:"synonyms,omitempty"`
}

type Dialog struct {
	Delegation DialogDelegation `json:"delegationStrategy"`
	Intents    *[]DialogIntent  `json:"intents"`
}

type DialogDelegation string

const (
	SkillResponse DialogDelegation = "SKILL_RESPONSE"
)

type DialogIntent struct {
	Name         string             `json:"name"`
	Confirmation bool               `json:"confirmationRequired"`
	Prompts      struct{}           `json:"prompts"`
	Slots        []DialogIntentSlot `json:"slots"`
}

type DialogIntentSlot struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Confirmation bool        `json:"confirmationRequired"`
	Elicitation  bool        `json:"elicitationRequired"`
	Prompts      SlotPrompts `json:"prompts,omitempty"`
}

type SlotPrompts struct {
	Elicitation string `json:"elicitation"`
}

type ModelPrompt struct {
	Id         string             `json:"id"`
	Variations []PromptVariations `json:"variations"`
}

type PromptVariations struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
