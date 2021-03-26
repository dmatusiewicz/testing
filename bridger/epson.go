package bridger

import "fmt"

type epson struct {
}

func (e epson) print() {
	fmt.Println("Epson is printing")
}
