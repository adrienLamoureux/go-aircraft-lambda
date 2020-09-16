package databases

// IFlightDatabase is the Flight Database contract
type IFlightDatabase interface {
	// CreateFlight create a flight
	CreateFlight(flightInfo *CreateFlightInfo) error

	// GetFlightInfo get a flight
	GetFlightInfo(flightID string) (*FlightInfo, error)

	// GetFlightInfosByAircraftID get a flight list of an Aircraft
	GetFlightInfosByAircraftID(aircraftID string) ([]*FlightInfo, error)
}
