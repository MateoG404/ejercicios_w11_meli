// This code contains the server implementation for the application

package internal

import (
	"fmt"
	"net/http"
	"post/app/internal/handlers"
	repository "post/app/internal/repository"
	service "post/app/internal/service"

	"github.com/go-chi/chi/v5"
)

// Struct represent the server of the application using Chi
type ServerChi struct {
	// Adress is the direction where the server will be listening
	adress string
}

// Create new server and configurate it using Chi
func NewServerChi(adress string) *ServerChi {
	if adress != "" {
		adress = ":8080"
	}
	return &ServerChi{
		adress: adress,
	}
}

// Start the server using Chi

func (s *ServerChi) Run() (err error) {
	// To create a server we need to create in a reversre way
	// First we need to configure the dependencies

	// 1. Repository - db

	// Create the repository for the products
	repoProducts := repository.NewProductMap()

	// 2. Service - usecase
	sv := service.NewProductRepoService(*repoProducts)

	// 3. Handler - controller
	hd := handlers.NewProductsHandler(*sv)
	// 4. Router - router

	rt := chi.NewRouter()

	// Then we need to start the server
	rt.Post("/products", hd.CreateProduct())

	fmt.Println("ss")
	// 5. Server - server
	err = http.ListenAndServe(s.adress, rt)
	fmt.Println("Server running on port", s.adress)
	return
}
