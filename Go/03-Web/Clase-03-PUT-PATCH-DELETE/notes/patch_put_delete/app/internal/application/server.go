// This script contains the server of the application using the chi router

package application

import (
	"net/http"

	"patch_put_delete/app/internal/handlers"
	repository "patch_put_delete/app/internal/repository"
	"patch_put_delete/app/internal/service"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	address string
}

// Create a new server using the chi router
func NewServerChi(address string) *Server {
	address = ":" + address
	return &Server{
		address: address,
	}
}

// Configurate and Run the server

func (s *Server) Run() error {
	// Dependencies configuration

	// - repository
	db := repository.NewProductMap()

	// - service

	sv := service.NewService(*db)

	// - handler

	hd := handlers.NewHandler(*sv)

	// - router
	router := chi.NewRouter()

	// Configuration new endpoints

	// - GET /products/{id}
	router.Get("/products/{id}", hd.GetProductById)

	// - POST /products

	router.Post("/products", hd.CreateProduct)

	// - PUT /products/{id}
	router.Put("/products/{id}", hd.UpdateProductById)

	// - PATCH /products/{id}
	router.Patch("/products/{id}", hd.UpdateProductByIdPatch)
	// run server
	// run server
	err := http.ListenAndServe(s.address, router)
	if err != nil {
		return err
	}
	return nil

}
