package builder

import (
	"fmt"
	"strings"
)

const indentSize = 2

type HTMLElement struct {
	tag, text string
	elements  []HTMLElement
}

func (e *HTMLElement) String() string {
	return e.string(0)
}

func (e *HTMLElement) string(indent int) string {
	sb := strings.Builder{}
	i := strings.Repeat(" ", indentSize*indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.tag))

	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.elements {
		sb.WriteString(el.string(indent + 1))
	}

	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.tag))
	return sb.String()
}
