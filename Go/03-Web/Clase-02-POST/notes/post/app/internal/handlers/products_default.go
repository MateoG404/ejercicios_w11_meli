// This code contains the implementation of the default handler of the product model

package handlers

import (
	"fmt"
	"net/http"
	"post/app/internal/service"
	"post/app/platform/request"
)

// __________________________________________________________________

// Struct for the handler implementation
type ProductsHandler struct {
	sv service.ProductRepoService
}

// Create new handler

func NewProductsHandler(sv service.ProductRepoService) *ProductsHandler {
	return &ProductsHandler{
		sv: sv,
	}
}

// Struct to use in the Request and Response body of the product model

// RequestBodyProduct is the struct that represents the request body of a product
type RequestBodyProduct struct {
	Name     string  `json:"name"`
	Type     string  `json:"type"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

// ResponseBodyProduct is the struct that represents the response body of a product
type ResponseBodyProduct struct {
	Message string `json:"message"`
	Data    *struct {
		Id       int     `json:"id"`
		Name     string  `json:"name"`
		Type     string  `json:"type"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	} `json:"data"`
	Error bool `json:"error"`
}

// __________________________________________________________________

// Handler to create a product in the storage using the RequestBodyProduct struct and the ResponseBodyProduct struct

func (h *ProductsHandler) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request

		// Get the request body and pass it to the RequestBodyProduct struct

		var requestBody RequestBodyProduct

		// Use the platform request to use the method to serialize the request body in a RequestBodyProduct struct

		if err := request.ToJSON(*r, requestBody); err != nil {
			// If there is an error, return the error
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("Request body: ", requestBody)
		// Process

		// Response

		// Return the response
	}

}
