// This script is created to use the package json in Go

package main

import (
	"encoding/json"
	"fmt"
)

// Creation of the struct to use the package json

type Note struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {

	// Creation new data for the struct
	n := Note{
		ID:     1,
		Title:  "Note 1",
		Author: "Mateo Gutierrez",
	}

	// Creation of the json using the package json

	JsonData, err := json.Marshal(n)
	if err != nil {
		fmt.Println(err)
	}

	// Print the json data
	fmt.Println(string(JsonData))

	// Use of unmarshall function

	JsonData = []byte(`{"ID": 2, "Title": "Note 2", "Author": "Mateo Gutierrez"}`)

	var p Note

	if err := json.Unmarshal(JsonData, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(p.Author)
}
