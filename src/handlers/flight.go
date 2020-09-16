package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/darahayes/go-boom"
)

type flightRequestBody struct {
	FlightID         string `json:"flight_number"`
	AircraftID       string `json:"registration"`
	DepartureAirport string `json:"departure_airport"`
	DepartureTime    int64  `json:"departure_timestamp"`
	ArrivalAirport   string `json:"arrival_airport"`
	ArrivalTime      int64  `json:"arrival_timestamp"`
}

type flightResponseBody struct {
	FlightID string `json:"flightId"`
}

func handleCreateFlight(w http.ResponseWriter, r *http.Request) {
	var reqBody flightRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	// Check mandatory body params
	if len(reqBody.FlightID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("flight_number"))
		return
	}
	if len(reqBody.AircraftID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("registration"))
		return
	}
	if len(reqBody.DepartureAirport) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("departure_airport"))
		return
	}
	if reqBody.DepartureTime == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("departure_timestamp"))
		return
	}
	if len(reqBody.ArrivalAirport) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("arrival_airport"))
		return
	}
	if reqBody.ArrivalTime == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("arrival_timestamp"))
		return
	}
	flightID, err := FlightService.CreateFlight(&services.FlightVO{
		FlightID:         reqBody.FlightID,
		AircraftID:       reqBody.AircraftID,
		DepartureAirport: reqBody.DepartureAirport,
		DepartureTime:    reqBody.DepartureTime,
		ArrivalAirport:   reqBody.ArrivalAirport,
		ArrivalTime:      reqBody.ArrivalTime,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(&flightResponseBody{
		FlightID: flightID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}
