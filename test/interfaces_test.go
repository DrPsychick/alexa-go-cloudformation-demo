package test_test

import (
	"github.com/drpsychick/alexa-go-cloudformation-demo/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	f := test.Foo{}
	f.SetName("bar")
	assert.Equal(t, "bar", f.GetName(), "Original unchanged")
}

func TestBar(t *testing.T) {
	// SetName is pointer reciever -> it works
	f := test.Foo{}
	f.SetName("bar")
	assert.NotEmpty(t, f.GetName())

	// Foo stores a pointer to itself by default
	f2 := f.GetInstance()
	assert.NotEmpty(t, f2.GetName())

	b := test.Bar{}
	b.SetName("bar")
	// will not work:
	//
	//bi := test.MySubInterface(b)
	f.SetInstance(b)
	f3 := f.GetInstance()
	assert.NotEmpty(t, f3.GetName())

}
