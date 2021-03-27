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

func (c *Container) list() []asset {
	var assets []asset
	for _, j := range c.assets {
		switch j.(type) {
		case *Container:
			{
				x := j.(*Container).list()
				for _, jj := range x {
					assets = append(assets, jj)
				}
				assets = append(assets, j)
			}
		default:
			{
				assets = append(assets, j)
			}
		}

	}

	return assets
}

func (c *Container) show() (name string, value int) {
	return c.name, c.value
}
