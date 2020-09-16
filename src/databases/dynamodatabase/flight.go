package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// FlightDynamoDB is the Flight Database implementation for DynamoDB
type FlightDynamoDB struct {
}

// CreateFlight create a flight with DynamoDB
func (flightDynamoDB *FlightDynamoDB) CreateFlight(flightInfo *databases.CreateFlightInfo) error {
	timeNow := time.Now().Unix()
	return createItem(&flightTableData{
		FlightID:         flightInfo.FlightID,
		AircraftID:       flightInfo.AircraftID,
		DepartureAirport: flightInfo.DepartureAirport,
		DepartureTime:    flightInfo.DepartureTime,
		ArrivalAirport:   flightInfo.ArrivalAirport,
		ArrivalTime:      flightInfo.ArrivalTime,
		CreateTm:         timeNow,
		UpdateTm:         timeNow,
	}, &keyItemInfo{
		HashKeyName: flightTableHashKeyName,
	}, flightTableName)
}

// GetFlightInfo get a flight with DynamoDB
func (flightDynamoDB *FlightDynamoDB) GetFlightInfo(flightID string) (*databases.FlightInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  flightTableHashKeyName,
		HashKeyValue: flightID,
	}, flightTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := flightTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

// GetFlightInfosByAircraftID get a flight list of an Aircraft with DynamoDB
func (flightDynamoDB *FlightDynamoDB) GetFlightInfosByAircraftID(aircraftID string) ([]*databases.FlightInfo, error) {
	indexName := flightTableAircraftIndexName
	result, err := getItems(&keyItemInfo{
		HashKeyName:  flightTableAircraftIndexHashKeyName,
		HashKeyValue: aircraftID,
		IndexName:    &indexName,
	}, flightTableName)
	if err != nil {
		return []*databases.FlightInfo{}, err
	}
	if len(result.Items) == 0 {
		return []*databases.FlightInfo{}, nil
	}

	flightInfoList := make([]*databases.FlightInfo, len(result.Items))
	for i, item := range result.Items {
		flightData := flightTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &flightData)
		if err != nil {
			return []*databases.FlightInfo{}, err
		}
		flightInfoList[i] = flightData.toAbstract()
	}
	return flightInfoList, nil
}
