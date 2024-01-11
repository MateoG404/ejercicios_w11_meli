// This code is the implementation of the service for the product model
// that recives the repository as a dependency and its used by the handler

package service

// Struct to represent the service of the product model

// Struct to represent the service of the product model
import "post/app/internal/repository"

type ProductRepoService struct {
	// The service needs the repository, for that reason we define it as a dependency
	// We create an attribute of type ProductMap
	rp repository.ProductMap
}

// Method to create a new product service
func NewProductRepoService(rp repository.ProductMap) *ProductRepoService {
	return &ProductRepoService{
		rp: rp,
	}
}
