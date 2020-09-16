package databases

// IPortfolioDatabase is the Portfolio Database contract
type IPortfolioDatabase interface {
	// CreatePortfolio create a portfolio
	CreatePortfolio(portfolioInfo *CreatePortfolioInfo) error

	// GetPortfolioInfo get a portfolio
	GetPortfolioInfo(portfolioID string) (*PortfolioInfo, error)

	// GetPortfolioInfoList get a portfolio list
	GetPortfolioInfoList() ([]*PortfolioInfo, error)

	// DeletePortfolio delete a portfolio and the related portfolio name
	DeletePortfolio(portfolioID string) error
}
