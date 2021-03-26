package adapter

func Run() {
	c := &client{}
	kds := &knownDataSource{}

	c.ReadSomeDataFromSource(kds)

	uds := &unknownDataSource{}
	udsa := &unknownDataSourceAdapter{
		unknownDataSource: uds,
	}

	c.ReadSomeDataFromSource(udsa)

}
