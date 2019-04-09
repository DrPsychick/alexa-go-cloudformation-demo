package gen_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa"
	"github.com/drpsychick/alexa-go-cloudformation-demo/pkg/alexa/gen"
	"github.com/stretchr/testify/assert"
	"testing"
)

var s = gen.NewSkill()

func TestSkill(t *testing.T) {
	s.SetCategory(alexa.CategoryCommunication)
	assert.NotEmpty(t, s.Category)

}
func TestIntent(t *testing.T) {
	s.AddIntent(gen.NewIntent("Foo"))
	assert.NotEmpty(t, s.Intents)
}

func TestIntentWithSlots(t *testing.T) {
	ty := gen.NewType("MY_Type")
	i := gen.NewIntent("WithSlots")
	i.AddSlot(gen.NewSlot("MySlot-1", ty))
	s.AddIntent(i)
	// 3 basic + 2 added in this test case:
	assert.Equal(t, 5, len(s.Intents))
}

func TestValidateTypes(t *testing.T) {
	assert.Error(t, s.ValidateTypes())
	s.AddTypeString("MY_Type")
	assert.NoError(t, s.ValidateTypes())
}
