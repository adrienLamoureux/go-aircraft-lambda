package databases

// IAirportDatabase is the Airport Database contract
type IAirportDatabase interface {
	// CreateAirport create an airport
	CreateAirport(airportInfo *CreateAirportInfo) error

	// GetAirportInfo get an airport
	GetAirportInfo(airportID string) (*AirportInfo, error)
}
