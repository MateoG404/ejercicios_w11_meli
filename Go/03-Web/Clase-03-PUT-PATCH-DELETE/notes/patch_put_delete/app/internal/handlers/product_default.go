// This is the script to create a handler for the default product

package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	service "patch_put_delete/app/internal/service"
)

// Create the struct for the handler
type ProductDefaultHandler struct {
	s service.ProductService
}

// Struct for the request
type RequestJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Struct for the response
type ResponseJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Create a new handler

func NewHandler(service service.ProductService) *ProductDefaultHandler {
	return &ProductDefaultHandler{
		s: service,
	}
}

// Create a new product handler

func (h *ProductDefaultHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	// Read the request body

	body := RequestJSON{}

	// Verify if the request body is valid that all the data are given in the request body by the user
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}
	// Validate the request

	// Validate all the data are given in the request body by the user

	if body.ID == 0 || body.Name == "" || body.Description == "" || body.Price == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	// Process the request
	// Create a new product
	_, err := h.s.CreateProduct(body.ID, body.Name, body.Description, body.Price)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid request body"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Set Status
	w.WriteHeader(http.StatusCreated)

	// marshal body to json
	bytes, err := json.Marshal(body)
	if err != nil {
		// default error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(bytes)
}
