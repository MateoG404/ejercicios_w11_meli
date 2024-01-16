// This is the code created for the server using the package chi

package platform

import (
	"desafio_cierre/app/internal/handler"
	repository "desafio_cierre/app/internal/repository"
	"desafio_cierre/app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// This struct is used to create the server using the package chi
type ServerChi struct {
	addr string
}

// This is the constructor of the struct ServerChi
func NewServerChi(addr string) *ServerChi {
	if addr == "" {
		addr = ":8080"
	}
	return &ServerChi{
		addr: addr,
	}
}

// This is the method Start of the struct ServerChi

func (s *ServerChi) Start() error {

	// Configuration dependencies

	// - Create the repository
	rp := repository.NewTicketMap()

	// - Load the files to the repository
	rp.LoadTickets()

	// - Create the service
	sv := service.NewTicketService(*rp)

	// - Create the handler
	hd := handler.NewTicketHandler(*sv)

	// Create the server
	router := chi.NewRouter()

	// Configurate and create the middlewares
	//router.Use(middleware.Logger)

	// - Create the router and endpoints
	router.Get("/tickets", hd.GetTickets)
	router.Get("/tickets/{country}", hd.GetTicketsByCountry)
	//router.Get("/tickets/proportion/{country}", hd.GetTicketProportion)

	// Listen the server
	http.ListenAndServe(s.addr, router)
	return nil
}
