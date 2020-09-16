package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	portfolioAircraftTestDefaultPortfolioID = "pID"
	portfolioAircraftTestDefaultAircraftID  = "aID"
)

// PortfolioAircraftTestSuite is the test suite on any DB
type PortfolioAircraftTestSuite struct {
	suite.Suite
	Database databases.IPortfolioAircraftDatabase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *PortfolioAircraftTestSuite) SetupTest() {
	err := suite.Helper.CreatePortfolioAircraftTable()
	if err != nil {
		panic(err)
	}
	err = suite.createPortfolioAircraft()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *PortfolioAircraftTestSuite) TearDownTest() {
	err := suite.Helper.DeletePortfolioAircraftTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreatePortfolioAircraftWhenExist TestShouldNotCreatePortfolioAircraftWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotCreatePortfolioAircraftWhenExist() {
	assert.Error(suite.T(), suite.createPortfolioAircraft())
}

// TestShouldGetPortfolioAircraftWhenExist TestShouldGetPortfolioAircraftWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldGetPortfolioAircraftWhenExist() {
	portfolioAircraftInfo, err := suite.Database.GetPortfolioAircraftInfo(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), portfolioAircraftInfo)
	assert.Equal(suite.T(), portfolioAircraftTestDefaultPortfolioID, portfolioAircraftInfo.PortfolioID)
	assert.Equal(suite.T(), portfolioAircraftTestDefaultAircraftID, portfolioAircraftInfo.AircraftID)
}

// TestShouldNotGetPortfolioAircraftWhenNotExist TestShouldNotGetPortfolioAircraftWhenNotExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotGetPortfolioAircraftWhenNotExist() {
	portfolioAircraftInfo, err := suite.Database.GetPortfolioAircraftInfo(portfolioAircraftTestDefaultPortfolioID+"a", portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), portfolioAircraftInfo)
	portfolioAircraftInfo, err = suite.Database.GetPortfolioAircraftInfo(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID+"a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), portfolioAircraftInfo)
}

// TestShouldGetPortfolioAircraftInfosWhenExist TestShouldGetPortfolioAircraftInfosWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldGetPortfolioAircraftInfosWhenExist() {
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfos(portfolioAircraftTestDefaultPortfolioID)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), portfolioAircraftInfoList)
}

// TestShouldNotGetPortfolioAircraftInfosWhenNotExist TestShouldNotGetPortfolioAircraftInfosWhenNotExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotGetPortfolioAircraftInfosWhenNotExist() {
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfos(portfolioAircraftTestDefaultPortfolioID + "a")
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), portfolioAircraftInfoList)
}

// TestShouldGetPortfolioAircraftInfosByAircraftIDWhenExist TestShouldGetPortfolioAircraftInfosByAircraftIDWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldGetPortfolioAircraftInfosByAircraftIDWhenExist() {
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfosByAircraftID(portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), portfolioAircraftInfoList)
}

// TestShouldNotGetPortfolioAircraftInfosByAircraftIDWhenNotExist TestShouldNotGetPortfolioAircraftInfosByAircraftIDWhenNotExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotGetPortfolioAircraftInfosByAircraftIDWhenNotExist() {
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfosByAircraftID(portfolioAircraftTestDefaultAircraftID + "a")
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), portfolioAircraftInfoList)
}

// TestShouldDeletePortfolioAircraftWhenExist  TestShouldDeletePortfolioAircraftWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldDeletePortfolioAircraftWhenExist() {
	err := suite.Database.DeletePortfolioAircraft(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	portfolioAircraftInfo, err := suite.Database.GetPortfolioAircraftInfo(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), portfolioAircraftInfo)
}

// TestShouldNotDeletePortfolioAircraftWhenNotExist TestShouldNotDeletePortfolioAircraftWhenNotExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotDeletePortfolioAircraftWhenNotExist() {
	err := suite.Database.DeletePortfolioAircraft(portfolioAircraftTestDefaultPortfolioID+"a", portfolioAircraftTestDefaultAircraftID)
	assert.Error(suite.T(), err)
	err = suite.Database.DeletePortfolioAircraft(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID+"a")
	assert.Error(suite.T(), err)
	portfolioAircraftInfo, err := suite.Database.GetPortfolioAircraftInfo(portfolioAircraftTestDefaultPortfolioID, portfolioAircraftTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), portfolioAircraftInfo)
}

// TestShouldDeletePortfolioAircraftsByPortfolioIDWhenExist TestShouldDeletePortfolioAircraftsByPortfolioIDWhenExist
func (suite *PortfolioAircraftTestSuite) TestShouldDeletePortfolioAircraftsByPortfolioIDWhenExist() {
	err := suite.Database.DeletePortfolioAircraftsByPortfolioID(portfolioAircraftTestDefaultPortfolioID)
	assert.NoError(suite.T(), err)
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfos(portfolioAircraftTestDefaultPortfolioID)
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), portfolioAircraftInfoList)
}

// TestShouldNotDeletePortfolioAircraftsByPortfolioIDWhenNotExist TestShouldNotDeletePortfolioAircraftsByPortfolioIDWhenNotExist
func (suite *PortfolioAircraftTestSuite) TestShouldNotDeletePortfolioAircraftsByPortfolioIDWhenNotExist() {
	err := suite.Database.DeletePortfolioAircraftsByPortfolioID(portfolioAircraftTestDefaultPortfolioID + "a")
	assert.Error(suite.T(), err)
	portfolioAircraftInfoList, err := suite.Database.GetPortfolioAircraftInfos(portfolioAircraftTestDefaultPortfolioID)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), portfolioAircraftInfoList)
}

func (suite *PortfolioAircraftTestSuite) createPortfolioAircraft() error {
	return suite.Database.CreatePortfolioAircraft(&databases.CreatePortfolioAircraftInfo{
		PortfolioID: portfolioAircraftTestDefaultPortfolioID,
		AircraftID:  portfolioAircraftTestDefaultAircraftID,
	})
}
