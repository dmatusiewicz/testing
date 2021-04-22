package facader

import "fmt"

const (
	produceSCV    productionIdentifier = "SCV"
	produceMarine productionIdentifier = "Marine"
)

type productionIdentifier string
type CommandCenter struct {
	unitBacklog    *unitBacklog
	rallyPoint     *rallyPoint
	crystalStorage *crystalStorage
}

type Unit interface {
	GetCost() int
}
type Marine struct{}

func (m Marine) GetCost() int {
	return 500
}

type SCV struct{}

func (m SCV) GetCost() int {
	return 600
}

func New(name string, crystal, gas int, p point) *CommandCenter {
	return &CommandCenter{
		unitBacklog:    &unitBacklog{},
		rallyPoint:     &rallyPoint{point: p},
		crystalStorage: &crystalStorage{max: 5000, current: crystal},
	}
}

func (cc CommandCenter) AddResources(r int) error {
	err := cc.crystalStorage.addResource(r)
	if err != nil {
		// fmt.Printf("Can't store additional: %d crystals\n", r)
		fmt.Printf("Resource operation failed with: %s\n", err.Error())
		return err
	}
	return nil
}

func (cc CommandCenter) ProduceUnit(u productionIdentifier) (Unit, error) {
	var unit Unit

	if u == "SCV" {
		unit = SCV{}
		fmt.Println("Prodducing SCV")
	} else if u == "Marine" {
		unit = Marine{}
		fmt.Println("Prodducing Marine")
	}
	err := cc.AddResources(-unit.GetCost())
	if err != nil {
		return unit, err
	}
	return unit, nil
}
