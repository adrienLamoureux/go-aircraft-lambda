package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// AirportDynamoDB is the Airport Database implementation for DynamoDB
type AirportDynamoDB struct {
}

// CreateAirport create an airport with DynamoDB
func (airportDynamoDB *AirportDynamoDB) CreateAirport(airportInfo *databases.CreateAirportInfo) error {
	timeNow := time.Now().Unix()
	return createItem(&airportTableData{
		ID:       airportInfo.ID,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, &keyItemInfo{
		HashKeyName: airportTableHashKeyName,
	}, airportTableName)
}

// GetAirportInfo get an airport from DynamoDB
func (airportDynamoDB *AirportDynamoDB) GetAirportInfo(airportID string) (*databases.AirportInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  airportTableHashKeyName,
		HashKeyValue: airportID,
	}, airportTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := airportTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}
