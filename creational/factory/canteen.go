package factory

func SimpleFactoryExample() {
	soup := SimpleSoupCreator(Egusi)
	soup.Prepare()

	soup2 := SimpleSoupCreator(Efo)
	soup2.Prepare()
}