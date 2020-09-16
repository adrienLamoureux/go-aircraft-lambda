package dynamodatabase

import "github.com/adrienLamoureux/go-aircraft-lambda/src/databases"

type aircraftTableData struct {
	ID       string `json:"id"`
	Model    string `json:"model"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (aircraftData *aircraftTableData) toAbstract() *databases.AircraftInfo {
	if aircraftData == nil {
		return nil
	}
	return &databases.AircraftInfo{
		ID:       aircraftData.ID,
		Model:    aircraftData.Model,
		CreateTm: aircraftData.CreateTm,
		UpdateTm: aircraftData.UpdateTm,
	}
}

type aircraftModelTableData struct {
	ID       string `json:"id"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (aircraftModelData *aircraftModelTableData) toAbstract() *databases.AircraftModelInfo {
	return &databases.AircraftModelInfo{
		ID:       aircraftModelData.ID,
		CreateTm: aircraftModelData.CreateTm,
		UpdateTm: aircraftModelData.UpdateTm,
	}
}

type portfolioTableData struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (portfolioData *portfolioTableData) toAbstract() *databases.PortfolioInfo {
	return &databases.PortfolioInfo{
		ID:       portfolioData.ID,
		Name:     portfolioData.Name,
		CreateTm: portfolioData.CreateTm,
		UpdateTm: portfolioData.UpdateTm,
	}
}

type portfolioNameTableData struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

type portfolioAircraftTableData struct {
	PortfolioID string `json:"portfolioId"`
	AircraftID  string `json:"aircraftId"`
	CreateTm    int64  `json:"createTm"`
	UpdateTm    int64  `json:"updateTm"`
}

func (portfolioAircraftData *portfolioAircraftTableData) toAbstract() *databases.PortfolioAircraftInfo {
	return &databases.PortfolioAircraftInfo{
		PortfolioID: portfolioAircraftData.PortfolioID,
		AircraftID:  portfolioAircraftData.AircraftID,
		CreateTm:    portfolioAircraftData.CreateTm,
		UpdateTm:    portfolioAircraftData.UpdateTm,
	}
}

type flightTableData struct {
	FlightID         string `json:"flightId"`
	AircraftID       string `json:"aircraftId"`
	DepartureAirport string `json:"departureAirport"`
	DepartureTime    int64  `json:"departureTm"`
	ArrivalAirport   string `json:"arrivalAirport"`
	ArrivalTime      int64  `json:"arrivalTm"`
	CreateTm         int64  `json:"createTm"`
	UpdateTm         int64  `json:"updateTm"`
}

func (flightData *flightTableData) toAbstract() *databases.FlightInfo {
	return &databases.FlightInfo{
		FlightID:         flightData.FlightID,
		AircraftID:       flightData.AircraftID,
		DepartureAirport: flightData.DepartureAirport,
		DepartureTime:    flightData.DepartureTime,
		ArrivalAirport:   flightData.ArrivalAirport,
		ArrivalTime:      flightData.ArrivalTime,
		CreateTm:         flightData.CreateTm,
		UpdateTm:         flightData.UpdateTm,
	}
}

type airportTableData struct {
	ID       string `json:"id"`
	CreateTm int64  `json:"createTm"`
	UpdateTm int64  `json:"updateTm"`
}

func (airportData *airportTableData) toAbstract() *databases.AirportInfo {
	return &databases.AirportInfo{
		ID:       airportData.ID,
		CreateTm: airportData.CreateTm,
		UpdateTm: airportData.UpdateTm,
	}
}
