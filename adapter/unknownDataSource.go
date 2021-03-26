package adapter

import "fmt"

type unknownDataSource struct {
}

func (uds unknownDataSource) ReadUnknownData() {
	fmt.Println("Reading unknown data.")
}
