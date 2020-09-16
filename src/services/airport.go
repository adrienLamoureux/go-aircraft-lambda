package services

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
)

// IAirportService is the Airport Service contract
type IAirportService interface {
	// CreateAirport create an allowed airport
	CreateAirport(airportVO *AirportVO) (string, error)

	// GetAirport get an allowed airport
	GetAirport(airportID string) (*AirportVO, error)
}

// AirportVO AirportVO
type AirportVO struct {
	ID string
}

// AirportDefaultService is the default service implementing the logic for the project
type AirportDefaultService struct {
	AirportDatabase databases.IAirportDatabase
}

// CreateAirport create an allowed airport
func (service *AirportDefaultService) CreateAirport(airportVO *AirportVO) (string, error) {
	return airportVO.ID, service.AirportDatabase.CreateAirport(&databases.CreateAirportInfo{
		ID: airportVO.ID,
	})
}

// GetAirport get an allowed airport
func (service *AirportDefaultService) GetAirport(airportID string) (*AirportVO, error) {
	airportInfo, err := service.AirportDatabase.GetAirportInfo(airportID)
	if err != nil {
		return nil, err
	}
	return convertAirportInfoToVO(airportInfo), err
}

func convertAirportInfoToVO(airportInfo *databases.AirportInfo) *AirportVO {
	if airportInfo == nil {
		return nil
	}
	return &AirportVO{
		ID: airportInfo.ID,
	}
}
