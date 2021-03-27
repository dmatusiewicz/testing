package compositer

type asset interface {
	calculateValue() int
	show() (name string, value int)
	showAllAssets() []asset
}
