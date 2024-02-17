// This script contains the tests for the handler package
package handler_test

import (
	"ejercicio4/internal/handler"
	hunter "ejercicio4/internal/hunter"
	"ejercicio4/internal/prey"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	// ErrCanNotHunt is returned when the hunter can not hunt the prey
	ErrCanNotHunt = errors.New("can not hunt the prey")
)

// Test for Hunter ConfigurePrey handler

func TestHandlerConfigurePreyHandler(t *testing.T) {

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

	// Case 2 : Fail configuration of a prey for the hunter handler with a status code 400 and message "Bad request" created by incorrect JSON format
	t.Run("Fail configuration of a prey for the hunter handler with a status code 400 and message \"Bad request\"", func(t *testing.T) {
		// Arrange
		// - Set the prey to use in the hunter handler
		prey := prey.NewPreyStub()

		// - Set the handler
		hd := handler.NewHunter(nil, prey)

		// - Set HandlerFunc for the handler
		hdFunc := hd.ConfigurePrey() // This will be tested
		// - Set the body
		body := "invalid request body"
		// Act
		// - Create the request
		request := httptest.NewRequest("POST", "/", strings.NewReader(body))
		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler
		hdFunc(response, request)

		// Assert
		// - Check if the status code is 400
		expectedCode := 400
		require.Equal(t, expectedCode, response.Code)

		// - Check if the message is "Bad request"
		expectedMessage := `{"message":"invalid request body", "status":"Bad Request"}`
		require.JSONEq(t, expectedMessage, response.Body.String())

		// - Check Header Content-Type is application/json
		expectedContentType := "application/json"
		require.Equal(t, expectedContentType, response.Header().Get("Content-Type"))
	})
}

// Test for Hunter ConfigureHunter handler

func TestHandlerConfigureHunter(t *testing.T) {
	// Case 1 : Success configuration of a hunter handler with a status code 200 and message "Hunter configured"
	t.Run("Success configuration of a hunter handler with a status code 200 and message \"Hunter configured\"", func(t *testing.T) {
		// Arrange
		// - Set the hunter to use in the hunter handler
		hunter := hunter.NewHunterMock()

		// - Set Hunter Handler
		hd := handler.NewHunter(hunter, nil)

		// - Set HandlerFunc for the handler. (This will be tested)
		hdFunc := hd.ConfigureHunter()
		// - Set body
		body := `{
			"speed": 10.0,
			"position": {
				"x": 10.0,
				"y": 10.0,
				"z": 10.0
			}
		}`
		// Act
		// - Create the request
		request := httptest.NewRequest("POST", "/configure-hunter", strings.NewReader(body))

		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler using the handler function
		hdFunc(response, request)

		// Assert
		// - Check if the status code is 200
		expectedCode := 200
		require.Equal(t, expectedCode, response.Code)

		// - Check if the message is "Hunter configured"
		expectedMessage := `{"message":"hunter configured","data":null}`
		require.JSONEq(t, expectedMessage, response.Body.String())

		// - Check Header Content-Type is application/json
		expectedContentType := "application/json"
		require.Equal(t, expectedContentType, response.Header().Get("Content-Type"))
	})

	// Case 2 : Fail configuration of a hunter handler with a status code 400 and message "Bad request" created by incorrect JSON format
	t.Run("Fail configuration of a hunter handler with a status code 400 and message \"Bad request\"", func(t *testing.T) {
		// Arrange
		// - Set the hunter to use in the hunter handler
		hunter := hunter.NewHunterMock()

		// - Set the handler with the hunter
		handler := handler.NewHunter(hunter, nil)

		// - Set HandlerFunc for the handler
		handlerFunc := handler.ConfigureHunter() // This will be tested

		// - Set the body
		body := "invalid request body"

		// Act
		// - Create the request
		request := httptest.NewRequest("POST", "/configure-hunter", strings.NewReader(body))

		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler using the handler function
		handlerFunc(response, request)

		// Assert
		// - Check if the status code is 400
		expectedCode := 400
		require.Equal(t, expectedCode, response.Code)

		// - Check if the message is "Bad request"
		expectedMessage := `{"status":"Bad Request","message":"invalid request body"}`
		require.JSONEq(t, expectedMessage, response.Body.String())

		// - Check Header Content-Type is application/json
		require.Equal(t, "application/json", response.Header().Get("Content-Type"))
	})
}

// Test for Hunt

func TestHandlerHunt(t *testing.T) {
	// Case 1 : The hunter handler successfully hunts the prey with a status code 200 and message "Hunt done"
	// 			And data show the info of the hunter and time of the hunt

	t.Run("The hunter handler successfully hunts the prey with a status code 200 and message \"Hunt done\"", func(t *testing.T) {
		// Arrange
		// - Set the prey to use in the hunter handler
		prey := prey.NewPreyStub()

		// - Set the hunter to use in the hunter handler
		hunter := hunter.NewHunterMock()

		// - Set the handler
		handler := handler.NewHunter(hunter, prey)

		// - Set the handler function
		handlerFunc := handler.Hunt()

		// Act
		// - Create the request
		request := httptest.NewRequest("POST", "/hunt", nil)

		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler
		handlerFunc(response, request)

		// Assert

		// - Check if the status code is 200
		expectedCode := 200
		require.Equal(t, expectedCode, response.Code)

		// - Check if the message is "Hunt done"
		expectedMessage := `{
			"data": {
				"duration": 0,
				"success": true
			},
			"message": "hunt done"
		}`
		require.JSONEq(t, expectedMessage, response.Body.String())

		// - Check Header Content-Type is application/json
		expectedContentType := "application/json"
		require.Equal(t, expectedContentType, response.Header().Get("Content-Type"))
	})

	// Case 2 : The hunter handler fails to hunt the prey with a status code 400 and message "Bad request"
	t.Run("The hunter handler fails to hunt the prey with a status code 400 and message \"Bad request\"", func(t *testing.T) {
		// Arrange
		// - Set the prey to use in the hunter handler
		p := prey.NewPreyStub()
		// - Set the hunter to use in the hunter handler
		hunter := hunter.NewHunterMock()
		// - Set the hunter cannot hunt the prey
		hunter.HuntFunc = func(p prey.Prey) (duration float64, err error) {
			return 0.0, ErrCanNotHunt
		}
		// - Set the handler
		handler := handler.NewHunter(hunter, p)
		// - Set the handler function
		handlerFunc := handler.Hunt() // This will be tested

		// Act
		// - Create the request
		request := httptest.NewRequest("POST", "/hunt", nil)

		// - Configure the header Content-Type as application/json
		request.Header.Set("Content-Type", "application/json")

		// - Set the response recorder
		response := httptest.NewRecorder()

		// - Execute the handler
		handlerFunc(response, request)

		// Assert
		// - Check if the status code is 400
		expectedCode := 500
		require.Equal(t, expectedCode, response.Code)
		// - Check if the message is "Bad request"
		expectedMessage := `{"status":"Internal Server Error","message":"internal server error"}`
		require.JSONEq(t, expectedMessage, response.Body.String())
		// - Check Header Content-Type is application/json
		expectedContentType := "application/json"
		require.Equal(t, expectedContentType, response.Header().Get("Content-Type"))
	})
}
