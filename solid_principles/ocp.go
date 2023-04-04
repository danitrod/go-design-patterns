package solid

// Open Closed Principal
// Types should be open for extension, closed for modification

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(
	products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

// If we were to add a filter for other property, we would have to duplicate this method

// Instead, we can make an extensible type, using the Specification interface
type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	Color Color
}

func (c ColorSpecification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)
	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}
	return result
}

// Now the filter is closed for modification, but extensible (you can use other specifications
// as input)

// Composite specification
type AndSpecification struct {
	First, Second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.First.IsSatisfied(p) && a.Second.IsSatisfied(p)
}
