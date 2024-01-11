// This code contains the product model
package internal

// Product is the model of a product that contains the information of a product
type Product struct {
	Id       int
	Name     string
	Type     string
	Quantity int
	Price    float64
}
