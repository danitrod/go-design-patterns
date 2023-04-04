package builder

type HTMLBuilder struct {
	rootTag string
	root    HTMLElement
}

func NewHTMLBuilder(rootTag string) *HTMLBuilder {
	return &HTMLBuilder{rootTag,
		HTMLElement{rootTag, "", []HTMLElement{}}}
}

func (b *HTMLBuilder) String() string {
	return b.root.String()
}

func (b *HTMLBuilder) AddChild(tag, text string) {
	e := HTMLElement{tag, text, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
}

func (b *HTMLBuilder) AddChildFluent(tag, text string) *HTMLBuilder {
	e := HTMLElement{tag, text, []HTMLElement{}}
	b.root.elements = append(b.root.elements, e)
	return b
}
