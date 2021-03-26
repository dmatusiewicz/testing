package bridger

import "fmt"

type mac struct {
	printer printer
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

func (m mac) print() {
	fmt.Println("Printing on mac")
	m.printer.print()
}
