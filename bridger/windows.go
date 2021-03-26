package bridger

import "fmt"

type windows struct {
	printer printer
}

func (w windows) print() {
	fmt.Println("Print on windwos")
	w.printer.print()
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}
