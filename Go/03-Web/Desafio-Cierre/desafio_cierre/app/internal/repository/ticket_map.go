// This code is created for the repository using the package map

package repository

import (
	internal "desafio_cierre/app/internal"
	"encoding/csv"
	"fmt"
	"os"
)

// Struct for the repository using the package map (Representation of the database)
type TicketMap struct {
	tickets map[string]internal.Ticket
	lastID  int
}

// This is the constructor of the struct TicketMap
func NewTicketMap() *TicketMap {
	return &TicketMap{
		tickets: make(map[string]internal.Ticket),
		lastID:  0,
	}
}

// Function created to load the tickets from the files to the repository and save in the memory
func (r *TicketMap) LoadTickets() (err error) {
	// Open the file
	filepath := os.Getenv("FILECSV")
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	// Close the file
	defer file.Close()

	// Read the file

	reader := csv.NewReader(file)

	// read all the records
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, eachrecord := range records {
		r.AddTicket(eachrecord)
	}
	return
}

// Function to Add a ticket to the repository
func (r *TicketMap) AddTicket(record []string) (err error) {
	// Create the ticket
	ticket := internal.Ticket{
		ID:             record[0],
		Name:           record[1],
		Email:          record[2],
		Country:        record[3],
		Time_Arrival:   record[4],
		Time_Departure: record[5]}
	// Add the ticket to the repository
	r.tickets[ticket.ID] = ticket
	r.lastID++
	return

}

// Function to Get a ticket from the repository
func (r *TicketMap) GetAllTickets() (tickets map[string]internal.Ticket, err error) {
	return r.tickets, nil
}

// Function to Get all the tickets by country from the repository

func (r *TicketMap) GetTicketsByCountry(country string) (tickets map[string]internal.Ticket, err error) {
	// Create a map to save the tickets by country
	tickets = make(map[string]internal.Ticket)
	// Iterate the tickets and save the tickets by country in the map
	for _, ticket := range r.tickets {
		if ticket.Country == country {
			tickets[ticket.ID] = ticket
		}
	}
	return tickets, nil
}
