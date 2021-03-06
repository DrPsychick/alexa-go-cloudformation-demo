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
	f := &test.Foo{}
	f.SetName("bar1")
	assert.Equal(t, "bar1", f.GetName())

	f.SetName("bar2")
	// Foo stores a pointer to itself by default
	f2 := f.GetInstance()
	assert.Equal(t, "bar2", f2.GetName())
	assert.Equal(t, f, f2)

	b := &test.Bar{}
	b.SetName("bar3")
	// INFO: need to pass as pointer!
	f.SetInstance(b)
	f3 := f.GetInstance()
	assert.Equal(t, "bar3", f3.GetName())
	assert.Equal(t, b, f3)

}
