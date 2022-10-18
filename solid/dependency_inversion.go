package solid

// Dependency Inversion Principle
// High level modules should not depend on low level modules
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
}

type Relation struct {
	From         *Person
	Relationship Relationship
	To           *Person
}

// low level module
type Relationships struct {
	Relations []Relation
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.Relations = append(r.Relations, Relation{parent, Parent, child})
	r.Relations = append(r.Relations, Relation{child, Child, parent})
}

// high level module
type Research struct {
	// this breaks dependency inversion principle
	Relationships Relationships
}

func (r *Research) Investigate() string {
	relations := r.Relationships.Relations
	for _, rel := range relations {
		if rel.From.Name == "John" && rel.Relationship == Parent {
			return "John has a child called " + rel.To.Name
		}
	}
	return "No interesting relationships found"
}

// If we were to switch database, for example, like the way to store relationships, the Investigate
// function would break.
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type BetterResearch struct {
	Browser RelationshipBrowser
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.Relations {
		if v.Relationship == Parent && v.From.Name == name {
			result = append(result, r.Relations[i].To)
		}
	}

	return result
}

func (r *BetterResearch) Investigate() string {
	out := ""
	for _, p := range r.Browser.FindAllChildrenOf("John") {
		out += "John has a child called " + p.Name + "\n"
	}

	if out == "" {
		return "No interesting relationships found"
	}
	return out
}

// Now, investigate does not depend on what database is being used... We can simply implement
// other RelationshipBrowser if needed
