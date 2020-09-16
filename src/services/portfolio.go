package services

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/google/uuid"
)

// IPortfolioService is the Portfolio Service contract
type IPortfolioService interface {
	// CreatePortfolio create a portfolio
	CreatePortfolio(portfolioVO *PortfolioVO) (string, error)

	// CreatePortfolioAircraft create a portfolio aircraft link
	CreatePortfolioAircraft(portfolioAircraftVO *PortfolioAircraftVO) error

	// GetPortfolio get a portfolio
	GetPortfolio(portfolioID string) (*PortfolioVO, error)

	// GetPortfolioList get the portfolio list
	GetPortfolioList() ([]*PortfolioVO, error)

	// DeletePortfolio delete a porfolio
	DeletePortfolio(portfolioID string) error

	// DeletePortfolioAircraft delete a portfolio aircraft link
	DeletePortfolioAircraft(portfolioID, aircraftID string) error
}

// PortfolioVO PortfolioVO
type PortfolioVO struct {
	ID             string
	Name           string
	AircraftVOList []*AircraftVO
}

// PortfolioAircraftVO PortfolioAircraftVO
type PortfolioAircraftVO struct {
	PortfolioID string
	AircraftID  string
}

// PortfolioDefaultService is the default service implementing the logic for the project
type PortfolioDefaultService struct {
	PortfolioDatabase         databases.IPortfolioDatabase
	PortfolioAircraftDatabase databases.IPortfolioAircraftDatabase
	AircraftDatabase          databases.IAircraftDabatase
}

// CreatePortfolio create a portfolio
func (service *PortfolioDefaultService) CreatePortfolio(portfolioVO *PortfolioVO) (string, error) {
	if len(portfolioVO.ID) == 0 {
		portfolioUUID, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		portfolioVO.ID = portfolioUUID.String()
	}
	return portfolioVO.ID, service.PortfolioDatabase.CreatePortfolio(&databases.CreatePortfolioInfo{
		ID:   portfolioVO.ID,
		Name: portfolioVO.Name,
	})
}

// CreatePortfolioAircraft create a portfolio aircraft link
func (service *PortfolioDefaultService) CreatePortfolioAircraft(portfolioAircraftVO *PortfolioAircraftVO) error {
	portfolioInfo, err := service.PortfolioDatabase.GetPortfolioInfo(portfolioAircraftVO.PortfolioID)
	if err != nil {
		return err
	}
	if portfolioInfo == nil {
		return common.NewErrorCode(common.ErrorItemNotFoundPortfolioCode)
	}
	aircraftInfo, err := service.AircraftDatabase.GetAircraftInfo(portfolioAircraftVO.AircraftID)
	if err != nil {
		return err
	}
	if aircraftInfo == nil {
		return common.NewErrorCode(common.ErrorItemNotFoundAircraftCode)
	}
	return service.PortfolioAircraftDatabase.CreatePortfolioAircraft(&databases.CreatePortfolioAircraftInfo{
		PortfolioID: portfolioAircraftVO.PortfolioID,
		AircraftID:  portfolioAircraftVO.AircraftID,
	})
}

// GetPortfolio get a portfolio
func (service *PortfolioDefaultService) GetPortfolio(portfolioID string) (*PortfolioVO, error) {
	portfolioInfo, err := service.PortfolioDatabase.GetPortfolioInfo(portfolioID)
	if err != nil {
		return nil, err
	}
	// Since it's a unique item, we can afford doing another GET request to get aircraft IDs everytime
	portfolioAircraftInfoList, err := service.PortfolioAircraftDatabase.GetPortfolioAircraftInfos(portfolioID)
	if err != nil {
		return nil, err
	}
	aircraftVOList := make([]*AircraftVO, len(portfolioAircraftInfoList))
	for i, portfolioAircraftInfo := range portfolioAircraftInfoList {
		aircraftVOList[i] = &AircraftVO{
			ID: portfolioAircraftInfo.AircraftID,
		}
	}
	portfolioVO := convertPortfolioInfoToVO(portfolioInfo)
	if portfolioVO != nil {
		portfolioVO.AircraftVOList = aircraftVOList
	}
	return portfolioVO, nil
}

// GetPortfolioList get the portfolio list
func (service *PortfolioDefaultService) GetPortfolioList() ([]*PortfolioVO, error) {
	portfolioInfoList, err := service.PortfolioDatabase.GetPortfolioInfoList()
	if err != nil {
		return []*PortfolioVO{}, err
	}
	portfolioVOList := make([]*PortfolioVO, len(portfolioInfoList))
	for i, portfolioInfo := range portfolioInfoList {
		portfolioVOList[i] = convertPortfolioInfoToVO(portfolioInfo)
	}
	return portfolioVOList, nil
}

// DeletePortfolio delete a porfolio
func (service *PortfolioDefaultService) DeletePortfolio(portfolioID string) error {
	err := service.PortfolioDatabase.DeletePortfolio(portfolioID)
	if err != nil {
		return err
	}
	portfolioAircraftInfoList, err := service.PortfolioAircraftDatabase.GetPortfolioAircraftInfos(portfolioID)
	if err != nil {
		return err
	}
	if len(portfolioAircraftInfoList) == 0 {
		return nil
	}
	return service.PortfolioAircraftDatabase.DeletePortfolioAircraftsByPortfolioID(portfolioID)
}

// DeletePortfolioAircraft delete a portfolio aircraft link
func (service *PortfolioDefaultService) DeletePortfolioAircraft(portfolioID, aircraftID string) error {
	return service.PortfolioAircraftDatabase.DeletePortfolioAircraft(portfolioID, aircraftID)
}

func convertPortfolioInfoToVO(portfolioInfo *databases.PortfolioInfo) *PortfolioVO {
	if portfolioInfo == nil {
		return nil
	}
	return &PortfolioVO{
		ID:   portfolioInfo.ID,
		Name: portfolioInfo.Name,
	}
}
