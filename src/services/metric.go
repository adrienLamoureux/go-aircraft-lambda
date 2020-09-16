package services

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
)

// IMetricService is the Metric Service contract
type IMetricService interface {

	// GetPortfolioFlightMetric get the flight metric of a specific portfolio since a certain duration
	GetPortfolioFlightMetric(portfolioID string, duration *int64) ([]*AircraftFlightMetricVO, error)

	// GetAircraftModelFlightMetric get the flight metric for every aircraft model since a certain duration
	GetAircraftModelFlightMetric(duration *int64) ([]*AircraftModelFlightMetricVO, error)
}

// AircraftFlightMetricVO AircraftFlightMetricVO
type AircraftFlightMetricVO struct {
	AircraftID    string
	AircraftModel string
	FlightCount   int32
	FlightTime    int64
}

// AircraftModelFlightMetricVO AircraftModelFlightMetricVO
type AircraftModelFlightMetricVO struct {
	AircraftModel string
	FlightCount   int32
	FlightTime    int64
}

// MetricDefaultService is the default service implementing the logic for the project
type MetricDefaultService struct {
	PortfolioDatabase         databases.IPortfolioDatabase
	PortfolioAircraftDatabase databases.IPortfolioAircraftDatabase
	AircraftDatabase          databases.IAircraftDabatase
	FlightDatabase            databases.IFlightDatabase
}

// GetPortfolioFlightMetric get the flight metric of a specific portfolio since a certain duration
// Metrics computation are done per request and in-memory at every call of this methode
func (service *MetricDefaultService) GetPortfolioFlightMetric(portfolioID string, duration *int64) ([]*AircraftFlightMetricVO, error) {
	// Retrieve a portfolio with its aircrafts
	portfolioInfo, err := service.PortfolioDatabase.GetPortfolioInfo(portfolioID)
	if err != nil {
		return nil, err
	}
	if portfolioInfo == nil {
		return nil, common.NewErrorCode(common.ErrorItemNotFoundPortfolioCode)
	}
	portfolioAircraftInfoList, err := service.PortfolioAircraftDatabase.GetPortfolioAircraftInfos(portfolioID)
	if err != nil {
		return nil, err
	}

	// Compute the flight metric for every aircraft of a porfolio
	aircraftFlightMetricVOList := make([]*AircraftFlightMetricVO, len(portfolioAircraftInfoList))
	for i, portfolioAircraftInfo := range portfolioAircraftInfoList {
		aircraftFlightMetricVO, err := service.computeFlightMetricPerAircraft(portfolioAircraftInfo.AircraftID, duration)
		if err != nil {
			return nil, err
		}
		aircraftFlightMetricVOList[i] = aircraftFlightMetricVO
	}

	// Find the aircraft model of each computed aircraft
	// TODO: Optimize with getByIds
	for _, aircraftFlightMetricVO := range aircraftFlightMetricVOList {
		aircraftInfo, err := service.AircraftDatabase.GetAircraftInfo(aircraftFlightMetricVO.AircraftID)
		if err != nil {
			return nil, err
		}
		if aircraftInfo == nil {
			return nil, common.NewErrorCode(common.ErrorItemNotFoundAircraftCode)
		}
		aircraftFlightMetricVO.AircraftModel = aircraftInfo.Model
	}

	return aircraftFlightMetricVOList, nil
}

// GetAircraftModelFlightMetric get the flight metric for every aircraft model since a certain duration
// Metrics computation are done per request and in-memory at every call of this methode
func (service *MetricDefaultService) GetAircraftModelFlightMetric(duration *int64) ([]*AircraftModelFlightMetricVO, error) {
	// Get all aircraft since an aircraft has a model.
	// Aircraft table can be scanned since its going to be way smaller than flights
	aircraftInfoList, err := service.AircraftDatabase.GetAircraftInfoList()
	if err != nil {
		return []*AircraftModelFlightMetricVO{}, err
	}
	aircraftModelMap := make(map[string][]*databases.AircraftInfo)
	for _, aircraftInfo := range aircraftInfoList {
		aircraftModelMap[aircraftInfo.Model] = append(aircraftModelMap[aircraftInfo.Model], aircraftInfo)
	}

	// Compute the flight metric for every aircraft model
	aircraftModelFlightMetricVOList := make([]*AircraftModelFlightMetricVO, 0, len(aircraftModelMap))
	for model, aircraftInfoList := range aircraftModelMap {
		flightCount := int32(0)
		flightTime := int64(0)

		// Compute the flight metric for every aircraft of a specific model and cumulate the results
		for _, aircraftInfo := range aircraftInfoList {
			aircraftFlightMetricVO, err := service.computeFlightMetricPerAircraft(aircraftInfo.ID, duration)
			if err != nil {
				return []*AircraftModelFlightMetricVO{}, err
			}
			flightCount += aircraftFlightMetricVO.FlightCount
			flightTime += aircraftFlightMetricVO.FlightTime
		}
		aircraftModelFlightMetricVOList = append(aircraftModelFlightMetricVOList, &AircraftModelFlightMetricVO{
			AircraftModel: model,
			FlightCount:   flightCount,
			FlightTime:    flightTime,
		})
	}
	return aircraftModelFlightMetricVOList, nil
}

func (service *MetricDefaultService) computeFlightMetricPerAircraft(aircraftID string, duration *int64) (*AircraftFlightMetricVO, error) {
	// TODO: Optimize with getByIds
	flightInfoList, err := service.FlightDatabase.GetFlightInfosByAircraftID(aircraftID)
	if err != nil {
		return nil, err
	}
	flightTime := int64(0)
	if duration == nil {
		for _, flightInfo := range flightInfoList {
			flightTime += flightInfo.ArrivalTime - flightInfo.DepartureTime
		}
	} else {
		timeNow := time.Now().Unix()
		timeThreshold := timeNow - *duration
		for _, flightInfo := range flightInfoList {
			if flightInfo.DepartureTime >= timeThreshold {
				flightTime += flightInfo.ArrivalTime - flightInfo.DepartureTime
			}
		}
	}
	return &AircraftFlightMetricVO{
		AircraftID:  aircraftID,
		FlightCount: int32(len(flightInfoList)),
		FlightTime:  flightTime,
	}, nil
}
