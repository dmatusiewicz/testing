package bridger

import "fmt"

func Run() {

	hp := &hp{}
	ep := &epson{}
	mac := &mac{}
	win := &windows{}
	mac.setPrinter(hp)
	win.setPrinter(ep)

	mac.print()
	win.print()

	mac.setPrinter(ep)
	win.setPrinter(hp)

	fmt.Println()
	mac.print()
	win.print()
}
