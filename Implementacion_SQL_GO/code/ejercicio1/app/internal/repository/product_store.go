package repository

import (
	"database/sql"
	"ejercicio1/app/internal"
)

// Struct for the repository of the product

type ProductStore struct {
	db *sql.DB
}

// Delete implements internal.RepositoryProduct.
func (*ProductStore) Delete(id int) (err error) {
	panic("unimplemented")
}

// FindById implements internal.RepositoryProduct.
func (*ProductStore) FindById(id int) (p internal.Product, err error) {
	panic("unimplemented")
}

// Save implements internal.RepositoryProduct.
func (*ProductStore) Save(p *internal.Product) (err error) {
	panic("unimplemented")
}

// Update implements internal.RepositoryProduct.
func (*ProductStore) Update(p *internal.Product) (err error) {
	panic("unimplemented")
}

// UpdateOrSave implements internal.RepositoryProduct.
func (*ProductStore) UpdateOrSave(p *internal.Product) (err error) {
	panic("unimplemented")
}

// Method to create a new repository of the product
func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}
