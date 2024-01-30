// This script contains the implementation of the product service

package internal

// Struct for the service of the product
type ProductService interface {
	// Delete implements internal.RepositoryProduct.
	Delete(id int) (err error)
	// FindById implements internal.RepositoryProduct.
	FindById(id int) (p Product, err error)
	// Save implements internal.RepositoryProduct.
	Save(p *Product) (err error)
	// Update implements internal.RepositoryProduct.
	Update(p *Product) (err error)
	// UpdateOrSave implements internal.RepositoryProduct.
	UpdateOrSave(p *Product) (err error)
	// ReadAll reads all products from the store.
	ReadAll() (p map[int]Product, err error)
	// WriteAll writes all products to the store.
	WriteAll(p map[int]Product) (err error)
}
