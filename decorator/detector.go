package decorator

import "fmt"

type detector struct {
	robot robot
}

func (d detector) Detect() {
	fmt.Printf("Robot %s is detecting...\n", d.robot.reportName())
}
