package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/darahayes/go-boom"
)

type aircraftModelRequestBody struct {
	Model string `json:"model"`
}

type aircraftModelResponseBody struct {
	ID string `json:"id"`
}

func handleCreateAircraftModel(w http.ResponseWriter, r *http.Request) {
	var reqBody aircraftModelRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	if len(reqBody.Model) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("model"))
		return
	}
	aircraftModelID, err := AircraftService.CreateAircraftModel(&services.AircraftModelVO{
		ID: reqBody.Model,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(&aircraftModelResponseBody{
		ID: aircraftModelID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}
