package factory

type Employee struct {
	Name, Position string
	Income         int
}

// Functional
// With this generator function, we can create a factory of employees from a specific position and
// income
func NewEmployeeFactory(position string, income int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{
			Name:     name,
			Position: position,
			Income:   income,
		}
	}
}

// Structural
type EmployeeFactory struct {
	Position string
	Income   int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Income}
}

func NewEmployeeFactoryStructural(position string, income int) *EmployeeFactory {
	return &EmployeeFactory{position, income}
}
