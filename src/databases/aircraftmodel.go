package databases

// IAircraftModelDatabase is the Aircraft Model Database contract
type IAircraftModelDatabase interface {
	// CreateAircraftModel create an aircraft model
	CreateAircraftModel(aircraftModelInfo *CreateAircraftModelInfo) error

	// GetAircraftModelInfo get an aircraft model
	GetAircraftModelInfo(aircraftModelID string) (*AircraftModelInfo, error)

	// GetAircraftModelInfoList get an aircraft model list
	GetAircraftModelInfoList() ([]*AircraftModelInfo, error)
}
