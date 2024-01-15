// This script contains the struct of the product and the methods to interact with the database

package repository

import (
	"fmt"
	internal "patch_put_delete/app"
)

// Var Errors

var (
	ErrProductNotFound    = fmt.Errorf("product not found")
	ErrInvalidID          = fmt.Errorf("invalid id")
	ErrProductExists      = fmt.Errorf("product already exists")
	ErrInvalidName        = fmt.Errorf("invalid name")
	ErrInvalidDescription = fmt.Errorf("invalid description")

	ErrInvalidPrice = fmt.Errorf("invalid price")
)

// This is the struct of the product
type ProductMap struct {
	db     map[int]internal.Product
	lastId int
}

// Methdos to interact with the database

// Create new ProductMap

func NewProductMap() *ProductMap {
	return &ProductMap{
		db:     make(map[int]internal.Product),
		lastId: 0,
	}
}

// Create new Product
func (p *ProductMap) Create(id int, name string, description string, price float64) (internal.Product, error) {

	// Verify if the product exists
	if p.lastId+1 < id {
		return internal.Product{}, ErrInvalidID
	}

	// Create the product
	product := internal.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}

	// Add the product to the database
	p.db[id] = product
	p.lastId++
	return p.db[id], nil

}

// Get a product by id
func (p *ProductMap) GetProductById(id int) (internal.Product, error) {

	product, ok := p.db[id]
	if !ok {
		return internal.Product{}, ErrProductNotFound
	}
	return product, nil

}

// Update or create a product

func (p *ProductMap) UpdateOrCreate(product internal.Product) (internal.Product, error) {
	// Verify if the product exists
	if _, ok := p.db[product.ID]; !ok {
		return internal.Product{}, ErrProductExists
	}

	// Update the product
	p.db[product.ID] = product
	// Return the product

	return p.db[product.ID], nil
}
