package solid

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct{}

func (m MultiFunctionPrinter) Print(d Document) {
	// ...
}
func (m MultiFunctionPrinter) Fax(d Document) {
	// ...
}
func (m MultiFunctionPrinter) Scan(d Document) {
	// ...
}

type OldFashionedPrinter struct {
}

// You will have to implement all the methods again

func (o OldFashionedPrinter) Print(d Document) {
	// ...
}
func (o OldFashionedPrinter) Fax(d Document) {
	// ...
}
func (o OldFashionedPrinter) Scan(d Document) {
	// ...
	panic("operation not supported")
}

// Although the OldFashionedPrinter does not implement the scan method, for it to be a Machine
// interface it needs to have the method.

// Interface Segregation Principle
// to apply this principle, we could simply separate the Machine interface's methods into separate
// interfaces
type Scanner interface {
	Scan(d Document)
}

type Printer interface {
	Print(d Document)
}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {
	// ...
}

func (p Photocopier) Scan(d Document) {
	// ...
}

// We could use the decorator pattern to use multiple interfaces in the same struct:
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func (m MultiFunctionMachine) Scan(d Document) {
	m.scanner.Scan(d)
}
