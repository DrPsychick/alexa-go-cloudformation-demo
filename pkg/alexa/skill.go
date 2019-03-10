package alexa

type skill struct {
	Manifest manifest `json:"manifest"`
}

type manifest struct {
	Version     string     `json:"manifestVersion"`
	Publishing  publishing `json:"publishingInformation"`
	Apis        apis       `json:"apis,omitempty""`
	Permissions []string   `json:"permissions"`
	Privacy     privacy    `json:"privacyAndCompliance"`
}

type publishing struct {
	Locales   map[string]localeDef `json:"locales"`
	Worldwide bool                 `json:"isAvailableWorldwide"`
	Category  string               `json:"category"`
	Countries []country            `json:"distributionCountries"`
}

//type locale struct {
//	German		localeDef	`json:"de-DE,omitempty"`
//	USEnglish	localeDef	`json:"en-US,omitempty"`
//}

type localeDef struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Examples    []string `json:"examplePhrases"`
	Keywords    []string `json:"keywords`
}

type country string

type apis struct {
	Custom     custom   `json:"custom"`
	Interfaces []string `json:"interfaces"`
}
type custom struct {
	Endpoint endpoint `json:"endpoint"`
}
type endpoint struct {
	Uri string `json:"uri"`
}

type privacy struct {
	Compliant bool `json:"isExportCompliant"`
	Ads       bool `json:"containsAds"`
}

var example = skill{
	Manifest: manifest{
		Publishing: publishing{
			Locales: map[string]localeDef{
				"de-DE": {
					Name:        "name",
					Description: "description",
					Summary:     "summary",
					Keywords:    []string{},
					Examples:    []string{},
				},
			},
		},
		Apis: apis{
			Custom: custom{
				Endpoint: endpoint{
					Uri: "arn:...",
				},
			},
			Interfaces: []string{},
		},
	},
}
