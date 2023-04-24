package factory

import "fmt"

type Greeter interface {
	SayHello() string
}

type Person struct {
	Name          string
	Age, EyeCount int
}

func (p *Person) SayHello() string {
	return fmt.Sprintf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

type tiredPerson struct {
	name          string
	age, eyeCount int
}

func (p *tiredPerson) SayHello() string {
	return "Sorry, I'm tired"
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age, 2} // Eye count will always be 2, so we can set it from the factory
}

func NewGreeter(name string, age int) Greeter {
	// We can return different implementations of a same interface from a factory
	if age > 70 {
		return &tiredPerson{name, age, 2}
	}

	return &Person{name, age, 2}
}

func (p *Person) removeEye() {
	if p.EyeCount > 0 {
		p.EyeCount -= 1
	}
}
