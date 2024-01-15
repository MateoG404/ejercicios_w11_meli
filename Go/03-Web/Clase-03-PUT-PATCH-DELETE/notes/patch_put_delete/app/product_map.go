// This is the squeleton of the product map ( Simulation of a database)

package app

// Interface of the product map
type ProductMap interface {
	// Create new ProductMap
	NewProductMap() *ProductMap

	// Create new Product
	Create(id int, name string, description string, price float64) error

	// Get product by ID
	Get(id int) (Product, error)

	//UpdateorCreate product

	UpdateOrCreate(product Product) (Product, error)
}
