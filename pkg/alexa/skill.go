package alexa

// Alexa `skill.json` top element
type Skill struct {
	Manifest Manifest `json:"manifest"`
}

// `skill.json` manifest definition
type Manifest struct {
	Version     string        `json:"manifestVersion"`
	Publishing  Publishing    `json:"publishingInformation"`
	Apis        *Apis         `json:"apis,omitempty"`
	Permissions *[]Permission `json:"permissions"`
	Privacy     *Privacy      `json:"privacyAndCompliance"`
}

// Publishing information
type Publishing struct {
	Locales   map[Locale]LocaleDef `json:"locales"`
	Worldwide bool                 `json:"isAvailableWorldwide"`
	Category  string               `json:"category"`
	Countries []Country            `json:"distributionCountries"`
}

// Description for each locale
type LocaleDef struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Summary     string   `json:"summary"`
	Examples    []string `json:"examplePhrases"`
	Keywords    []string `json:"keywords"`
}

// Country constants
type Country string

const (
	CountryAustrialia   Country = "AU"
	CountryCanada       Country = "CA"
	CountryGermany      Country = "DE"
	CountryGreatBritain Country = "GB"
	CountryIndia        Country = "IN"
	CountryItaly        Country = "IT"
	CountryJapan        Country = "JP"
	CountryUnitedStates Country = "US"
)

// Alexa connected APIs
type Apis struct {
	Custom *Custom `json:"custom"`
	//FlashBriefing *FlashBriefing `json:"flashBriefing"`
	//Health     *Health	`json:"health"`
	Interfaces *[]string `json:"interfaces"`
}

// Custom API endpoint
type Custom struct {
	Endpoint   *Endpoint             `json:"endpoint"`
	Regions    *map[Region]RegionDef `json:"regions,omitempty"`
	Interfaces *[]Interface          `json:"interfaces"`
}

// Endpoint definition
type Endpoint struct {
	Uri                string `json:"uri"`
	SslCertificateType string `json:"sslCertificateType,omitempty"`
}

type Region string

const (
	RegionNorthAmerica Region = "NA"
	RegionEurope       Region = "EU"
	RegionFarEast      Region = "FE"
)

type RegionDef struct {
	Endpoint *Endpoint `json:"endpoint"`
}

// API interface
type Interface struct {
	Type InterfaceType `json:"type"`
}

type InterfaceType string

const (
	InterfaceTypeAlexaPresentationAPL    InterfaceType = "ALEXA_PRESENTATION_APL"
	InterfaceTypeAudioPlayer             InterfaceType = "AUDIO_PLAYER"
	InterfaceTypeCanFulfillIntentRequest InterfaceType = "CAN_FULFILL_INTENT_REQUEST"
	InterfaceTypeGadgetController        InterfaceType = "GADGET_CONTROLLER"
	InterfaceTypeGameEngine              InterfaceType = "GAME_ENGINE"
	InterfaceTypeRenderTemplate          InterfaceType = "RENDER_TEMPLATE"
	InterfaceTypeVideoApp                InterfaceType = "VIDEO_APP"
)

// Permission string
type Permission struct {
	Name string `json:"name"`
}

// Privacy definition
type Privacy struct {
	IsExportCompliant bool                         `json:"isExportCompliant"`
	ContainsAds       bool                         `json:"containsAds"`
	AllowsPurchases   bool                         `json:"allowsPurchases"`
	UsesPersonalInfo  bool                         `json:"usesPersonalInfo"`
	IsChildDirected   bool                         `json:"isChildDirected"`
	Locales           *map[Locale]PrivacyLocaleDef `json:"locales"`
}

type PrivacyLocaleDef struct {
	PrivacyPolicyUrl string `json:"privacyPolicyUrl"`
	TermsOfUse       string `json:"termsOfUse"`
}
