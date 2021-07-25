package test

/* Use case:
Two objects types (interfaces), where one uses the other.
Both shall be open for a specific implementation and defined as interfaces.
*/

// MyInterface is a test interface
type MyInterface interface {
	SetName(name string)
	// CANNOT be pointer!
	// link to an object complying to the sub interface
	SetInstance(i MySubInterface)
	GetInstance() MySubInterface
}

// MySubInterface is another test interface
type MySubInterface interface {
	GetName() string
	SetName(name string)
}

// Foo implements both interfaces
type Foo struct {
	name     string
	instance MySubInterface
}

// SetName sets the name
func (f *Foo) SetName(n string) {
	f.name = n
}

// GetName returns the name
func (f *Foo) GetName() string {
	return f.name
}

// SetInstance sets the instance
func (f *Foo) SetInstance(i MySubInterface) {
	f.instance = i
}

// GetInstance returns the instance
func (f *Foo) GetInstance() MySubInterface {
	if f.instance == nil {
		// what exactly does this do?
		// why can I not used this in the test?
		f.instance = f
	}
	return f.instance
}

// Bar implements only the sub interface
type Bar struct {
	name string
}

// SetName sets the name
func (b *Bar) SetName(n string) {
	b.name = n
}

// GetName returns the name
func (b *Bar) GetName() string {
	return b.name
}
