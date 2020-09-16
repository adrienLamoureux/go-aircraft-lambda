package services

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
)

// IAircraftService is the Aircraft Service contract
type IAircraftService interface {
	// CreateAircraft create an aircraft
	CreateAircraft(aircraftVO *AircraftVO) (string, error)

	// CreateAircraftModel create an aircraft model
	CreateAircraftModel(AircraftModelVO *AircraftModelVO) (string, error)

	// GetAircraft get an aircraft
	GetAircraft(aircraftID string) (*AircraftVO, error)
}

// AircraftVO AircraftVO
type AircraftVO struct {
	ID    string
	Model string
}

// AircraftModelVO AircraftModelVO
type AircraftModelVO struct {
	ID string
}

// AircraftDefaultService is the default service implementing the logic for the project
type AircraftDefaultService struct {
	AircraftDatabase      databases.IAircraftDabatase
	AircraftModelDatabase databases.IAircraftModelDatabase
}

// CreateAircraft create an aircraft
func (service *AircraftDefaultService) CreateAircraft(aircraftVO *AircraftVO) (string, error) {
	aircraftModelInfo, err := service.AircraftModelDatabase.GetAircraftModelInfo(aircraftVO.Model)
	if err != nil {
		return aircraftVO.ID, err
	}
	if aircraftModelInfo == nil {
		return aircraftVO.ID, common.NewErrorCode(common.ErrorItemNotFoundAircraftModelCode)
	}
	return aircraftVO.ID, service.AircraftDatabase.CreateAircraft(&databases.CreateAircraftInfo{
		ID:    aircraftVO.ID,
		Model: aircraftVO.Model,
	})
}

// CreateAircraftModel create an aircraft model
func (service *AircraftDefaultService) CreateAircraftModel(aircraftModelVO *AircraftModelVO) (string, error) {
	return aircraftModelVO.ID, service.AircraftModelDatabase.CreateAircraftModel(&databases.CreateAircraftModelInfo{
		ID: aircraftModelVO.ID,
	})
}

// GetAircraft get an aircraft
func (service *AircraftDefaultService) GetAircraft(aircraftID string) (*AircraftVO, error) {
	aircraftInfo, err := service.AircraftDatabase.GetAircraftInfo(aircraftID)
	if err != nil {
		return nil, err
	}
	return convertAircraftInfoToVO(aircraftInfo), err
}

func convertAircraftInfoToVO(aircraftInfo *databases.AircraftInfo) *AircraftVO {
	if aircraftInfo == nil {
		return nil
	}
	return &AircraftVO{
		ID:    aircraftInfo.ID,
		Model: aircraftInfo.Model,
	}
}
