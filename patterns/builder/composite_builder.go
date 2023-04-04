package builder

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (b *PersonBuilder) Lives() *AddressBuilder {
	return &AddressBuilder{*b}
}

func (b *PersonBuilder) Works() *JobBuilder {
	return &JobBuilder{*b}
}

type AddressBuilder struct {
	PersonBuilder
}

func (b *AddressBuilder) At(streetAddress string) *AddressBuilder {
	b.person.Address.StreetAddress = streetAddress
	return b
}

func (b *AddressBuilder) In(city string) *AddressBuilder {
	b.person.Address.City = city
	return b
}

func (b *AddressBuilder) WithPostCode(postCode string) *AddressBuilder {
	b.person.Address.PostCode = postCode
	return b
}

type JobBuilder struct {
	PersonBuilder
}

func (b *JobBuilder) At(company string) *JobBuilder {
	b.person.Job.CompanyName = company
	return b
}

func (b *JobBuilder) AsA(position string) *JobBuilder {
	b.person.Job.Position = position
	return b
}

func (b *JobBuilder) Earning(income int) *JobBuilder {
	b.person.Job.Income = income
	return b
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}
