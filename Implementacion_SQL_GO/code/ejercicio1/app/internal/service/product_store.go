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

// Method to validate if a product exists
func (s *ProductService) ValidateProductExists(id int) (productExists bool) {

	// Use the repository to find the product by id
	_, err := s.repository.FindById(id)
	if err != nil {
		return errors.Is(err, repository.ErrServiceProductNotFound)
	}
	return true
}

// Save implements internal.RepositoryProduct to save a product.
func (s *ProductService) Save(p *internal.Product) (err error) {
	// Bussiness logic to save a product

	// - Validate if some product with the same id exists
	productExists := s.ValidateProductExists(p.Id)

	// If exists, change the id to the last id + 1
	if productExists {
		lastID, err := s.repository.GetLastID()
		if err != nil {
			return err
		}
		p.Id = lastID + 1
	}

	// Save the product
	err = s.repository.Save(p)
	if err != nil {
		return err
	}
	return nil
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
