package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	portfolioTestDefaultID   = "UUID-1"
	portfolioTestDefaultName = "p1"
)

// PortfolioTestSuite is the test suite on any DB
type PortfolioTestSuite struct {
	suite.Suite
	Database databases.IPortfolioDatabase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *PortfolioTestSuite) SetupTest() {
	err := suite.Helper.CreatePortfolioTable()
	if err != nil {
		panic(err)
	}
	err = suite.createPortfolio()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *PortfolioTestSuite) TearDownTest() {
	err := suite.Helper.DeletePortfolioTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreatePortfolioWhenExist TestShouldNotCreatePortfolioWhenExist
func (suite *PortfolioTestSuite) TestShouldNotCreatePortfolioWhenExist() {
	err := suite.Database.CreatePortfolio(&databases.CreatePortfolioInfo{
		ID:   portfolioTestDefaultID,
		Name: portfolioTestDefaultName + "a",
	})
	assert.Error(suite.T(), err)
	err = suite.Database.CreatePortfolio(&databases.CreatePortfolioInfo{
		ID:   portfolioTestDefaultID + "a",
		Name: portfolioTestDefaultName,
	})
	assert.Error(suite.T(), err)
}

// TestShouldGetPortfolioWhenExist TestShouldGetPortfolioWhenExist
func (suite *PortfolioTestSuite) TestShouldGetPortfolioWhenExist() {
	portfolioInfo, err := suite.Database.GetPortfolioInfo(portfolioTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), portfolioInfo)
	assert.Equal(suite.T(), portfolioTestDefaultID, portfolioInfo.ID)
	assert.Equal(suite.T(), portfolioTestDefaultName, portfolioInfo.Name)
}

// TestShouldNotGetPortfolioWhenNotExist TestShouldNotGetPortfolioWhenNotExist
func (suite *PortfolioTestSuite) TestShouldNotGetPortfolioWhenNotExist() {
	portfolioInfo, err := suite.Database.GetPortfolioInfo(portfolioTestDefaultID + "a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), portfolioInfo)
}

// TestShouldGetPortfolioInfoList TestShouldGetPortfolioInfoList
func (suite *PortfolioTestSuite) TestShouldGetPortfolioInfoList() {
	portfolioInfoList, err := suite.Database.GetPortfolioInfoList()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), portfolioInfoList)
}

// TestShouldDeletePortfolioWhenExist TestShouldDeletePortfolioWhenExist
func (suite *PortfolioTestSuite) TestShouldDeletePortfolioWhenExist() {
	err := suite.Database.DeletePortfolio(portfolioTestDefaultID)
	assert.NoError(suite.T(), err)
	portfolioInfo, err := suite.Database.GetPortfolioInfo(portfolioTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), portfolioInfo)
}

// TestShouldNotDeletePortfolioWhenNotExist TestShouldNotDeletePortfolioWhenNotExist
func (suite *PortfolioTestSuite) TestShouldNotDeletePortfolioWhenNotExist() {
	assert.Error(suite.T(), suite.Database.DeletePortfolio(portfolioTestDefaultID+"a"))
	portfolioInfo, err := suite.Database.GetPortfolioInfo(portfolioTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), portfolioInfo)
}

func (suite *PortfolioTestSuite) createPortfolio() error {
	return suite.Database.CreatePortfolio(&databases.CreatePortfolioInfo{
		ID:   portfolioTestDefaultID,
		Name: portfolioTestDefaultName,
	})
}
