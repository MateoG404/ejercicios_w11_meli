// This code has the schema of the ticket repository (map)

// Interface for the ticket repository
package internal

type TicketRepository interface {
	// Load the tickets from the files to the repository and save in the memory
	LoadTickets() error
	// Get the tickets from the repository and return in a slice of tickets
	GetTickets() []Ticket
	// Get the tickets from the repository and return in a slice of tickets according to the country input
	GetTicketsByCountry() []Ticket
	// Get the tickets from the repository and return in a slice of tickets according to the proportion input
	GetTicketProportion() int

	// Add, update and delete tickets from the repository
	AddTicket(ticket Ticket) error
	UpdateTicket(ticket Ticket) error
	DeleteTicket(ID string) error
}
