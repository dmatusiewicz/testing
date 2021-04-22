package facader

import "fmt"

func Run() {
	commandCenter := New("Main", 400, 150, point{0, 0})
	commandCenter.AddResources(500)
	commandCenter.AddResources(-100)
	fmt.Printf("Resources: %d\n", commandCenter.crystalStorage.getResourceCount())
	marine, err := commandCenter.ProduceUnit(produceMarine)
	if err != nil {
		fmt.Printf("You should not use this marine.\n")
	}
	fmt.Printf("Marine HP %d\n", marine.GetCost())
	fmt.Printf("Resources: %d\n", commandCenter.crystalStorage.getResourceCount())
	scv, err := commandCenter.ProduceUnit(produceSCV)
	if err != nil {
		fmt.Printf("You should not use this SCV.\n")
	}
	fmt.Printf("SCV HP %d\n", scv.GetCost())
	fmt.Printf("Resources: %d\n", commandCenter.crystalStorage.getResourceCount())
	commandCenter.AddResources(6000)

}
