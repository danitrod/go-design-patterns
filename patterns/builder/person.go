package builder

type Person struct {
	Name    string
	Address address
	Job     job
}

type address struct {
	StreetAddress, PostCode, City string
}

type job struct {
	CompanyName, Position string
	Income                int
}
