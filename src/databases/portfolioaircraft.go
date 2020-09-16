package databases

// IPortfolioAircraftDatabase is the Portfolio Aircraft Database contract
type IPortfolioAircraftDatabase interface {
	// CreatePortfolioAircraft create a portfolio aircraft
	CreatePortfolioAircraft(portfolioAircraftInfo *CreatePortfolioAircraftInfo) error

	// GetPortfolioAircraftInfo get a portfolio aircraft
	GetPortfolioAircraftInfo(portfolioAircraftPortfolioID, portfolioAircraftAircraftID string) (*PortfolioAircraftInfo, error)

	// GetPortfolioAircraftInfos get a portfolio aircraft list of a portfolio
	GetPortfolioAircraftInfos(portfolioAircraftPortfolioID string) ([]*PortfolioAircraftInfo, error)

	// GetPortfolioAircraftInfosByAircraftID get a portfolio aircraft list of an aircraft
	GetPortfolioAircraftInfosByAircraftID(portfolioAircraftAircraftID string) ([]*PortfolioAircraftInfo, error)

	// DeletePortfolioAircraft delete a portfolio aircraft
	DeletePortfolioAircraft(portfolioAircraftPortfolioID, portfolioAircraftAircraftID string) error

	// DeletePortfolioAircraftsByPortfolioID delete a portfolio aircraft list of a portfolio
	DeletePortfolioAircraftsByPortfolioID(portfolioAircraftPortfolioID string) error
}
