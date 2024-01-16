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
	ErrDataNotFound   = fmt.Errorf("data not found")
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
func (s *TicketService) GetTicketProportion(country string) (float64, error) {
	// Bussiness logic

	// - Capitalize the first letter of the country
	country = strings.Title(country)

	// - If the country is valid, use the repository to get the tickets
	if country == "" {
		return 0, ErrInvalidCountry
	}
	// Use the repository to get the tickets
	tickets, err := s.rp.GetTicketsByCountry(country)
	if err != nil {
		return 0, err
	}
	// Use the repository to get all the tickets
	ticketsAll, err := s.rp.GetAllTickets()
	if err != nil {
		return 0, err
	}
	fmt.Println("ticketsAll", len(ticketsAll))
	// Calculate the proportion of tickets by country
	if len(ticketsAll) == 0 {
		// If there are no tickets, return an error
		return 0, ErrDataNotFound
	}
	if len(tickets) == 0 {
		// If there are no tickets, return an error
		return 0, ErrDataNotFound
	}
	// Calculate the average
	avg := float64(len(tickets)) / float64(len(ticketsAll))

	return avg, nil // return proportion as a percentage
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
