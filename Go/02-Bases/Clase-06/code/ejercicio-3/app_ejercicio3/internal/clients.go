// Script that implements the struct and methods of the clients
package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
)

// Global variable to store if a panic occurs
var (
	PanicErros = false
	red        = color.New(color.FgRed).SprintFunc()
)

// Function to update the panic errors to true if a panic occurs
func updatePanicErros() {
	PanicErros = true
}

// Client struct that implements the clients information
type Client struct {
	FileName    string   `json:"file_name"`
	File        *os.File `json:"file"`
	Name        string   `json:"name"`
	ID          string   `json:"id"`
	PhoneNumber int      `json:"phone_number"`
	HomeNumber  int      `json:"home_number"`
}

// Methods

// NewClient method that creates a new client
func NewClient(fileName string, name string, id string, phoneNumber int, homeNumber int) Client {
	return Client{
		FileName:    fileName,
		Name:        name,
		ID:          id,
		PhoneNumber: phoneNumber,
		HomeNumber:  homeNumber,
	}
}

// Verify method that verifies if the client exists in the file
func (c Client) VerifyClientFile() bool {
	// Defers
	// Defer to run if there is a panic
	defer func() {
		if r := recover(); r != nil {
			updatePanicErros()
		}
	}()

	// Verify if the file exists
	file, ok := c.VerifyFile()
	if ok {
		// Defer file.Close() to close the file when the function ends
		defer func() {
			if err := file.Close(); err != nil {
				panic(red(err))
			}
		}()

		// Decode the JSON file into a slice of Client
		var clients []Client
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&clients); err != nil && err != io.EOF {
			panic(errors.New(red("The file could not be read")))
		}

		// Search for the client ID in the slice of Client
		for _, client := range clients {
			if client.ID == c.ID {
				return true
			}
		}
	}
	return false
}

// VerifyFile is a method created to verify if the file exists
func (c Client) VerifyFile() (*os.File, bool) {
	// Defer section
	defer func() {
		// If a panic occurs, the file will be closed and updated the panic errors to true
		if r := recover(); r != nil {
			updatePanicErros()
			fmt.Println("Panic occurred:", r)
		}
	}()

	// Open the file
	file, err := os.Open(c.FileName)
	if err != nil {
		// If there was an error opening the file or if the file does not exist, panic
		panic(errors.New(red("The indicated file was not found or is damaged")))
	}

	// If the file was opened successfully, close it when the function ends
	defer func() {
		if err := file.Close(); err != nil {
			panic(red(err))
		}
	}()

	return file, true
}
