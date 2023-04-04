package builder

type personMod func(*Person)
type PersonBuilderFunctional struct {
	actions []personMod
}

func (b *PersonBuilderFunctional) Called(name string) *PersonBuilderFunctional {
	b.actions = append(b.actions, func(p *Person) {
		p.Name = name
	})
	return b
}

func (b *PersonBuilderFunctional) WorksAsA(position string) *PersonBuilderFunctional {
	b.actions = append(b.actions, func(p *Person) {
		p.Job.Position = position
	})
	return b
}

func (b *PersonBuilderFunctional) Build() *Person {
	p := Person{}
	for _, a := range b.actions {
		a(&p)
	}

	return &p
}
