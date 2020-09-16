package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	aircraftModelTestDefaultID = "A320-200"
)

// AircraftModelTestSuite is the test suite on any DB
type AircraftModelTestSuite struct {
	suite.Suite
	Database databases.IAircraftModelDatabase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *AircraftModelTestSuite) SetupTest() {
	err := suite.Helper.CreateAircraftModelTable()
	if err != nil {
		panic(err)
	}
	err = suite.createAircraftModel()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *AircraftModelTestSuite) TearDownTest() {
	err := suite.Helper.DeleteAircraftModelTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreateAircraftModelWhenExist TestShouldNotCreateAircraftModelWhenExist
func (suite *AircraftModelTestSuite) TestShouldNotCreateAircraftModelWhenExist() {
	assert.Error(suite.T(), suite.createAircraftModel())
}

// TestShouldGetAircraftModelInfoList TestShouldGetAircraftModelInfoList
func (suite *AircraftModelTestSuite) TestShouldGetAircraftModelInfoList() {
	aircraftModelInfoList, err := suite.Database.GetAircraftModelInfoList()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), aircraftModelInfoList)
}

// TestShouldGetAircraftWhenExist TestShouldGetAircraftWhenExist
func (suite *AircraftModelTestSuite) TestShouldGetAircraftWhenExist() {
	aircraftModelInfo, err := suite.Database.GetAircraftModelInfo(aircraftModelTestDefaultID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), aircraftModelInfo)
	assert.Equal(suite.T(), aircraftModelTestDefaultID, aircraftModelInfo.ID)
}

// TestShouldNotGetAircraftWhenNotExist TestShouldNotGetAircraftWhenNotExist
func (suite *AircraftModelTestSuite) TestShouldNotGetAircraftWhenNotExist() {
	aircraftModelInfo, err := suite.Database.GetAircraftModelInfo(aircraftModelTestDefaultID + "a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), aircraftModelInfo)
}

func (suite *AircraftModelTestSuite) createAircraftModel() error {
	return suite.Database.CreateAircraftModel(&databases.CreateAircraftModelInfo{
		ID: aircraftModelTestDefaultID,
	})
}
