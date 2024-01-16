// This code contains a struct to create the JSON response

package response

import (
	"encoding/json"
	"net/http"
)

// Function to return a JSON response

func JSONResponse(w http.ResponseWriter, status int, body any) {

	// Write the header
	w.Header().Set("Content-Type", "application/json")

	// Set the status code
	w.WriteHeader(status)

	// Write the body

	if body == nil {
		w.WriteHeader(status)
	}

	// We need to use the marshall to transform the body to JSON

	marshalJson, err := json.Marshal(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(marshalJson)
}
