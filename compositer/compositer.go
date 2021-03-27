package compositer

import "fmt"

func Run() {
	charger := Item{name: "Charger", value: 15}
	phone := Item{name: "Iphone11", value: 800}
	chargerContainer := Container{name: "Charger box", value: 1}
	chargerContainer.add(charger)
	phoneContainer := Container{name: "Phone box", value: 1}
	phoneContainer.add(phone)
	box := Container{name: "Box above all"}
	box.add(&chargerContainer)
	box.add(&phoneContainer)
	box.add(&Item{name: "Invoice"})
	fmt.Printf("Box value %d$\n", box.calculateValue())

	for _, j := range box.showAllAssets() {
		s, v := j.show()
		fmt.Printf("- \"%s\" of value %d$\n", s, v)
	}
}
