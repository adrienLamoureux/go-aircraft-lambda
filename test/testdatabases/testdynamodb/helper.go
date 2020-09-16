package testdynamodb

import (
	"fmt"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases/dynamodatabase"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDBHelper provide methods to manipulate the DynamoDB specifically for the tests
type DynamoDBHelper struct {
	TableSuffix                string
	Client                     *dynamodb.DynamoDB
	portfolioTableName         *string
	portfolioNameTableName     *string
	aircraftTableName          *string
	aircraftModelTableName     *string
	portfolioAircraftTableName *string
	flightTableName            *string
	airportTableName           *string
}

// SetPortfolioTableName SetPortfolioTableName
func (dynamoHelper *DynamoDBHelper) SetPortfolioTableName(tableName string) {
	dynamoHelper.portfolioTableName = dynamoHelper.buildTableName(tableName)
}

// GetPortfolioTableName GetPortfolioTableName
func (dynamoHelper *DynamoDBHelper) GetPortfolioTableName() *string {
	return dynamoHelper.portfolioTableName
}

// SetPortfolioNameTableName SetPortfolioNameTableName
func (dynamoHelper *DynamoDBHelper) SetPortfolioNameTableName(tableName string) {
	dynamoHelper.portfolioNameTableName = dynamoHelper.buildTableName(tableName)
}

// GetPortfolioNameTableName GetPortfolioNameTableName
func (dynamoHelper *DynamoDBHelper) GetPortfolioNameTableName() *string {
	return dynamoHelper.portfolioNameTableName
}

// SetAircraftTableName SetAircraftTableName
func (dynamoHelper *DynamoDBHelper) SetAircraftTableName(tableName string) {
	dynamoHelper.aircraftTableName = dynamoHelper.buildTableName(tableName)
}

// GetAircraftTableName GetAircraftTableName
func (dynamoHelper *DynamoDBHelper) GetAircraftTableName() *string {
	return dynamoHelper.aircraftTableName
}

// SetAircraftModelTableName SetAircraftModelTableName
func (dynamoHelper *DynamoDBHelper) SetAircraftModelTableName(tableName string) {
	dynamoHelper.aircraftModelTableName = dynamoHelper.buildTableName(tableName)
}

// GetAircraftModelTableName GetAircraftModelTableName
func (dynamoHelper *DynamoDBHelper) GetAircraftModelTableName() *string {
	return dynamoHelper.aircraftModelTableName
}

// SetPortfolioAircraftTableName SetPortfolioAircraftTableName
func (dynamoHelper *DynamoDBHelper) SetPortfolioAircraftTableName(tableName string) {
	dynamoHelper.portfolioAircraftTableName = dynamoHelper.buildTableName(tableName)
}

// GetPortfolioAircraftTableName GetPortfolioAircraftTableName
func (dynamoHelper *DynamoDBHelper) GetPortfolioAircraftTableName() *string {
	return dynamoHelper.portfolioAircraftTableName
}

// SetFlightTableName SetFlightTableName
func (dynamoHelper *DynamoDBHelper) SetFlightTableName(tableName string) {
	dynamoHelper.flightTableName = dynamoHelper.buildTableName(tableName)
}

// GetFlightTableName GetFlightTableName
func (dynamoHelper *DynamoDBHelper) GetFlightTableName() *string {
	return dynamoHelper.flightTableName
}

// SetAirportTableName SetAirportTableName
func (dynamoHelper *DynamoDBHelper) SetAirportTableName(tableName string) {
	dynamoHelper.airportTableName = dynamoHelper.buildTableName(tableName)
}

// GetAirportTableName GetAirportTableName
func (dynamoHelper *DynamoDBHelper) GetAirportTableName() *string {
	return dynamoHelper.airportTableName
}

// CreatePortfolioTable create the portfolio table
// Create also the portfolio name table
func (dynamoHelper *DynamoDBHelper) CreatePortfolioTable() error {
	err := dynamodatabase.CreatePortfolioTable(dynamoHelper.portfolioTableName)
	if err != nil {
		return err
	}
	return dynamodatabase.CreatePortfolioNameTable(dynamoHelper.portfolioNameTableName)
}

// DeletePortfolioTable delete the portfolio table
// Delete also the portfolio name table
func (dynamoHelper *DynamoDBHelper) DeletePortfolioTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.portfolioTableName,
	})
	if err != nil {
		return err
	}
	_, err = dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.portfolioNameTableName,
	})
	return err
}

// CreateAircraftTable create the aircraft table
func (dynamoHelper *DynamoDBHelper) CreateAircraftTable() error {
	return dynamodatabase.CreateAircraftTable(dynamoHelper.aircraftTableName)
}

// DeleteAircraftTable delete the aircraft table
func (dynamoHelper *DynamoDBHelper) DeleteAircraftTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.aircraftTableName,
	})
	return err
}

// CreateAircraftModelTable create the aircraft model table
func (dynamoHelper *DynamoDBHelper) CreateAircraftModelTable() error {
	return dynamodatabase.CreateAircraftModelTable(dynamoHelper.aircraftModelTableName)
}

// DeleteAircraftModelTable delete the aircraft model table
func (dynamoHelper *DynamoDBHelper) DeleteAircraftModelTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.aircraftModelTableName,
	})
	return err
}

// CreatePortfolioAircraftTable create the portfolio aircraft table
func (dynamoHelper *DynamoDBHelper) CreatePortfolioAircraftTable() error {
	return dynamodatabase.CreatePortfolioAircraftTable(dynamoHelper.portfolioAircraftTableName)
}

// DeletePortfolioAircraftTable delete the portfolio aircraft table
func (dynamoHelper *DynamoDBHelper) DeletePortfolioAircraftTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.portfolioAircraftTableName,
	})
	return err
}

// CreateFlightTable create the flight table
func (dynamoHelper *DynamoDBHelper) CreateFlightTable() error {
	return dynamodatabase.CreateFlightTable(dynamoHelper.flightTableName)
}

// DeleteFlightTable delete the flight table
func (dynamoHelper *DynamoDBHelper) DeleteFlightTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.flightTableName,
	})
	return err
}

// CreateAirportTable create the airport table
func (dynamoHelper *DynamoDBHelper) CreateAirportTable() error {
	return dynamodatabase.CreateAirportTable(dynamoHelper.airportTableName)
}

// DeleteAirportTable delete the airport table
func (dynamoHelper *DynamoDBHelper) DeleteAirportTable() error {
	_, err := dynamoHelper.Client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: dynamoHelper.airportTableName,
	})
	return err
}

func (dynamoHelper *DynamoDBHelper) buildTableName(tableName string) *string {
	return aws.String(fmt.Sprintf("%s_%s", tableName, dynamoHelper.TableSuffix))
}
