package builder_test

import (
	"strings"
	"testing"

	"github.com/danitrod/go-design-patterns/patterns/builder"
	"github.com/stretchr/testify/assert"
)

func TestGoStringsBuilder(t *testing.T) {
	t.Run("Should use Go strings builder to build an HTML paragraph", func(t *testing.T) {
		text := "hello"
		sb := strings.Builder{}
		sb.WriteString("<p>")
		sb.WriteString(text)
		sb.WriteString("</p>")

		assert.Equal(t, "<p>hello</p>", sb.String())
	})

	t.Run("Should use Go strings builder to build an HTML list", func(t *testing.T) {
		words := []string{"hello", "world"}
		sb := strings.Builder{}
		sb.WriteString("<ul>")
		for _, w := range words {
			sb.WriteString("<li>")
			sb.WriteString(w)
			sb.WriteString("</li>")
		}
		sb.WriteString("</ul>")

		assert.Equal(t, "<ul><li>hello</li><li>world</li></ul>", sb.String())
	})
}

func TestHTMLBuilder(t *testing.T) {
	t.Run("Should use HTML builder to build a stringified HTML list", func(t *testing.T) {
		b := builder.NewHTMLBuilder("ul")
		b.AddChild("li", "hello")
		b.AddChild("li", "world")

		assert.Equal(t, `<ul>
  <li>
    hello
  </li>
  <li>
    world
  </li>
</ul>
`, b.String())
	})

	t.Run("Should use HTML builder with fluent interface to build a stringified HTML list", func(t *testing.T) {
		result := builder.NewHTMLBuilder("ul").
			AddChildFluent("li", "hello").
			AddChildFluent("li", "world").
			String()

		assert.Equal(t, `<ul>
  <li>
    hello
  </li>
  <li>
    world
  </li>
</ul>
`, result)
	})
}

func TestPersonBuilder(t *testing.T) {
	t.Run("Should use aggregate builders to build a person", func(t *testing.T) {
		pb := builder.NewPersonBuilder()
		pb.
			Lives().
			At("123 London Road").
			In("London").
			WithPostCode("123M4C").
			Works().
			At("Fabrikam").
			AsA("Programmer").
			Earning(123000)

		person := pb.Build()

		assert.Equal(t, "123 London Road", person.Address.StreetAddress)
		assert.Equal(t, "London", person.Address.City)
		assert.Equal(t, "123M4C", person.Address.PostCode)
		assert.Equal(t, "Fabrikam", person.Job.CompanyName)
		assert.Equal(t, "Programmer", person.Job.Position)
		assert.Equal(t, 123000, person.Job.Income)
	})
}

func TestBuilderParameter(t *testing.T) {
	t.Run("Should send an e-mail without having access to the Email struct", func(t *testing.T) {
		err := builder.SendEmail(func(b *builder.EmailBuilder) {
			b.From("foo@example.com").
				To("bar@example.com").
				Subject("Hello").
				Body("Testing...")
		})

		assert.Nil(t, err)
	})
}

func TestFunctionalBuilder(t *testing.T) {
	t.Run("Should use functional builder to create a person", func(t *testing.T) {
		b := builder.PersonBuilderFunctional{}
		p := b.Called("Dimitri").WorksAsA("Developer").Build()

		assert.Equal(t, "Dimitri", p.Name)
		assert.Equal(t, "Developer", p.Job.Position)
	})
}
