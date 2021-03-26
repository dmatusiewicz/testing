package adapter

import "fmt"

type unknownDataSourceAdapter struct {
	unknownDataSource *unknownDataSource
}

func (udsa unknownDataSourceAdapter) ReadData() {
	fmt.Println("Translatind data from unknown data soruce to known data source.")
	udsa.unknownDataSource.ReadUnknownData()
}
