package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	airportTestDefaultID = "DUB"
)

// AirportTestSuite is the test suite on any DB
type AirportTestSuite struct {
	suite.Suite
	Database databases.IAirportDatabase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *AirportTestSuite) SetupTest() {
	err := suite.Helper.CreateAirportTable()
	if err != nil {
		panic(err)
	}
	err = suite.createAirport()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *AirportTestSuite) TearDownTest() {
	err := suite.Helper.DeleteAirportTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreateAirportWhenExist TestShouldNotCreateAirportWhenExist
func (suite *AirportTestSuite) TestShouldNotCreateAirportWhenExist() {
	assert.Error(suite.T(), suite.createAirport())
}

// TestShouldGetAirportWhenExist TestShouldGetAirportWhenExist
func (suite *AirportTestSuite) TestShouldGetAirportWhenExist() {
	airportInfo, err := suite.Database.GetAirportInfo(airportTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), airportInfo)
	assert.Equal(suite.T(), airportTestDefaultID, airportInfo.ID)
}

// TestShouldNotGetAirportWhenNotExist TestShouldNotGetAirportWhenNotExist
func (suite *AirportTestSuite) TestShouldNotGetAirportWhenNotExist() {
	airportInfo, err := suite.Database.GetAirportInfo(airportTestDefaultID + "a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), airportInfo)
}

func (suite *AirportTestSuite) createAirport() error {
	return suite.Database.CreateAirport(&databases.CreateAirportInfo{
		ID: airportTestDefaultID,
	})
}
