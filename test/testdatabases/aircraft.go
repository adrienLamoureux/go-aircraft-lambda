package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	aircraftTestDefaultID    = "ZS-GAO"
	aircraftTestDefaultModel = "A320-200"
)

// AircraftTestSuite is the test suite on any DB
type AircraftTestSuite struct {
	suite.Suite
	Database databases.IAircraftDabatase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *AircraftTestSuite) SetupTest() {
	err := suite.Helper.CreateAircraftTable()
	if err != nil {
		panic(err)
	}
	err = suite.createAircraft()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *AircraftTestSuite) TearDownTest() {
	err := suite.Helper.DeleteAircraftTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreateAircraftWhenExist TestShouldNotCreateAircraftWhenExist
func (suite *AircraftTestSuite) TestShouldNotCreateAircraftWhenExist() {
	assert.Error(suite.T(), suite.createAircraft())
}

// TestShouldGetAircraftInfoList TestShouldGetAircraftInfoList
func (suite *AircraftTestSuite) TestShouldGetAircraftInfoList() {
	aircraftInfoList, err := suite.Database.GetAircraftInfoList()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), aircraftInfoList)
}

// TestShouldGetAircraftWhenExist TestShouldGetAircraftWhenExist
func (suite *AircraftTestSuite) TestShouldGetAircraftWhenExist() {
	aircraftInfo, err := suite.Database.GetAircraftInfo(aircraftTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), aircraftInfo)
	assert.Equal(suite.T(), aircraftTestDefaultID, aircraftInfo.ID)
	assert.Equal(suite.T(), aircraftTestDefaultModel, aircraftInfo.Model)
}

// TestShouldNotGetAircraftWhenNotExist TestShouldNotGetAircraftWhenNotExist
func (suite *AircraftTestSuite) TestShouldNotGetAircraftWhenNotExist() {
	aircraftInfo, err := suite.Database.GetAircraftInfo(aircraftTestDefaultID + "a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), aircraftInfo)
}

func (suite *AircraftTestSuite) createAircraft() error {
	return suite.Database.CreateAircraft(&databases.CreateAircraftInfo{
		ID:    aircraftTestDefaultID,
		Model: aircraftTestDefaultModel,
	})
}
