// This file contains the implementation of the server using the chi library
package application

import (
	"database/sql"
	handler "ejercicio1/app/internal/handler"
	repository "ejercicio1/app/internal/repository"
	service "ejercicio1/app/internal/service"
	"fmt"

	"github.com/go-chi/chi/v5"
)

type ServerChi struct {
	address  string
	mySqlDSN string
}

// Create the struct for the configuration of the server
type ServerChiConfig struct {
	Address  string
	MySqlDSN string
}

// Create a new server using the chi library
func NewServerChi(serverChiConfig ServerChiConfig) *ServerChi {
	// Verify the address and the mySqlDSN
	if serverChiConfig.Address == "" {
		serverChiConfig.Address = ":8080"
	}

	return &ServerChi{
		address:  serverChiConfig.Address,
		mySqlDSN: serverChiConfig.MySqlDSN,
	}
}

// Start the server using the chi library

func (s *ServerChi) Start() (err error) {

	// Configuration dependencies

	// - Database

	db, err := sql.Open("mysql", s.mySqlDSN)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test connection using Ping
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")

	// - Create the repository
	rp := repository.NewProductStore(db)

	// - Create the service
	svc := service.NewProductService(rp)

	// - Create the handler
	hd := handler.NewHandlerProduct(svc)

	// Create the server
	router := chi.NewRouter()

	router.Get("/products", hd.GetById())
	// - Listen the server
	return
}
