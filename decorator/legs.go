package decorator

import "fmt"

type legs struct {
	robot robot
}

func (lc legs) Move() {
	fmt.Printf("Moving robot %s to different postion\n", lc.robot.reportName())
}
