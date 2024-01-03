package internal

// Animal is a struct that represents an animal
type Animal struct {
	// Name is the name of the animal
	Name string `json:"name"`
	// Age is the age of the animal
	Age int `json:"age"`
	// Race is the race of the animal
	Race string `json:"race"`
}

// Method to print the name
