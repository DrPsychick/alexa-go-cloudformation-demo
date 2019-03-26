package alexa

// Skill is the Alexa `skill.json` top element
type Skill struct {
	Manifest Manifest `json:"manifest"`
}

// Manifest is the parent for all other elements
type Manifest struct {
	Version     string        `json:"manifestVersion"`
	Publishing  Publishing    `json:"publishingInformation"`
	Apis        *Apis         `json:"apis,omitempty"`
	Permissions *[]Permission `json:"permissions"`
	Privacy     *Privacy      `json:"privacyAndCompliance"`
}

// Publishing information
type Publishing struct {
	Locales             map[Locale]LocaleDef `json:"locales"`
	Worldwide           bool                 `json:"isAvailableWorldwide"`
	Category            Category             `json:"category"`
	Countries           []Country            `json:"distributionCountries"`
	TestingInstructions string               `json:"testingInstructions"`
}

// LocaleDef description of each locale
type LocaleDef struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Summary      string   `json:"summary"`
	Examples     []string `json:"examplePhrases"`
	Keywords     []string `json:"keywords"`
	SmallIconUri string   `json:"smallIconUri"`
	LargeIconUri string   `json:"largeIconUri"`
}

// Country constants
type Country string

const (
	// CountryAustralia is AU
	CountryAustrialia Country = "AU"
	// CountryCanada is CA
	CountryCanada Country = "CA"
	// CountryGermany is DE
	CountryGermany Country = "DE"
	// CountryGreatBritain is GB
	CountryGreatBritain Country = "GB"
	// CountryIndia is IN
	CountryIndia Country = "IN"
	// CountryItalia is IT
	CountryItaly Country = "IT"
	// CountryJapan is JP
	CountryJapan Country = "JP"
	// CountryUnitedStates is US
	CountryUnitedStates Country = "US"
)

// Category of the Skill
type Category string

const (
	//ALARMS_AND_CLOCKS
	//ASTROLOGY
	//BUSINESS_AND_FINANCE
	//CALCULATORS
	//CALENDARS_AND_REMINDERS
	//CHILDRENS_EDUCATION_AND_REFERENCE
	//CHILDRENS_GAMES
	//CHILDRENS_MUSIC_AND_AUDIO
	//CHILDRENS_NOVELTY_AND_HUMOR
	//COMMUNICATION
	//CONNECTED_CAR
	//COOKING_AND_RECIPE
	//CURRENCY_GUIDES_AND_CONVERTERS
	//DATING
	//DELIVERY_AND_TAKEOUT
	//DEVICE_TRACKING
	//EDUCATION_AND_REFERENCE
	//EVENT_FINDERS
	//EXERCISE_AND_WORKOUT
	//FASHION_AND_STYLE
	//FLIGHT_FINDERS
	//FRIENDS_AND_FAMILY
	//GAME_INFO_AND_ACCESSORY
	//GAMES
	//HEALTH_AND_FITNESS
	//HOTEL_FINDERS
	//KNOWLEDGE_AND_TRIVIA
	//MOVIE_AND_TV_KNOWLEDGE_AND_TRIVIA
	//MOVIE_INFO_AND_REVIEWS
	//MOVIE_SHOWTIMES
	//MUSIC_AND_AUDIO_ACCESSORIES
	//MUSIC_AND_AUDIO_KNOWLEDGE_AND_TRIVIA
	//MUSIC_INFO_REVIEWS_AND_RECOGNITION_SERVICE
	//NAVIGATION_AND_TRIP_PLANNER
	//NEWS
	//NOVELTY
	// CategoryOrganizersAndAssistants is ORGANIZERS_AND_ASSISTANTS
	CategoryOrganizersAndAssistants Category = "ORGANIZERS_AND_ASSISTANTS"
	//PETS_AND_ANIMAL
	//PODCAST
	//PUBLIC_TRANSPORTATION
	//RELIGION_AND_SPIRITUALITY
	//RESTAURANT_BOOKING_INFO_AND_REVIEW
	//SCHOOLS
	//SCORE_KEEPING
	//SELF_IMPROVEMENT
	//SHOPPING
	//SMART_HOME
	//SOCIAL_NETWORKING
	//SPORTS_GAMES
	//SPORTS_NEWS
	//STREAMING_SERVICE
	//TAXI_AND_RIDESHARING
	//TO_DO_LISTS_AND_NOTES
	//TRANSLATORS
	//TV_GUIDES
	//UNIT_CONVERTERS
	//WEATHER
	//WINE_AND_BEVERAGE
	//ZIP_CODE_LOOKUP
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

// Region for Alexa
type Region string

const (
	// RegionNorthAmerica is NA
	RegionNorthAmerica Region = "NA"
	// RegionEurope is EU
	RegionEurope Region = "EU"
	// RegionFarEast is FE
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
	// InterfaceTypeAlexaPresentationAPL is ALEXA_PRESENTATION_APL
	InterfaceTypeAlexaPresentationAPL InterfaceType = "ALEXA_PRESENTATION_APL"
	// InterfaceTypeAudioPlayer is AUDIO_PLAYER
	InterfaceTypeAudioPlayer InterfaceType = "AUDIO_PLAYER"
	// InterfaceTypeCanFulfillIntentRequest is CAN_FULFILL_INTENT_REQUEST
	InterfaceTypeCanFulfillIntentRequest InterfaceType = "CAN_FULFILL_INTENT_REQUEST"
	// InterfaceTypeGadgetController is GADGET_CONTROLLER
	InterfaceTypeGadgetController InterfaceType = "GADGET_CONTROLLER"
	// InterfaceTypeGameEngine is GAME_ENGINE
	InterfaceTypeGameEngine InterfaceType = "GAME_ENGINE"
	// InterfaceTypeRenderTemplate is RENDER_TEMPLATE
	InterfaceTypeRenderTemplate InterfaceType = "RENDER_TEMPLATE"
	// InterfaceTypeVideoApp is VIDEO_APP
	InterfaceTypeVideoApp InterfaceType = "VIDEO_APP"
)

// Permission string
type Permission struct {
	Name string `json:"name"`
}

// Privacy definition
type Privacy struct {
	IsExportCompliant bool                        `json:"isExportCompliant"`
	ContainsAds       bool                        `json:"containsAds"`
	AllowsPurchases   bool                        `json:"allowsPurchases"`
	UsesPersonalInfo  bool                        `json:"usesPersonalInfo"`
	IsChildDirected   bool                        `json:"isChildDirected"`
	Locales           map[Locale]PrivacyLocaleDef `json:"locales,omitempty"`
}

// PrivacyLocaleDef
type PrivacyLocaleDef struct {
	PrivacyPolicyUrl string `json:"privacyPolicyUrl"`
	TermsOfUse       string `json:"termsOfUse"`
}
