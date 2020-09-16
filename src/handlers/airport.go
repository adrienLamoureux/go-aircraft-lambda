package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/darahayes/go-boom"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
)

type airportRequestBody struct {
	AirportID string `json:"airportId"`
}

type airportResponseBody struct {
	AirportID string `json:"airportId"`
}

func handleCreateAirport(w http.ResponseWriter, r *http.Request) {
	var reqBody airportRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	if len(reqBody.AirportID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("airportId"))
		return
	}
	airportID, err := AirportService.CreateAirport(&services.AirportVO{
		ID: reqBody.AirportID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(&airportResponseBody{
		AirportID: airportID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}
