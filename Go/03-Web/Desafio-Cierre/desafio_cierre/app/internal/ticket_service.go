// This code is create to show the schema of the ticket service

// Interface for the ticket service
package internal

type TicketService interface {
	// Get the tickets from the repository and return in a slice of tickets according to the country input
	GetTicketsByCountry() []Ticket
	// Get the tickets from the repository and return in a slice of tickets according to the proportion input
	GetTicketProportion() int
	// Add, update and delete tickets from the repository
	AddTicket(ticket Ticket) error
	UpdateTicket(ticket Ticket) error
	DeleteTicket(ID string) error
}
