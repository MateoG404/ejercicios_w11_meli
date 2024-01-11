// This code contains the controller of the product model

package internal

// RepositoryProducts is the interface that represents the repository of the product model
type RepositoryProducts interface {
	// Save the product in the storage
	Save(product *Product) error
}
