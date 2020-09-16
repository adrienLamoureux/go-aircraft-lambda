package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// AircraftDynamoDB is the Aircraft Database implementation for DynamoDB
type AircraftDynamoDB struct {
}

// CreateAircraft create an aircraft with DynamoDB
func (aircraftDynamoDB *AircraftDynamoDB) CreateAircraft(aircraftInfo *databases.CreateAircraftInfo) error {
	timeNow := time.Now().Unix()
	return createItem(&aircraftTableData{
		ID:       aircraftInfo.ID,
		Model:    aircraftInfo.Model,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, &keyItemInfo{
		HashKeyName: aircraftTableHashKeyName,
	}, aircraftTableName)
}

// GetAircraftInfo get an aircraft from DynamoDB
func (aircraftDynamoDB *AircraftDynamoDB) GetAircraftInfo(aircraftID string) (*databases.AircraftInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  aircraftTableHashKeyName,
		HashKeyValue: aircraftID,
	}, aircraftTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := aircraftTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

// GetAircraftInfoList get an aircraft from DynamoDB
func (aircraftDynamoDB *AircraftDynamoDB) GetAircraftInfoList() ([]*databases.AircraftInfo, error) {
	result, err := scan(aircraftTableName)
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return []*databases.AircraftInfo{}, nil
	}
	aircraftInfoList := make([]*databases.AircraftInfo, len(result.Items))
	for i, item := range result.Items {
		aircraftData := aircraftTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &aircraftData)
		if err != nil {
			return []*databases.AircraftInfo{}, err
		}
		aircraftInfoList[i] = aircraftData.toAbstract()
	}
	return aircraftInfoList, nil
}
