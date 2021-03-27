package compositer

type Container struct {
	items      []Item
	containers []Container
	value      int
}

func (c Container) calculateValue() int {
	c.value = 0
	for _, j := range c.items {
		c.value += j.calculateValue()
	}
	for _, j := range c.containers {
		c.value += j.calculateValue()
	}
	return c.value
}

func (c *Container) addItem(i Item) {
	c.items = append(c.items, i)
}

func (c *Container) addContainer(cont Container) {
	c.containers = append(c.containers, cont)
}

func (c *Container) getContent() []Item {
	var content []Item

	for _, j := range c.items {
		content = append(content, j)
	}

	for _, j := range c.containers {
		for _, jj := range j.getContent() {
			content = append(content, jj)
		}
	}

	return content
}
