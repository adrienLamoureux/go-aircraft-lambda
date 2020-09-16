package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// PortfolioAircraftDynamoDB is the Portfolio Aircraft Database implementation for DynamoDB
type PortfolioAircraftDynamoDB struct {
}

// CreatePortfolioAircraft create a portfolio aircraft with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) CreatePortfolioAircraft(portfolioAircraftInfo *databases.CreatePortfolioAircraftInfo) error {
	timeNow := time.Now().Unix()
	rangeKeyName := portfolioAircraftTableRangeKeyName
	return createItem(&portfolioAircraftTableData{
		PortfolioID: portfolioAircraftInfo.PortfolioID,
		AircraftID:  portfolioAircraftInfo.AircraftID,
		CreateTm:    timeNow,
		UpdateTm:    timeNow,
	}, &keyItemInfo{
		HashKeyName:  portfolioAircraftTableHashKeyName,
		RangeKeyName: &rangeKeyName,
	}, portfolioAircraftTableName)
}

// GetPortfolioAircraftInfo get a portfolio aircraft with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) GetPortfolioAircraftInfo(portfolioAircraftPortfolioID, portfolioAircraftAircraftID string) (*databases.PortfolioAircraftInfo, error) {
	var rangeKeyName = portfolioAircraftTableRangeKeyName
	result, err := getItem(&keyItemInfo{
		HashKeyName:   portfolioAircraftTableHashKeyName,
		HashKeyValue:  portfolioAircraftPortfolioID,
		RangeKeyName:  &rangeKeyName,
		RangeKeyValue: &portfolioAircraftAircraftID,
	}, portfolioAircraftTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := portfolioAircraftTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

// GetPortfolioAircraftInfos get a portfolio aircraft list of a portfolio with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) GetPortfolioAircraftInfos(portfolioAircraftPortfolioID string) ([]*databases.PortfolioAircraftInfo, error) {
	result, err := getItems(&keyItemInfo{
		HashKeyName:  portfolioAircraftTableHashKeyName,
		HashKeyValue: portfolioAircraftPortfolioID,
	}, portfolioAircraftTableName)
	if err != nil {
		return []*databases.PortfolioAircraftInfo{}, err
	}
	if len(result.Items) == 0 {
		return []*databases.PortfolioAircraftInfo{}, nil
	}

	portfolioAircraftInfoList := make([]*databases.PortfolioAircraftInfo, len(result.Items))
	for i, item := range result.Items {
		portfolioAircraftData := portfolioAircraftTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &portfolioAircraftData)
		if err != nil {
			return []*databases.PortfolioAircraftInfo{}, err
		}
		portfolioAircraftInfoList[i] = portfolioAircraftData.toAbstract()
	}
	return portfolioAircraftInfoList, nil
}

// GetPortfolioAircraftInfosByAircraftID get a portfolio aircraft list of an aircraft with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) GetPortfolioAircraftInfosByAircraftID(portfolioAircraftAircraftID string) ([]*databases.PortfolioAircraftInfo, error) {
	var indexName = portfolioAircraftTableAircraftIndexName
	result, err := getItems(&keyItemInfo{
		HashKeyName:  portfolioAircraftTableAircraftIndexHashKeyName,
		HashKeyValue: portfolioAircraftAircraftID,
		IndexName:    &indexName,
	}, portfolioAircraftTableName)
	if err != nil {
		return []*databases.PortfolioAircraftInfo{}, err
	}
	if len(result.Items) == 0 {
		return []*databases.PortfolioAircraftInfo{}, nil
	}

	portfolioAircraftInfoList := make([]*databases.PortfolioAircraftInfo, len(result.Items))
	for i, item := range result.Items {
		portfolioAircraftData := portfolioAircraftTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &portfolioAircraftData)
		if err != nil {
			return []*databases.PortfolioAircraftInfo{}, err
		}
		portfolioAircraftInfoList[i] = portfolioAircraftData.toAbstract()
	}
	return portfolioAircraftInfoList, nil
}

// DeletePortfolioAircraft delete a portfolio aircraft with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) DeletePortfolioAircraft(portfolioAircraftPortfolioID, portfolioAircraftAircraftID string) error {
	rangeKeyName := portfolioAircraftTableRangeKeyName
	return deleteItem(&keyItemInfo{
		HashKeyName:   portfolioAircraftTableHashKeyName,
		HashKeyValue:  portfolioAircraftPortfolioID,
		RangeKeyName:  &rangeKeyName,
		RangeKeyValue: &portfolioAircraftAircraftID,
	}, portfolioAircraftTableName)
}

// DeletePortfolioAircraftsByPortfolioID delete a portfolio aircraft list of a portfolio with DynamoDB
func (portfolioAircraftDynamoDB *PortfolioAircraftDynamoDB) DeletePortfolioAircraftsByPortfolioID(portfolioAircraftPortfolioID string) error {
	// With DynamoDB, we need both hash - range keys for deletion. Can not delete based on an index too (only for querying)
	result, err := getItems(&keyItemInfo{
		HashKeyName:  portfolioAircraftTableHashKeyName,
		HashKeyValue: portfolioAircraftPortfolioID,
	}, portfolioAircraftTableName)
	if err != nil {
		return err
	}
	if len(result.Items) == 0 {
		return common.NewErrorCode(common.ErrorItemNotFoundPortfolioCode)
	}
	// TODO: Optimize with batch delete
	rangeKeyName := portfolioAircraftTableRangeKeyName
	for _, item := range result.Items {
		portfolioAircraftData := portfolioAircraftTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &portfolioAircraftData)
		if err != nil {
			return err
		}
		aircraftID := portfolioAircraftData.AircraftID
		err = deleteItem(&keyItemInfo{
			HashKeyName:   portfolioAircraftTableHashKeyName,
			HashKeyValue:  portfolioAircraftData.PortfolioID,
			RangeKeyName:  &rangeKeyName,
			RangeKeyValue: &aircraftID,
		}, portfolioAircraftTableName)
		if err != nil {
			return err
		}
	}
	return nil
}
