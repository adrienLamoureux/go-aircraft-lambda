package databases

// IAircraftDabatase is the Aircraft Database contract
type IAircraftDabatase interface {
	// CreateAircraft create an aircraft
	CreateAircraft(aircraftInfo *CreateAircraftInfo) error

	// GetAircraftInfo get an aircraft
	GetAircraftInfo(aircraftID string) (*AircraftInfo, error)

	// GetAircraftInfoList get the aircraft list
	GetAircraftInfoList() ([]*AircraftInfo, error)
}
