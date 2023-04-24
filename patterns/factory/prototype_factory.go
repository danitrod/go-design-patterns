package factory

const (
	Developer = iota
	Manager
)

func NewEmployee(name string, role int) *Employee {
	switch role {
	case Developer:
		return &Employee{name, "developer", 1000}
	case Manager:
		return &Employee{name, "manager", 2000}
	default:
		return &Employee{name, "unknown", 0}
	}
}
