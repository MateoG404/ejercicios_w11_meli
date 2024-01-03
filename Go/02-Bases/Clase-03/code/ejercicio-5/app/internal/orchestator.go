package internal

// Create a type funcType that is a func() string
type funcType func() string

const (
	dog = "dog"
	cat = "cat"

	hamster   = "hamster"
	tarantula = "tarantula"
)

func Orchestator(animal string) funcType {

	switch animal {
	case dog:
		return AnimalDog
	case cat:
		return AnimalCat
	case hamster:
		return AnimalHamster
	case tarantula:
		return AnimalTarantula
	default:
		return func() string {
			return "Animal no encontrado"
		}

	}
}
