package service

import (
	internal "ejercicio1/app/internal"
	repository "ejercicio1/app/internal/repository"
	"errors"
)

// Struct for the service of the product
type ProductService struct {
	repository repository.ProductStore
}

// Method to create a new service of the product

func NewProductService(repository repository.ProductStore) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

// Delete implements internal.RepositoryProduct.
func (s *ProductService) Delete(id int) (err error) {
	panic("unimplemented")
}

// FindById implements internal.RepositoryProduct.
func (s *ProductService) FindById(id int) (p internal.Product, err error) {
	// Call the method of the repository
	product, err := s.repository.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrServiceProductNotFound):
			return internal.Product{}, repository.ErrServiceProductNotFound
		default:
		}
		return
	}
	return product, nil
}

// Save implements internal.RepositoryProduct.
func (*ProductService) Save(p *internal.Product) (err error) {
	panic("unimplemented")
}

// Update implements internal.RepositoryProduct.
func (*ProductService) Update(p *internal.Product) (err error) {
	panic("unimplemented")
}

/*
// UpdateOrSave implements internal.RepositoryProduct.
func (*ProductService) UpdateOrSave(p *repository.Product) (err error) {
	panic("unimplemented")
}

*/
