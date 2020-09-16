package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/darahayes/go-boom"

	"github.com/gorilla/mux"
)

type createPortfolioAircraftRequestBody struct {
	AircraftID string `json:"aircraftId"`
}

func handleCreatePortfolioAircraft(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	portfolioID := params["portfolioId"]
	if len(portfolioID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("portfolioId"))
		return
	}
	var reqBody createPortfolioAircraftRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	err = PortfolioService.CreatePortfolioAircraft(&services.PortfolioAircraftVO{
		PortfolioID: portfolioID,
		AircraftID:  reqBody.AircraftID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func handleDeletePortfolioAircraft(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	portfolioID := params["portfolioId"]
	if len(portfolioID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("portfolioId"))
		return
	}
	aircraftID := params["aircraftId"]
	if len(aircraftID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("aircraftId"))
		return
	}
	err := PortfolioService.DeletePortfolioAircraft(portfolioID, aircraftID)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
