package handlers

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/darahayes/go-boom"

	"github.com/gorilla/mux"
)

type portfolioFlightMetricResponseBody struct {
	Aircrafts []*aircraftFlightMetricResponseBody `json:"aircrafts"`
}

type aircraftFlightMetricResponseBody struct {
	AircraftID    string  `json:"aircraftId"`
	AircraftModel string  `json:"aircraftModel"`
	FlightCount   int32   `json:"flightCount"`
	FlightTime    float64 `json:"flightTime"`
}

type aircraftModelFlightMetricListResponseBody struct {
	AircraftModels []*aircraftModelFlightMetricResponseBody `json:"aircraftModels"`
}

type aircraftModelFlightMetricResponseBody struct {
	AircraftModel string  `json:"aircraftModel"`
	FlightCount   int32   `json:"flightCount"`
	FlightTime    float64 `json:"flightTime"`
}

func handleGetPortfolioFlightMetricList(w http.ResponseWriter, r *http.Request) {
	// Check mandatory params
	params := mux.Vars(r)
	portfolioID := params["portfolioId"]
	if len(portfolioID) == 0 {
		boom.BadRequest(w, common.GetErrorMissingParam("portfolioId"))
		return
	}
	var durationSec *int64
	keys, ok := r.URL.Query()["duration"]
	if ok && len(keys[0]) > 0 {
		durationHours, err := strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			boom.BadRequest(w, common.GetErrorMissingParam("duration"))
			return
		}
		durationSeconds := durationHours * 3600
		durationSec = &durationSeconds
	}

	// Call service to get the computed portfolio flight metric
	aircraftFlightMetricVOList, err := MetricService.GetPortfolioFlightMetric(portfolioID, durationSec)
	if err != nil {
		common.WriteError(w, err)
		return
	}

	// Return the result in the response
	aircraftFlightMetricResponseBodyList := make([]*aircraftFlightMetricResponseBody, len(aircraftFlightMetricVOList))
	for i, aircraftFlightMetricVO := range aircraftFlightMetricVOList {
		aircraftFlightMetricResponseBodyList[i] = convertAircraftFlightMetricVOToResponse(aircraftFlightMetricVO)
	}
	response, err := json.Marshal(&portfolioFlightMetricResponseBody{
		Aircrafts: aircraftFlightMetricResponseBodyList,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}

func handleGetAircraftModelFlightMetricList(w http.ResponseWriter, r *http.Request) {
	// Check mandatory params
	var durationSec *int64
	keys, ok := r.URL.Query()["duration"]
	if ok && len(keys[0]) > 0 {
		durationHours, err := strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			boom.BadRequest(w, common.GetErrorMissingParam("duration"))
			return
		}
		durationSeconds := durationHours * 3600
		durationSec = &durationSeconds
	}

	// Call service to get the computed flight metric for every aircraft model
	aircraftModelFlightMetricVOList, err := MetricService.GetAircraftModelFlightMetric(durationSec)
	if err != nil {
		common.WriteError(w, err)
		return
	}

	// Return the result in the response
	aircraftModelFlightResponseBodyList := make([]*aircraftModelFlightMetricResponseBody, len(aircraftModelFlightMetricVOList))
	for i, aircraftModelFlightMetricVO := range aircraftModelFlightMetricVOList {
		aircraftModelFlightResponseBodyList[i] = convertAircraftModelFlightMetricVOToResponse(aircraftModelFlightMetricVO)
	}
	response, err := json.Marshal(&aircraftModelFlightMetricListResponseBody{
		AircraftModels: aircraftModelFlightResponseBodyList,
	})
	if err != nil {
		common.WriteError(w, err)
		return
	}
	w.Write(response)
}

func convertAircraftModelFlightMetricVOToResponse(aircraftModelFlightMetricVO *services.AircraftModelFlightMetricVO) *aircraftModelFlightMetricResponseBody {
	if aircraftModelFlightMetricVO == nil {
		return nil
	}
	return &aircraftModelFlightMetricResponseBody{
		AircraftModel: aircraftModelFlightMetricVO.AircraftModel,
		FlightCount:   aircraftModelFlightMetricVO.FlightCount,
		FlightTime:    flightTimeSecToHours(aircraftModelFlightMetricVO.FlightTime),
	}
}

func convertAircraftFlightMetricVOToResponse(aircraftFlightMetricVO *services.AircraftFlightMetricVO) *aircraftFlightMetricResponseBody {
	if aircraftFlightMetricVO == nil {
		return nil
	}
	return &aircraftFlightMetricResponseBody{
		AircraftID:    aircraftFlightMetricVO.AircraftID,
		AircraftModel: aircraftFlightMetricVO.AircraftModel,
		FlightCount:   aircraftFlightMetricVO.FlightCount,
		FlightTime:    flightTimeSecToHours(aircraftFlightMetricVO.FlightTime),
	}
}

func flightTimeSecToHours(flightTime int64) float64 {
	return math.Round(float64(flightTime)/float64(3600.0)*10) / 10
}
