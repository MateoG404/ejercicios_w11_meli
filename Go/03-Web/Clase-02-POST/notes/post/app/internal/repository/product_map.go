package repository

import "post/app/internal"

// This code containes the implementation of the default repository of the product model

// Struct to use in the storage of the product model
type ProductMap struct {
	// Map to store the products
	db map[int]internal.Product

	// LastId to generate the id of the product
	lastId int
}

// Method to create a new product map
func NewProductMap() *ProductMap {
	return &ProductMap{
		db:     make(map[int]internal.Product, 0),
		lastId: 0,
	}
}

// Method to save a product in the storage

func (p *ProductMap) Save(product *internal.Product) (err error) {

	// Verify if the product has an id
	if product.Id == 0 {
		// If not, generate the id
		p.lastId++
		product.Id = p.lastId
	}

	// Save the product in the storage
	p.db[product.Id] = *product

	return
}
