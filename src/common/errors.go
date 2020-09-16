package common

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
)

const (
	// Business Rule

	// ErrorIdenticalAirportCode ErrorIdenticalAirportCode
	ErrorIdenticalAirportCode    = "201"
	errorIdenticalAirportMessage = "Airports can not be the same"

	// ErrorDepartureTimeAfterArrivalTimeCode ErrorDepartureTimeAfterArrivalTimeCode
	ErrorDepartureTimeAfterArrivalTimeCode    = "202"
	errorDepartureTimeAfterArrivalTimeMessage = "Departure can not happen after Arrival"

	// Item not found

	// ErrorItemNotFoundPortfolioCode ErrorItemNotFoundPortfolioCode
	ErrorItemNotFoundPortfolioCode    = "401"
	errorItemNotFoundPortfolioMessage = "portfolio"

	// ErrorItemNotFoundAircraftCode ErrorItemNotFoundAircraftCode
	ErrorItemNotFoundAircraftCode    = "402"
	errorItemNotFoundAircraftMessage = "aircraft"

	// ErrorItemNotFoundAirportCode ErrorItemNotFoundAirportCode
	ErrorItemNotFoundAirportCode    = "403"
	errorItemNotFoundAirportMessage = "airport"

	// ErrorItemNotFoundAircraftModelCode ErrorItemNotFoundAircraftModelCode
	ErrorItemNotFoundAircraftModelCode    = "404"
	errorItemNotFoundAircraftModelMessage = "aircraftModel"
)

// WriteError write in the response writer the appropriate error message based on the error code
func WriteError(w http.ResponseWriter, err error) {
	switch err.Error() {
	case ErrorItemNotFoundPortfolioCode:
		boom.NotFound(w, getErrorItemNotFound(errorItemNotFoundPortfolioMessage))
		return
	case ErrorItemNotFoundAircraftCode:
		boom.NotFound(w, getErrorItemNotFound(errorItemNotFoundAircraftMessage))
		return
	case ErrorItemNotFoundAirportCode:
		boom.NotFound(w, getErrorItemNotFound(errorItemNotFoundAirportMessage))
		return
	case ErrorItemNotFoundAircraftModelCode:
		boom.NotFound(w, getErrorItemNotFound(errorItemNotFoundAircraftModelMessage))
		return
	case ErrorIdenticalAirportCode:
		boom.BadRequest(w, getErrorItemNotFound(errorIdenticalAirportMessage))
		return
	case ErrorDepartureTimeAfterArrivalTimeCode:
		boom.BadRequest(w, getErrorItemNotFound(errorDepartureTimeAfterArrivalTimeMessage))
		return
	default:
		boom.Internal(w)
	}
}

// NewErrorCode generate a new error containing the specific code
func NewErrorCode(code string) error {
	return errors.New(code)
}

// GetErrorMissingParam return a message for a missing param
func GetErrorMissingParam(paramName string) string {
	return fmt.Sprintf("Missing parameter - %s", paramName)
}

// GetErrorBadRequestBody return Bad Request Body
func GetErrorBadRequestBody() string {
	return "Bad Request Body"
}

func getErrorItemNotFound(itemName string) string {
	return fmt.Sprintf("Item not found - %s", itemName)
}
