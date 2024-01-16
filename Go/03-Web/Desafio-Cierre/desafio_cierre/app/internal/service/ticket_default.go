// This code contains the methods and functions to implement the ticket service
package service

import (
	"desafio_cierre/app/internal"
	"desafio_cierre/app/internal/repository"
	"fmt"
	"strings"
)

var (
	ErrInvalidCountry = fmt.Errorf("invalid country format ")
)

// This is the struct of the ticket service
type TicketService struct {
	rp repository.TicketMap
}

// This is the constructor of the struct TicketService
func NewTicketService(rp repository.TicketMap) *TicketService {
	return &TicketService{
		rp: rp,
	}
}

// Method to get all the tickets from the repository and return in a slice of tickets

func (s *TicketService) GetTickets() (tickets map[string]internal.Ticket, err error) {
	// Use the repository to get the tickets
	return s.rp.GetAllTickets()
}

// Method to get the tickets from the repository and return in a slice of tickets according to the country input
func (s *TicketService) GetTicketsByCountry(country string) (tickets map[string]internal.Ticket, err error) {
	// Bussiness logic

	// - Capitalize the first letter of the country
	country = strings.Title(country)
	fmt.Println("country", country)

	// - If the country is valid, use the repository to get the tickets
	if country == "" {
		return nil, ErrInvalidCountry
	}

	// Use the repository to get the tickets
	tickets, err = s.rp.GetTicketsByCountry(country)
	if err != nil {
		return nil, err
	}

	return tickets, nil

}

// Method to get the tickets from the repository and return in a slice of tickets according to the proportion input
func (s *TicketService) GetTicketProportion() int {
	return 0
}

// Method to add, update and delete tickets from the repository
func (s *TicketService) AddTicket(ticket internal.Ticket) error {
	return nil
}

func (s *TicketService) UpdateTicket(ticket internal.Ticket) error {
	return nil
}

func (s *TicketService) DeleteTicket(ID string) error {
	return nil
}
