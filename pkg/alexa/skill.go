package alexa

// Skill Alexa `skill.json` top element
type Skill struct {
	Manifest Manifest `json:"manifest"`
}

// Manifest definition for `skill.json`
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

// LocaleDef description of each locale
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
	// Country Australia
	CountryAustrialia Country = "AU"
	// Country Canada
	CountryCanada Country = "CA"
	// Country Germany
	CountryGermany Country = "DE"
	// Country Great Britain
	CountryGreatBritain Country = "GB"
	// Country India
	CountryIndia Country = "IN"
	// Country Italia
	CountryItaly Country = "IT"
	// Country Japan
	CountryJapan Country = "JP"
	// Country United States
	CountryUnitedStates Country = "US"
)

// Apis Alexa will be connected to
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
	// Alexa Region North America
	RegionNorthAmerica Region = "NA"
	// Alexa Region Europe
	RegionEurope Region = "EU"
	// Alexa Region Far East
	RegionFarEast Region = "FE"
)

// RegionDef for regional endpoints
type RegionDef struct {
	Endpoint *Endpoint `json:"endpoint"`
}

// Interface definition for API
type Interface struct {
	Type InterfaceType `json:"type"`
}

// InterfaceType string reference
type InterfaceType string

const (
	// Interface Type ???
	InterfaceTypeAlexaPresentationAPL InterfaceType = "ALEXA_PRESENTATION_APL"
	// Interface Type ???
	InterfaceTypeAudioPlayer InterfaceType = "AUDIO_PLAYER"
	// Interface Type for Lambda
	InterfaceTypeCanFulfillIntentRequest InterfaceType = "CAN_FULFILL_INTENT_REQUEST"
	// Interface Type ???
	InterfaceTypeGadgetController InterfaceType = "GADGET_CONTROLLER"
	// Interface Type ???
	InterfaceTypeGameEngine InterfaceType = "GAME_ENGINE"
	// Interface Type ???
	InterfaceTypeRenderTemplate InterfaceType = "RENDER_TEMPLATE"
	// Interface Type ???
	InterfaceTypeVideoApp InterfaceType = "VIDEO_APP"
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
