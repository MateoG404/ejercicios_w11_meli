package internal

// Person is a struct that represents a person
type Person struct {
	// Name is the name of the person
	FirstNmame string `json:"first_name"`
	// LastName is the last name of the person
	LastName string `json:"last_name"`
	// CellPhone is the cell phone of the person
	CellPhone string `json:"cell_phone"`
	// Address is the address of the person
	Address string `json:"address,omitempty"`
	// Password is the password of the person
	Password string `json:"-"`
}
