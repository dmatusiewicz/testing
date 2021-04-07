package facader

type storage interface {
	addResource(int) error
	getResourceCount() int
}
