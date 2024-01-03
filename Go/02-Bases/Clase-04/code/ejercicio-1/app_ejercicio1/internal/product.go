package internal

import "fmt"

// Product struct
var (
	SliceProducts []Product
)

type Product struct {

	// Variables

	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

// Method Save

// This method save the product in the sliceProduct
func (p Product) Save() {
	SliceProducts = append(SliceProducts, p)
}

// Method GetAll
// This method return all the products in the sliceProduct

func GetAll() {
	for pos, item := range SliceProducts {
		fmt.Println("La posicion ", pos, "tiene el item ", item)
	}
}

// Function GetByID
// This function return a product by ID

func GetByID(id int) Product {
	return SliceProducts[id]
}
