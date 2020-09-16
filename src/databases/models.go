package databases

// AircraftInfo AircraftInfo
type AircraftInfo struct {
	ID       string
	Model    string
	CreateTm int64
	UpdateTm int64
}

// CreateAircraftInfo CreateAircraftInfo
type CreateAircraftInfo struct {
	ID    string
	Model string
}

// AircraftModelInfo AircraftModelInfo
type AircraftModelInfo struct {
	ID       string
	CreateTm int64
	UpdateTm int64
}

// CreateAircraftModelInfo CreateAircraftModelInfo
type CreateAircraftModelInfo struct {
	ID string
}

// PortfolioInfo PortfolioInfo
type PortfolioInfo struct {
	ID       string
	Name     string
	CreateTm int64
	UpdateTm int64
}

// CreatePortfolioInfo CreatePortfolioInfo
type CreatePortfolioInfo struct {
	ID   string
	Name string
}

// PortfolioAircraftInfo PortfolioAircraftInfo
type PortfolioAircraftInfo struct {
	PortfolioID string
	AircraftID  string
	CreateTm    int64
	UpdateTm    int64
}

// CreatePortfolioAircraftInfo CreatePortfolioAircraftInfo
type CreatePortfolioAircraftInfo struct {
	PortfolioID string
	AircraftID  string
}

// FlightInfo FlightInfo
type FlightInfo struct {
	FlightID         string
	AircraftID       string
	DepartureAirport string
	DepartureTime    int64
	ArrivalAirport   string
	ArrivalTime      int64
	CreateTm         int64
	UpdateTm         int64
}

// CreateFlightInfo CreateFlightInfo
type CreateFlightInfo struct {
	FlightID         string
	AircraftID       string
	DepartureAirport string
	DepartureTime    int64
	ArrivalAirport   string
	ArrivalTime      int64
}

// AirportInfo AirportInfo
type AirportInfo struct {
	ID       string
	CreateTm int64
	UpdateTm int64
}

// CreateAirportInfo CreateAirportInfo
type CreateAirportInfo struct {
	ID string
}
