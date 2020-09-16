package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/darahayes/go-boom"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
)

type aircraftRequestBody struct {
	ID    string `json:"id"`
	Model string `json:"model"`
}

type aircraftResponseBody struct {
	ID string `json:"id"`
}

func handleCreateAircraft(w http.ResponseWriter, r *http.Request) {
	var reqBody aircraftRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	if len(reqBody.ID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("id"))
		return
	}
	if len(reqBody.Model) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("model"))
		return
	}
	aircraftID, err := AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    reqBody.ID,
		Model: reqBody.Model,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(&aircraftResponseBody{
		ID: aircraftID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}
