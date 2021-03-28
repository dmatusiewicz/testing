package decorator

import "fmt"

func Run() {

	robot := r2d2{Name: "234-2349823-f"}

	detector := detector{robot: robot}
	detector.Detect()

	legs := legs{robot: robot}
	legs.Move()

	fmt.Printf("Robot is reporting name: %s\n", robot.reportName())
}
