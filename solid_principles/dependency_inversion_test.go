package solid_test

import (
	"testing"

	solid "github.com/danitrod/go-design-patterns/solid_principles"
	"github.com/stretchr/testify/assert"
)

func TestDIP(t *testing.T) {
	t.Run("Research", func(t *testing.T) {
		parent := solid.Person{"John"}
		child := solid.Person{"Chris"}

		relationships := solid.Relationships{}
		relationships.AddParentAndChild(&parent, &child)

		r := solid.Research{relationships}
		res := r.Investigate()

		assert.Contains(t, res, "John")
		assert.Contains(t, res, "Chris")
	})

	t.Run("Better research", func(t *testing.T) {
		parent := solid.Person{"John"}
		child := solid.Person{"Chris"}

		relationships := solid.Relationships{}
		relationships.AddParentAndChild(&parent, &child)

		r := solid.BetterResearch{&relationships}
		res := r.Investigate()

		assert.Contains(t, res, "John")
		assert.Contains(t, res, "Chris")
	})
}
