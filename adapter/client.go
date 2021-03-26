package adapter

import "fmt"

type client struct {
}

func (c client) ReadSomeDataFromSource(s StringDataSource) {
	fmt.Println("The client is about to read some data from source")
	s.ReadData()
}
