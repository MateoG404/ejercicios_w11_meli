// This script contains the code for the service to use the database and pass the data to the handler

package service

import (
	"errors"
	"fmt"
	internal "patch_put_delete/app"
	repository "patch_put_delete/app/internal/repository"
)

// Struct Service

type ProductService struct {
	db repository.ProductMap
}

// Create a new service
func NewService(db repository.ProductMap) *ProductService {
	return &ProductService{
		db: db,
	}
}

// Service methods
// Method to validate the input data
func (p *ProductService) ValidateInputData(id int, name string, description string, price float64) error {

	// Verify if the id is valid
	if id < 0 {
		return repository.ErrInvalidID
	}

	// Verify if the name is valid
	if name == "" {
		return repository.ErrInvalidName
	}

	// Verify if the description is valid
	if description == "" {
		return repository.ErrInvalidDescription
	}

	// Verify if the price is valid
	if price < 0 {
		return repository.ErrInvalidPrice
	}

	// Return nil if the input data is valid
	return nil
}

// Create a new product
func (p *ProductService) CreateProduct(id int, name string, description string, price float64) (internal.Product, error) {

	// Bussiness logic

	// -Validate the input data

	if ok := p.ValidateInputData(id, name, description, price); ok != nil {
		return internal.Product{}, ok
	}
	// Verify if the product exists

	_, err := p.db.GetProductById(id)
	// If the product exists, return an error
	if !errors.Is(err, repository.ErrProductNotFound) {
		return internal.Product{}, repository.ErrProductExists
	}

	// Create the product
	product, err := p.db.Create(id, name, description, price)
	if err != nil {
		return internal.Product{}, err
	}
	fmt.Println("Aca 3")
	// Return the product for the body response
	return product, nil
}

// Get a product by id
