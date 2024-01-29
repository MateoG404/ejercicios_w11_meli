package service

import (
	repository "ejercicio1/app/internal"
)

// Struct for the service of the product
type ProductService struct {
	repository repository.RepositoryProduct
}

// Method to create a new service of the product

func NewProductService(repository repository.RepositoryProduct) *ProductService {
	return &ProductService{
		repository: repository,
	}
}
