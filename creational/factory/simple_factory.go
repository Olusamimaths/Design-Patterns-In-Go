package factory

const Egusi = "Egusi"
const Efo = "Efo"
const Tomato = "Tomato"

type ISoup interface {
	Type() string
	Prepare()
}

func SimpleSoupCreator(soupType string) ISoup {
	switch soupType {
	case Egusi:
		return &EgusiSoup{}
	case Efo:
		return &EfoSoup{}
	case Tomato:
		return &TomatoSoup{}
	default:
		return nil
	}
}

