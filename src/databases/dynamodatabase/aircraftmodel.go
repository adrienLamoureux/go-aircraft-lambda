package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// AircraftModelDynamoDB is the Aircraft Model Database implementation for DynamoDB
type AircraftModelDynamoDB struct {
}

// CreateAircraftModel create an aircraft model with DynamoDB
func (aircraftDynamoDB *AircraftModelDynamoDB) CreateAircraftModel(aircraftModelInfo *databases.CreateAircraftModelInfo) error {
	timeNow := time.Now().Unix()
	return createItem(&aircraftModelTableData{
		ID:       aircraftModelInfo.ID,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, &keyItemInfo{
		HashKeyName: aircraftModelTableHashKeyName,
	}, aircraftModelTableName)
}

// GetAircraftModelInfo get an aircraft model from DynamoDB
func (aircraftDynamoDB *AircraftModelDynamoDB) GetAircraftModelInfo(aircraftModelID string) (*databases.AircraftModelInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  aircraftModelTableHashKeyName,
		HashKeyValue: aircraftModelID,
	}, aircraftModelTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := aircraftModelTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

// GetAircraftModelInfoList get an aircraft model list from DynamoDB
func (aircraftDynamoDB *AircraftModelDynamoDB) GetAircraftModelInfoList() ([]*databases.AircraftModelInfo, error) {
	result, err := scan(aircraftModelTableName)
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return []*databases.AircraftModelInfo{}, nil
	}
	aircraftModelInfoList := make([]*databases.AircraftModelInfo, len(result.Items))
	for i, item := range result.Items {
		aircraftModelData := aircraftModelTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &aircraftModelData)
		if err != nil {
			return []*databases.AircraftModelInfo{}, err
		}
		aircraftModelInfoList[i] = aircraftModelData.toAbstract()
	}
	return aircraftModelInfoList, nil
}
