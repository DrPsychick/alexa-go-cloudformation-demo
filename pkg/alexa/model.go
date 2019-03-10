package alexa

type Model struct {
	Model InteractionModel `json:"interactionModel"`
}

type InteractionModel struct {
	Language Language      `json:"languageModel"`
	Dialog   Dialog        `json:"dialog"`
	Prompts  []ModelPrompt `json:"prompts"`
}

type Language struct {
	Invocation string        `json:"invocationName"`
	Intents    []ModelIntent `json:"intents"`
	Types      []ModelType   `json:"types,omitempty"`
}

type ModelType struct {
	Name   string      `json:"name"`
	Values []TypeValue `json:"values"`
}

type TypeValue struct {
	Name NameValue `json:"name"`
}

type NameValue struct {
	Value string `json:"value"`
}

type ModelIntent struct {
	Name    string   `json:"name"`
	Samples []string `json:"samples"`
	Slots   []struct {
		Name    string   `json:"name"`
		Type    string   `json:"type"`
		Samples []string `json:"samples"`
	} `json:"slots,omitempty"`
}

type Dialog struct {
	Intents    []DialogIntent `json:"intents"`
	Delegation string         `json:"delegationStrategy"`
}

type DialogIntent struct {
	Name         string `json:"name"`
	Confirmation bool   `json:"confirmationRequired"`
	Prompts      struct {
	} `json:"prompts"`
	Slots []struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Confirmation bool   `json:"confirmationRequired"`
		Elicitation  bool   `json:"elicitationRequired"`
		Prompts      struct {
			Elicitation string `json:"elicitation"`
		} `json:"prompts"`
	} `json:"slots"`
}

type ModelPrompt struct {
	Id         string `json:"id"`
	Variations []struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"variations"`
}
