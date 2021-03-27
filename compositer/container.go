package compositer

type Container struct {
	assets []asset
	value  int
	name   string
}

func (c *Container) calculateValue() int {
	var calculatedValue int
	calculatedValue += c.value
	for _, j := range c.assets {
		calculatedValue += j.calculateValue()
	}
	return calculatedValue
}

func (c *Container) add(i asset) {
	c.assets = append(c.assets, i)
}

func (c *Container) showAllAssets() []asset {
	var assets []asset
	for _, j := range c.assets {
		for _, jj := range j.showAllAssets() {
			assets = append(assets, jj)
		}
	}
	assets = append(assets, c)
	return assets
}

func (c *Container) show() (name string, value int) {
	return c.name, c.value
}
