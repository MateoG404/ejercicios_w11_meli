package internal

// Person is the model for the person entity
type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}
