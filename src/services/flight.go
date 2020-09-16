package services

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
)

// IFlightService is the Flight Service contract
type IFlightService interface {
	// CreateFlight create a flight record
	CreateFlight(flightVO *FlightVO) (string, error)

	// GetFlight get a flight record
	GetFlight(flightID string) (*FlightVO, error)
}

// FlightVO FlightVO
type FlightVO struct {
	FlightID         string
	AircraftID       string
	DepartureAirport string
	DepartureTime    int64
	ArrivalAirport   string
	ArrivalTime      int64
}

// FlightDefaultService is the default service implementing the logic for the project
type FlightDefaultService struct {
	FlightDatabase   databases.IFlightDatabase
	AircraftDatabase databases.IAircraftDabatase
	AirportDatabase  databases.IAirportDatabase
}

// CreateFlight create a flight record
func (service *FlightDefaultService) CreateFlight(flightVO *FlightVO) (string, error) {
	// Check airports are valid
	if flightVO.DepartureAirport == flightVO.ArrivalAirport {
		return flightVO.FlightID, common.NewErrorCode(common.ErrorIdenticalAirportCode)
	}
	airportInfo, err := service.AirportDatabase.GetAirportInfo(flightVO.DepartureAirport)
	if err != nil {
		return flightVO.FlightID, err
	}
	if airportInfo == nil {
		return flightVO.FlightID, common.NewErrorCode(common.ErrorItemNotFoundAirportCode)
	}
	airportInfo, err = service.AirportDatabase.GetAirportInfo(flightVO.ArrivalAirport)
	if err != nil {
		return flightVO.FlightID, err
	}
	if airportInfo == nil {
		return flightVO.FlightID, common.NewErrorCode(common.ErrorItemNotFoundAirportCode)
	}

	// Calculate and verify the UTC times
	departureTimeUTC := time.Unix(flightVO.DepartureTime, 0).Unix()
	arrivalTimeUTC := time.Unix(flightVO.ArrivalTime, 0).Unix()
	if departureTimeUTC > arrivalTimeUTC {
		return flightVO.FlightID, common.NewErrorCode(common.ErrorDepartureTimeAfterArrivalTimeCode)
	}

	// Verify the aircraft exist
	aircraftInfo, err := service.AircraftDatabase.GetAircraftInfo(flightVO.AircraftID)
	if err != nil {
		return flightVO.FlightID, err
	}
	if aircraftInfo == nil {
		return flightVO.FlightID, common.NewErrorCode(common.ErrorItemNotFoundAircraftCode)
	}

	// Create the flight record
	return flightVO.FlightID, service.FlightDatabase.CreateFlight(&databases.CreateFlightInfo{
		FlightID:         flightVO.FlightID,
		AircraftID:       flightVO.AircraftID,
		DepartureAirport: flightVO.DepartureAirport,
		DepartureTime:    departureTimeUTC,
		ArrivalAirport:   flightVO.ArrivalAirport,
		ArrivalTime:      arrivalTimeUTC,
	})
}

// GetFlight get a flight record
func (service *FlightDefaultService) GetFlight(flightID string) (*FlightVO, error) {
	flightInfo, err := service.FlightDatabase.GetFlightInfo(flightID)
	if err != nil {
		return nil, err
	}
	return convertFlightInfoToVO(flightInfo), err
}

func convertFlightInfoToVO(flightInfo *databases.FlightInfo) *FlightVO {
	if flightInfo == nil {
		return nil
	}
	return &FlightVO{
		FlightID:         flightInfo.FlightID,
		AircraftID:       flightInfo.AircraftID,
		DepartureAirport: flightInfo.DepartureAirport,
		DepartureTime:    flightInfo.DepartureTime,
		ArrivalAirport:   flightInfo.ArrivalAirport,
		ArrivalTime:      flightInfo.ArrivalTime,
	}
}
