package testdatabases

import (
	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	flightTestDefaultFlightID         = "FR1454"
	flightTestDefaultAircraftID       = "ZS-GAO"
	flightTestDefaultDepartureTime    = int64(1595931449)
	flightTestDefaultDepartureAirport = "DUB"
	flightTestDefaultArrivalTime      = int64(1595939700)
	flightTestDefaultArrivalAirport   = "MXP"
)

// FlightTestSuite is the test suite on any DB
type FlightTestSuite struct {
	suite.Suite
	Database databases.IFlightDatabase
	Helper   IDatabaseHelper
}

// SetupTest SetupTest
func (suite *FlightTestSuite) SetupTest() {
	err := suite.Helper.CreateFlightTable()
	if err != nil {
		panic(err)
	}
	err = suite.createFlight()
	if err != nil {
		panic(err)
	}
}

// TearDownTest TearDownTest
func (suite *FlightTestSuite) TearDownTest() {
	err := suite.Helper.DeleteFlightTable()
	if err != nil {
		panic(err)
	}
}

// TestShouldNotCreateFlightWhenExist TestShouldNotCreateFlightWhenExist
func (suite *FlightTestSuite) TestShouldNotCreateFlightWhenExist() {
	assert.Error(suite.T(), suite.createFlight())
}

// TestShouldGetFlightWhenExist TestShouldGetFlightWhenExist
func (suite *FlightTestSuite) TestShouldGetFlightWhenExist() {
	flightInfo, err := suite.Database.GetFlightInfo(flightTestDefaultFlightID)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), flightInfo)
	assert.Equal(suite.T(), flightTestDefaultFlightID, flightInfo.FlightID)
	assert.Equal(suite.T(), flightTestDefaultAircraftID, flightInfo.AircraftID)
	assert.Equal(suite.T(), flightTestDefaultDepartureAirport, flightInfo.DepartureAirport)
	assert.Equal(suite.T(), flightTestDefaultDepartureTime, flightInfo.DepartureTime)
	assert.Equal(suite.T(), flightTestDefaultArrivalAirport, flightInfo.ArrivalAirport)
	assert.Equal(suite.T(), flightTestDefaultArrivalTime, flightInfo.ArrivalTime)
}

// TestShouldNotGetFlightWhenNotExist TestShouldNotGetFlightWhenNotExist
func (suite *FlightTestSuite) TestShouldNotGetFlightWhenNotExist() {
	flightInfo, err := suite.Database.GetFlightInfo(flightTestDefaultFlightID + "a")
	assert.NoError(suite.T(), err)
	assert.Nil(suite.T(), flightInfo)
}

// TestShouldGetFlightInfosByAircraftIDWhenExist TestShouldGetFlightInfosByAircraftIDWhenExist
func (suite *FlightTestSuite) TestShouldGetFlightInfosByAircraftIDWhenExist() {
	flightInfoList, err := suite.Database.GetFlightInfosByAircraftID(flightTestDefaultAircraftID)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), flightInfoList)
}

// TestShouldGetFlightInfosByAircraftIDWhenNotExist TestShouldGetFlightInfosByAircraftIDWhenNotExist
func (suite *FlightTestSuite) TestShouldGetFlightInfosByAircraftIDWhenNotExist() {
	flightInfoList, err := suite.Database.GetFlightInfosByAircraftID(flightTestDefaultAircraftID + "a")
	assert.NoError(suite.T(), err)
	assert.Empty(suite.T(), flightInfoList)
}

func (suite *FlightTestSuite) createFlight() error {
	return suite.Database.CreateFlight(&databases.CreateFlightInfo{
		FlightID:         flightTestDefaultFlightID,
		AircraftID:       flightTestDefaultAircraftID,
		DepartureAirport: flightTestDefaultDepartureAirport,
		DepartureTime:    flightTestDefaultDepartureTime,
		ArrivalAirport:   flightTestDefaultArrivalAirport,
		ArrivalTime:      flightTestDefaultArrivalTime,
	})
}
