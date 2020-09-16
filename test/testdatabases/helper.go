package testdatabases

// IDatabaseHelper provide methods to manipulate the DB specifically for the tests
type IDatabaseHelper interface {
	// CreatePortfolioTable create the portfolio table
	CreatePortfolioTable() error

	// DeletePortfolioTable delete the portfolio table
	DeletePortfolioTable() error

	// CreateAircraftTable create the aircraft table
	CreateAircraftTable() error

	// DeleteAircraftTable delete the aircraft table
	DeleteAircraftTable() error

	// CreateAircraftModelTable create the aircraft model table
	CreateAircraftModelTable() error

	// DeleteAircraftModelTable delete the aircraft model table
	DeleteAircraftModelTable() error

	// CreatePortfolioAircraftTable create the portfolio aircraft table
	CreatePortfolioAircraftTable() error

	// DeletePortfolioAircraftTable delete the portfolio aircraft table
	DeletePortfolioAircraftTable() error

	// CreateFlightTable create the flight table
	CreateFlightTable() error

	// DeleteFlightTable delete the flight table
	DeleteFlightTable() error

	// CreateAirportTable create the airport table
	CreateAirportTable() error

	// DeleteAirportTable delete the airport table
	DeleteAirportTable() error
}
