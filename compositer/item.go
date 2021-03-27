package compositer

type Item struct {
	Name  string
	Value int
}

func (i Item) calculateValue() int {
	return i.Value
}
