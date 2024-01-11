// This code has the implementation to use the JSON files in the request and response body
package request

import (
	"encoding/json"
	"net/http"
)

func ToJSON(r http.Request, data any) (err error) {

	// Check if the request recives a JSON file
	if r.Header.Get("Content-Type") != "application/json" {
		return
	}

	// Get the body and serialize it in the data because the body is in bytes[]
	body := r.Body
	if err := json.NewDecoder(body).Decode(data); err != nil {
		return err
	}

	return
}
