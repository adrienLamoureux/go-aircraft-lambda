package dynamodatabase

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	readUnitsDefault  = 10
	writeUnitsDefault = 10

	portfolioTableHashKeyName                      = "id"
	portfolioNameTableHashKeyName                  = "name"
	aircraftTableHashKeyName                       = "id"
	aircraftModelTableHashKeyName                  = "id"
	portfolioAircraftTableHashKeyName              = "portfolioId"
	portfolioAircraftTableRangeKeyName             = "aircraftId"
	portfolioAircraftTableAircraftIndexName        = "aircraftIndex"
	portfolioAircraftTableAircraftIndexHashKeyName = "aircraftId"
	flightTableHashKeyName                         = "flightId"
	flightTableAircraftIndexName                   = "aircraftIndex"
	flightTableAircraftIndexHashKeyName            = "aircraftId"
	airportTableHashKeyName                        = "id"
)

var (
	portfolioTableName         = "portfolios"
	portfolioNameTableName     = "portfolio_names"
	aircraftTableName          = "aircrafts"
	aircraftModelTableName     = "aircraft_models"
	portfolioAircraftTableName = "portfolio_aircrafts"
	flightTableName            = "flights"
	airportTableName           = "airports"
)

// GetPortfolioTableName get the current portfolio table name
func GetPortfolioTableName() string {
	return portfolioTableName
}

// GetPortfolioNameTableName get the current portfolio name table name
func GetPortfolioNameTableName() string {
	return portfolioNameTableName
}

// GetAircraftTableName get the current aircraft table name
func GetAircraftTableName() string {
	return aircraftTableName
}

// GetAircraftModelTableName get the current aircraft model table name
func GetAircraftModelTableName() string {
	return aircraftModelTableName
}

// GetPortfolioAircraftTableName get the current portfolio aircraft table name
func GetPortfolioAircraftTableName() string {
	return portfolioAircraftTableName
}

// GetFlightTableName get the current flight table name
func GetFlightTableName() string {
	return flightTableName
}

// GetAirportTableName get the current airport table name
func GetAirportTableName() string {
	return airportTableName
}

// CreatePortfolioTable create the portfolio table on DynamoDB
func CreatePortfolioTable(tableName *string) error {
	if tableName != nil {
		portfolioTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(portfolioTableHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(portfolioTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(portfolioTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreatePortfolioNameTable create the portfolio name table on DynamoDB
func CreatePortfolioNameTable(tableName *string) error {
	// Specific for DynamoDB
	// Since Name is a unique attribute, it's better to make a dedicated table to ensure that constrait
	// Playing with projection or Range key, in portfolios table, would not work since Name has to be unique
	// We also don't want to make any column constraint for each Create request since optional columns are not indexed
	// On top of that, we could easily remove that constrait in the future by removing the table
	if tableName != nil {
		portfolioNameTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(portfolioNameTableHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(portfolioNameTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(portfolioNameTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreateAircraftTable create the aircraft table on DynamoDB
func CreateAircraftTable(tableName *string) error {
	if tableName != nil {
		aircraftTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(aircraftTableHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(aircraftTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(aircraftTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreateAircraftModelTable create the aircraft model table on DynamoDB
func CreateAircraftModelTable(tableName *string) error {
	if tableName != nil {
		aircraftModelTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(aircraftModelTableHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(aircraftModelTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(aircraftModelTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreatePortfolioAircraftTable create the portfolio aircraft table on DynamoDB
func CreatePortfolioAircraftTable(tableName *string) error {
	if tableName != nil {
		portfolioAircraftTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(portfolioAircraftTableHashKeyName),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(portfolioAircraftTableRangeKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(portfolioAircraftTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String(portfolioAircraftTableRangeKeyName),
				KeyType:       aws.String("RANGE"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			&dynamodb.GlobalSecondaryIndex{
				IndexName: aws.String(portfolioAircraftTableAircraftIndexName),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(portfolioAircraftTableAircraftIndexHashKeyName),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(readUnitsDefault),
					WriteCapacityUnits: aws.Int64(writeUnitsDefault),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(portfolioAircraftTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreateFlightTable create the flight table on DynamoDB
func CreateFlightTable(tableName *string) error {
	if tableName != nil {
		flightTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(flightTableHashKeyName),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String(flightTableAircraftIndexHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(flightTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			&dynamodb.GlobalSecondaryIndex{
				IndexName: aws.String(flightTableAircraftIndexName),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String(flightTableAircraftIndexHashKeyName),
						KeyType:       aws.String("HASH"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String(dynamodb.ProjectionTypeAll),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(readUnitsDefault),
					WriteCapacityUnits: aws.Int64(writeUnitsDefault),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(flightTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}

// CreateAirportTable create the airport table on DynamoDB
func CreateAirportTable(tableName *string) error {
	if tableName != nil {
		airportTableName = *tableName
	}
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(airportTableHashKeyName),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(airportTableHashKeyName),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readUnitsDefault),
			WriteCapacityUnits: aws.Int64(writeUnitsDefault),
		},
		TableName: aws.String(airportTableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		return err
	}
	return nil
}
