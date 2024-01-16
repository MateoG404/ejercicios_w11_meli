// This code contains the implementation of the handler for the ticket service
package handler

import (
	"desafio_cierre/app/internal/service"
	"desafio_cierre/app/platform/response"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// This is the struct of the ticket handler
type TicketHandler struct {
	sv service.TicketService
}

// This is the struct of the struct Ticket
type Ticket struct {
	ID             string
	Name           string
	Email          string
	Country        string
	Time_Arrival   string
	Time_Departure string
}

// This is the struct of the JSON Request
type JSONRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Country        string `json:"country"`
	Time_Arrival   string `json:"time_arrival"`
	Time_Departure string `json:"time_departure"`
}

// This is the struct of the JSON Response
type JSONResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Country        string `json:"country"`
	Time_Arrival   string `json:"time_arrival"`
	Time_Departure string `json:"time_departure"`
}

// This is the constructor of the struct TicketHandler
func NewTicketHandler(sv service.TicketService) *TicketHandler {
	return &TicketHandler{
		sv: sv,
	}
}

// This is the handler for the method GetTickets
func (h *TicketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	// This method dont receive any parameter in the body
	// This method dont receive any parameter in the url

	// Use the service to get the tickets

	tickets, err := h.sv.GetTickets()
	if err != nil {
		response.JSONResponse(w, http.StatusInternalServerError, err.Error())
	}

	// Create the response
	response.JSONResponse(w, http.StatusOK, tickets)

}

// This is the handler for the method GetTicketsByCountry

func (h *TicketHandler) GetTicketsByCountry(w http.ResponseWriter, r *http.Request) {
	// Process the request
	fmt.Println("entro al handler")
	country := chi.URLParam(r, "country")

	// Use the service to get the tickets
	tickets, err := h.sv.GetTicketsByCountry(country)

	// Create the response
	if err != nil {
		response.JSONResponse(w, http.StatusBadRequest, err.Error())
	}
	response.JSONResponse(w, http.StatusOK, tickets)
}
