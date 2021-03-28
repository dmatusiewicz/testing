package decorator

import "fmt"

type r2d2 struct {
	Name string
}

func (r r2d2) reportName() string {
	return r.Name
}

func (r r2d2) sendData() {
	fmt.Printf("%s is sending data\n", r.Name)
}

func (r r2d2) gatherData() {
	fmt.Printf("%s is gathering data\n", r.Name)
}
