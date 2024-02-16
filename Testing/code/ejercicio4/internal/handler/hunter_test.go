// This script contains the tests for the handler package
package handler_test

import (
	"ejercicio4/internal/handler"
	"ejercicio4/internal/prey"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test for Hunter ConfigurePrey handler

func TestHanlerConfigurePreyHandler(t *testing.T) {

	// Case 1 : Success configuration of a prey for the hunter handler with a status code 200 and message "Prey configured"
	t.Run("Success configuration of a prey for the hunter handler with a status code 200 and message \"Prey configured\"", func(t *testing.T) {
		// Arrange
		// - Set the prey to use in the hunter handler
		prey := prey.NewPreyStub()

		// - Set the handler
		hd := handler.NewHunter(nil, prey)

		// - Set HandlerFunc for the handler
		hdFunc := hd.ConfigurePrey()

		// - Set the RequestBodyConfigPrey struct in JSON format
		body := `{
			"speed": 10.0,
			"position": {
				"x": 10.0,
				"y": 10.0,
				"z": 10.0
			}
		}`
		// Act
		// - Do a request to the handler (POST) with a RequestBodyConfigPrey struct

		request := httptest.NewRequest("POST", "/configure-prey", strings.NewReader(body))

		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler
		hdFunc(response, request)

		// Assert
		// - Check if the status code is 200
		expectedStatus := 200
		require.Equal(t, expectedStatus, response.Code)

		// - Check if the message is "Prey configured"
		expectedMessage := `{"message":"prey configured","data":null}`
		require.JSONEq(t, expectedMessage, response.Body.String())

		// - Check Header Content-Type is application/json
		expectedContentType := "application/json"
		require.Equal(t, expectedContentType, response.Header().Get("Content-Type"))
	})
}
