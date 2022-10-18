package solid_test

import (
	"testing"

	"github.com/danitrod/go-design-patterns/solid"
	"github.com/stretchr/testify/assert"
)

func TestGetSetRectProps(t *testing.T) {
	rect := solid.Rectangle{2, 3}

	w := rect.GetWidth()
	rect.SetHeight(10)

	assert.Equal(t, 20, w*rect.GetHeight())
}

func TestGetSetUnsafeSquareProps(t *testing.T) {
	sq := solid.NewSquare(2)

	w := sq.GetWidth()
	sq.SetWidth(10)

	assert.Equal(t, 20, w*sq.GetHeight())
}
