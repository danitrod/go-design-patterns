package solid_test

import (
	"testing"

	"github.com/danitrod/go-design-patterns/solid"
	"github.com/stretchr/testify/assert"
)

func TestJournal(t *testing.T) {
	t.Run("should save and load from file using persistence", func(t *testing.T) {
		filename := "test"
		j := solid.Journal{}

		j.AddEntry("test")
		j.AddEntry("test2")

		p := solid.Persistence{LineSeparator: "\n"}
		p.SaveToFile(&j, filename)
		loaded, err := p.LoadFromFile(filename)

		assert.Nil(t, err)
		assert.Equal(t, j.String(), loaded.String())
	})
}
