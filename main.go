package main

import (
	"fmt"

	"github.com/dmatusiewicz/testing/adapter"
	"github.com/dmatusiewicz/testing/bridger"
	"github.com/dmatusiewicz/testing/compositer"
	"github.com/dmatusiewicz/testing/decorator"
	"github.com/dmatusiewicz/testing/facader"
	"github.com/dmatusiewicz/testing/heaper"
)

func main() {
	heaper.Run()
	fmt.Println()
	adapter.Run()
	fmt.Println()
	bridger.Run()
	fmt.Println()
	compositer.Run()
	fmt.Println()
	decorator.Run()
	fmt.Println()
	facader.Run()

}
