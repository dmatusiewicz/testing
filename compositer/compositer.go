package compositer

import "fmt"

func Run() {
	charger := Item{Name: "Charger", Value: 15}
	phone := Item{Name: "Iphone11", Value: 800}
	chargerContainer := &Container{}
	chargerContainer.addItem(charger)
	phoneContainer := &Container{}
	phoneContainer.addItem(phone)
	box := &Container{}
	box.addContainer(*chargerContainer)
	box.addContainer(*phoneContainer)
	box.addItem(Item{Name: "Invoice"})
	box.addItem(Item{Name: "Headphones", Value: 400})

	fmt.Printf("Box value %d$\n", box.calculateValue())

	for _, j := range box.getContent() {
		fmt.Printf("- %s, value: %d$\n", j.Name, j.Value)
	}
}
