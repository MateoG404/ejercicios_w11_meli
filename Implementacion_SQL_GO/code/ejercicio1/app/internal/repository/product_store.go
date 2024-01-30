package repository

import (
	"database/sql"
	"ejercicio1/app/internal"
	"errors"
	"log"
)

var (
	// ErrRepositoryProductNotFound is returned when a product is not found.
	ErrServiceProductNotFound = errors.New("repository: product not found")
	ErrProductExists          = errors.New("repository: product exists")
)

// Struct for the repository of the product

type ProductStore struct {
	db *sql.DB
}

// Delete implements internal.RepositoryProduct.
func (r *ProductStore) Delete(id int) (err error) {
	// Verify if the database is opened
	if err = r.VerifyOpenDB(); err != nil {
		return err
	}

	// Use the package database/sql to query the database
	_, err = r.db.Exec("DELETE FROM products WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

// VerifyOpenDB verifies if the database is opened
func (r *ProductStore) VerifyOpenDB() (err error) {
	// Verify if the database is opened
	if err = r.db.Ping(); err != nil {
		return err
	}
	return nil
}

// FindById implements internal.RepositoryProduct.
func (r *ProductStore) FindById(id int) (p internal.Product, err error) {
	// Verify if the database is opened
	if err = r.VerifyOpenDB(); err != nil {
		return internal.Product{}, err
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
func (r *ProductStore) Save(p *internal.Product) (err error) {
	// Verify if the database is opened
	if err = r.VerifyOpenDB(); err != nil {
		return err
	}
	// Use the package database/sql to query the database
	result, err := r.db.Exec("INSERT INTO products(id, name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?, ?)", p.Id, p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)
	// Manage the error when the query fails
	if err != nil {
		return err
	}
	// Get the last inserted id (id of the product saved)
	lastId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	// Set the id of the product
	p.Id = int(lastId)
	return nil
}

// GetLastId returns the last inserted id
func (r *ProductStore) GetLastID() (lastId int, err error) {
	// Verify if the database is opened
	if err = r.VerifyOpenDB(); err != nil {
		return 0, err
	}
	// Use the package database/sql to query the database
	var lastInsertID int
	err = r.db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	return lastInsertID, nil
}

// Update implements internal.RepositoryProduct.
func (r *ProductStore) Update(p *internal.Product) (err error) {
	// Verify if the database is opened
	if err = r.VerifyOpenDB(); err != nil {
		return err
	}
	// Use the package database/sql to query the database
	_, err = r.db.Exec("UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?", p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id)
	if err != nil {
		return err
	}
	return nil
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
