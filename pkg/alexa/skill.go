package alexa

type Skill struct {
	Manifest Manifest `json:"manifest"`
}

type Manifest struct {
	Version     string     `json:"manifestVersion"`
	Publishing  Publishing `json:"publishingInformation"`
	Apis        Apis       `json:"apis,omitempty""`
	Permissions []string   `json:"permissions"`
	Privacy     Privacy    `json:"privacyAndCompliance"`
}

type Publishing struct {
	Locales   map[string]LocaleDef `json:"locales"`
	Worldwide bool                 `json:"isAvailableWorldwide"`
	Category  string               `json:"category"`
	Countries []Country            `json:"distributionCountries"`
}

//type locale struct {
//	German		localeDef	`json:"de-DE,omitempty"`
//	USEnglish	localeDef	`json:"en-US,omitempty"`
//}

type LocaleDef struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Examples    []string `json:"examplePhrases"`
	Keywords    []string `json:"keywords`
}

type Country string

type Apis struct {
	Custom     Custom   `json:"custom,omitempty"`
	Interfaces []string `json:"interfaces,omitempty"`
}
type Custom struct {
	Endpoint Endpoint `json:"endpoint,omitempty"`
}
type Endpoint struct {
	Uri string `json:"uri,omitempty"`
}

type Privacy struct {
	Compliant bool `json:"isExportCompliant"`
	Ads       bool `json:"containsAds"`
}

/*
var Example = Skill{
	Manifest: Manifest{
		Version: "1.0",
		Publishing: Publishing{
			Locales: map[string]LocaleDef{
				"de-DE": {
					Name:        "name",
					Description: "description",
					Summary:     "summary",
					Keywords:    []string{"Demo"},
					Examples:    []string{"tell me how much beer people drink in germany"},
				},
			},
			Category: "mycategory",
			Countries: []Country{"DE"},
		},
		Apis: Apis{
			Custom: Custom{
				Endpoint: Endpoint{
					Uri: "arn:...",
				},
			},
			Interfaces: []string{},
		},
		Permissions: []string{},
		Privacy: Privacy{
			Compliant: true,
		},
	},
}
*/
