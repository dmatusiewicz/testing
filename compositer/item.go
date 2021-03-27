package compositer

type Item struct {
	name  string
	value int
}

func (i Item) calculateValue() int {
	return i.value
}

func (i Item) show() (name string, value int) {
	return i.name, i.value
}

func (i Item) showAllAssets() []asset {
	var assets []asset
	// fmt.Printf("Adding to empty list %v", i)
	return append(assets, i)
}
