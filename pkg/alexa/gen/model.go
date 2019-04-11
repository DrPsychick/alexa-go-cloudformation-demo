package gen

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/l10n"
)

// ModelBuilder builds an alexa.Model instance for a locale
type ModelBuilder struct {
	skillBuilder *SkillBuilder
	locale       *l10n.Locale
	intents      []*ModelIntentBuilder
}

func NewModelBuilder(locale string) *ModelBuilder {
	return &ModelBuilder{}
}

func (m *ModelBuilder) WithIntent(intent string) *ModelIntentBuilder {
	i := NewModelIntentBuilder(intent)
	m.intents = append(m.intents, i)
	return i
}

func (m *ModelBuilder) Build() alexa.Model {
	return alexa.Model{}
}

type ModelIntentBuilder struct {
	name string
}

func NewModelIntentBuilder(intent string) *ModelIntentBuilder {
	return &ModelIntentBuilder{name: intent}
}
