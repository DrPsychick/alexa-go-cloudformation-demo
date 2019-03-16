package alexa

type Skill struct {
	Manifest Manifest `json:"manifest"`
}

type Manifest struct {
	Version     string     `json:"manifestVersion"`
	Publishing  Publishing `json:"publishingInformation"`
	Apis        *Apis      `json:"apis,omitempty"`
	Permissions []string   `json:"permissions"`
	Privacy     Privacy    `json:"privacyAndCompliance"`
}

type Publishing struct {
	Locales   map[Locale]LocaleDef `json:"locales"`
	Worldwide bool                 `json:"isAvailableWorldwide"`
	Category  string               `json:"category"`
	Countries []Country            `json:"distributionCountries"`
}

type Locale string

const (
	German          Locale = "de-DE"
	AmericanEnglish Locale = "en-US"
)

//type locale struct {
//	German		localeDef	`json:"de-DE,omitempty"`
//	USEnglish	localeDef	`json:"en-US,omitempty"`
//}

type LocaleDef struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Examples    []string `json:"examplePhrases"`
	Keywords    []string `json:"keywords"`
}

type Country string

type Apis struct {
	Custom     *Custom  `json:"custom"`
	Interfaces []string `json:"interfaces"`
}
type Custom struct {
	Endpoint *Endpoint `json:"endpoint"`
}
type Endpoint struct {
	Uri string `json:"uri"`
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
