package test

/* Use case:
Two objects types (interfaces), where one uses the other.
Both shall be open for a specific implementation and defined as interfaces.
*/
type MyInterface interface {
	SetName(name string)
	// CANNOT be pointer!
	// link to an object complying to the sub interface
	SetInstance(i MySubInterface)
	GetInstance() MySubInterface
}

type MySubInterface interface {
	GetName() string
	SetName(name string)
}

// Foo implements both interfaces
type Foo struct {
	name     string
	instance *MySubInterface
}

func (f *Foo) SetName(n string) {
	f.name = n
}
func (f *Foo) GetName() string {
	return f.name
}
func (f *Foo) SetInstance(i MySubInterface) {
	f.instance = &i
}
func (f *Foo) GetInstance() MySubInterface {
	if f.instance == nil {
		// what exactly does this do?
		// why can I not used this in the test?
		var m MySubInterface = f
		f.instance = &m
	}
	return *f.instance
}

// Bar implements only the sub interface
type Bar struct {
	name string
}

func (b *Bar) SetName(n string) {
	b.name = n
}
func (b *Bar) GetName() string {
	return b.name
}
