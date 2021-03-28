package decorator

type robot interface {
	sendData()
	gatherData()
	reportName() string
}
