package main

import (
	"fmt"

	"github.com/dmatusiewicz/testing/adapter"
	"github.com/dmatusiewicz/testing/bridger"
	"github.com/dmatusiewicz/testing/heaper"
)

func main() {
	heaper.Run()
	fmt.Println()
	adapter.Run()
	fmt.Println()
	bridger.Run()
}
