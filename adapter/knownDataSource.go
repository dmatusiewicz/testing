package adapter

import "fmt"

type knownDataSource struct {
}

func (ksd knownDataSource) ReadData() {
	fmt.Println("Reading from known data source.")
}
