package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/darahayes/go-boom"
	"github.com/gorilla/mux"
)

type portfolioInfoListResponseBody struct {
	Portfolios []*portfolioInfoResponseBody `json:"portfolios"`
}

type portfolioInfoResponseBody struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	AircraftIds []string `json:"aircraftIds,omitempty"`
}

type createPortfolioRequestBody struct {
	Name string `json:"name"`
}

type createPortfolioResponseBody struct {
	ID string `json:"id"`
}

func handleGetPortfolio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	portfolioID := params["portfolioId"]
	if len(portfolioID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("portfolioId"))
		return
	}
	portfolioVO, err := PortfolioService.GetPortfolio(portfolioID)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(convertPortfolioVOToRequestBody(portfolioVO))
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}

func handleGetPortfolioList(w http.ResponseWriter, r *http.Request) {
	portfolioVOList, err := PortfolioService.GetPortfolioList()
	if err != nil {
		common.WriteError(w, err)
		return
	}
	portfolioInfoResponseList := make([]*portfolioInfoResponseBody, len(portfolioVOList))
	for i, portfolioVO := range portfolioVOList {
		portfolioInfoResponseList[i] = convertPortfolioVOToRequestBody(portfolioVO)
	}
	response, err := json.Marshal(&portfolioInfoListResponseBody{
		Portfolios: portfolioInfoResponseList,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}

func handleCreatePortfolio(w http.ResponseWriter, r *http.Request) {
	var reqBody createPortfolioRequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		boom.BadRequest(w, common.GetErrorBadRequestBody())
		return
	}
	if len(reqBody.Name) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("name"))
		return
	}
	portfolioID, err := PortfolioService.CreatePortfolio(&services.PortfolioVO{
		Name: reqBody.Name,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	response, err := json.Marshal(&createPortfolioResponseBody{
		ID: portfolioID,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}

func handleDeletePortfolio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	portfolioID := params["portfolioId"]
	if len(portfolioID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("portfolioId"))
		return
	}
	err := PortfolioService.DeletePortfolio(portfolioID)
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func convertPortfolioVOToRequestBody(portfolioVO *services.PortfolioVO) *portfolioInfoResponseBody {
	if portfolioVO == nil {
		return nil
	}
	aircraftIds := []string{}
	for _, aircraftVO := range portfolioVO.AircraftVOList {
		aircraftIds = append(aircraftIds, aircraftVO.ID)
	}
	return &portfolioInfoResponseBody{
		ID:          portfolioVO.ID,
		Name:        portfolioVO.Name,
		AircraftIds: aircraftIds,
	}
}
