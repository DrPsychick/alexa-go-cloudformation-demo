package alexa

// Model is the root of an interactionModel
type Model struct {
	Model InteractionModel `json:"interactionModel"`
}

// InteractionModel defines the base model structure
type InteractionModel struct {
	Language LanguageModel `json:"languageModel"`
	Dialog   *Dialog       `json:"dialog,omitempty"`
	Prompts  []ModelPrompt `json:"prompts,omitempty"`
}

// LanguageModel
type LanguageModel struct {
	Invocation string        `json:"invocationName"`
	Intents    []ModelIntent `json:"intents"`
	Types      []ModelType   `json:"types,omitempty"`
}

type ModelIntent struct {
	Name    string      `json:"name"`
	Samples []string    `json:"samples,omitempty"`
	Slots   []ModelSlot `json:"slots,omitempty"`
}

type ModelSlot struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Samples []string `json:"samples,omitempty"`
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
	Delegation string         `json:"delegationStrategy"`
	Intents    []DialogIntent `json:"intents,omitempty"`
}

const (
	DelegationSkillResponse string = "SKILL_RESPONSE"
	DelegationAlways        string = "ALWAYS"
)

type DialogIntent struct {
	Name         string             `json:"name"`
	Confirmation bool               `json:"confirmationRequired"`
	Delegation   string             `json:"delegationStrategy,omitempty"`
	Prompts      IntentPrompt       `json:"prompts,omitempty"`
	Slots        []DialogIntentSlot `json:"slots,omitempty"`
}

type IntentPrompt struct {
	Confirmation string `json:"confirmation,omitempty"`
}

type DialogIntentSlot struct {
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Confirmation bool             `json:"confirmationRequired"`
	Elicitation  bool             `json:"elicitationRequired"`
	Prompts      SlotPrompts      `json:"prompts,omitempty"`
	Validations  []SlotValidation `json:"validations,omitempty"`
}

type SlotPrompts struct {
	Elicitation  string `json:"elicitation,omitempty"`
	Confirmation string `json:"confirmation,omitempty"`
}

// see https://developer.amazon.com/docs/custom-skills/validate-slot-values.html#validation-rules
const (
	ValidationTypeHasMatch      string = "hasEntityResolutionMatch"
	ValidationTypeInSet         string = "isInSet"
	ValidationTypeNotInSet      string = "isNotInSet"
	ValidationTypeGreaterThan   string = "isGreaterThan"
	ValidationTypeGreaterEqal   string = "isGreaterThanOrEqualTo"
	ValidationTypeLessThan      string = "isLessThan"
	ValidationTypeLessEqual     string = "isLessThanOrEqualTo"
	ValidationTypeInDuration    string = "isInDuration"
	ValidationTypeNotInDuration string = "isNotInDuration"
)

type SlotValidation struct {
	Type   string   `json:"type"`   // see https://developer.amazon.com/docs/custom-skills/validate-slot-values.html#validation-rules
	Prompt string   `json:"prompt"` //
	Values []string `json:"values,omitempty"`
}

type ModelPrompt struct {
	Id         string            `json:"id"`
	Variations []PromptVariation `json:"variations"`
}

type PromptVariation struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
