package solid_test

import (
	"testing"

	"github.com/danitrod/go-design-patterns/solid"
	"github.com/stretchr/testify/assert"
)

func TestShouldFilterProductsByColor(t *testing.T) {
	apple := solid.Product{"Apple", solid.Red, solid.Small}
	tree := solid.Product{"Tree", solid.Green, solid.Medium}
	house := solid.Product{"House", solid.Blue, solid.Large}

	products := []solid.Product{apple, tree, house}
	f := solid.Filter{}

	filteredProducts := f.FilterByColor(products, solid.Blue)

	assert.Len(t, filteredProducts, 1)
	assert.Equal(t, "House", filteredProducts[0].Name)
}

func TestShouldFilterProductsUsingSpecification(t *testing.T) {
	apple := solid.Product{"Apple", solid.Red, solid.Small}
	tree := solid.Product{"Tree", solid.Green, solid.Medium}
	house := solid.Product{"House", solid.Blue, solid.Large}

	products := []solid.Product{apple, tree, house}
	greenSpecification := solid.ColorSpecification{solid.Green}
	bf := solid.BetterFilter{}

	filteredProducts := bf.Filter(products, greenSpecification)

	assert.Len(t, filteredProducts, 1)
	assert.Equal(t, "Tree", filteredProducts[0].Name)
}

func TestShouldFilterWithAndSpecification(t *testing.T) {
	apple := solid.Product{"Apple", solid.Red, solid.Small}
	tree := solid.Product{"Tree", solid.Green, solid.Medium}
	house := solid.Product{"House", solid.Blue, solid.Large}

	products := []solid.Product{apple, tree, house}

	smallSpec := solid.SizeSpecification{solid.Small}
	redSpec := solid.ColorSpecification{solid.Red}
	andSpec := solid.AndSpecification{smallSpec, redSpec}
	bf := solid.BetterFilter{}

	filteredProducts := bf.Filter(products, andSpec)

	assert.Len(t, filteredProducts, 1)
	assert.Equal(t, "Apple", filteredProducts[0].Name)
}
