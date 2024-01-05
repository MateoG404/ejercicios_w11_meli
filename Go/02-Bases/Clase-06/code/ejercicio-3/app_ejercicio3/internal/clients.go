// Script that implements the struct and methods of the clients
package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Global variable to store if a panic occurs
var (
	panicErros = false
	red        = color.New(color.FgRed).SprintFunc()
)

// Function to update the panic errors to true if a panic occurs
func updatePanicErros() {
	panicErros = true
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
	defer updatePanicErros()
	// Verify if the file exists
	file, ok := c.VerifyFile()
	if ok {

		// Read the file until the end to find the client

		reader, err := io.ReadAll(file)
		fmt.Println(file)
		if err != nil {
			panic(errors.New(red("The file could not ddddsbe read")))
		}

		// Convert the file to string
		fileString := string(reader)
		fmt.Println("acaaaaa")

		fmt.Println(fileString)
		// Search the client in the file
		if strings.Contains(fileString, c.ID) {
			return true
		}
	}
	return false
}

// VerifyFile is a method created to verify if the file exists
func (c Client) VerifyFile() (*os.File, bool) {

	// Open the file
	file, err := os.Open(c.FileName)

	// Check if there was an error opening the file
	if err != nil {
		fmt.Println("aca")
		panic(errors.New(red("The indicated file was not found or is damaged")))
	}

	// Defer file.Close()
	defer func() {
		if err := file.Close(); err != nil {
			panic(red(err))
		}
	}()

	defer updatePanicErros()

	// Check if the file exists if not throw a panic
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	return file, true
}
