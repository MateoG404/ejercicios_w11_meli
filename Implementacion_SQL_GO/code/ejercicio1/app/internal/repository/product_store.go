package repository

import (
	"database/sql"
	"ejercicio1/app/internal"
	"errors"
)

var (
	// ErrRepositoryProductNotFound is returned when a product is not found.
	ErrServiceProductNotFound = errors.New("repository: product not found")
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
func (r *ProductStore) FindById(id int) (p internal.Product, err error) {
	// Verify if the database is opened
	if err = r.db.Ping(); err != nil {
		return p, err
	}

	// Use the package database/sql to query the database
	row := r.db.QueryRow("SELECT p.id, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price FROM products AS p WHERE p.id = ?", id)

	// Manage the error
	if err := row.Err(); err != nil {
		return p, err
	}
	// Scan the row into the product
	// Initialize the product
	p = internal.Product{}

	err = row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price)

	// Manage the error when the scan fails
	if err != nil {
		switch {
		// If the SQL didn't find any row, return the error ErrRepositoryProductNotFound
		case errors.Is(err, sql.ErrNoRows):
			return internal.Product{}, internal.ErrRepositoryProductNotFound
		default:
			// If the error is different, return the error
			return internal.Product{}, err
		}
	}
	return p, nil
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
