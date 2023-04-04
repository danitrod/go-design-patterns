package solid

// An interface that works with a base class should also work with a derived class.

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	sq := Square{}
	sq.Width = size
	sq.Height = size
	return &sq
}

func (s *Square) SetWidth(width int) {
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height int) {
	s.Width = height
	s.Height = height
}

// These setters above are dangerous: see the second test in the test file.

// There are many ways to avoid this problem, one of them is not having squares at all, just using
// rects with same sized sides.
// Or do something like below, enabling the square to transform into a regular rectangle.
type SafeSquare struct {
	size int // size == width == height
}

func (s *SafeSquare) ToRect() Rectangle {
	return Rectangle{s.size, s.size}
}
