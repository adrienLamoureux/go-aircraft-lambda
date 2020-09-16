package dynamodatabase

import (
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/common"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// PortfolioDynamoDB is the Portfolio Database implementation for DynamoDB
type PortfolioDynamoDB struct {
}

// CreatePortfolio create a portfolio with DynamoDB
func (portfolioDynamoDB *PortfolioDynamoDB) CreatePortfolio(portfolioInfo *databases.CreatePortfolioInfo) error {
	timeNow := time.Now().Unix()
	err := createItem(&portfolioNameTableData{
		ID:       portfolioInfo.ID,
		Name:     portfolioInfo.Name,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, &keyItemInfo{
		HashKeyName: portfolioNameTableHashKeyName,
	}, portfolioNameTableName)
	if err != nil {
		return err
	}
	return createItem(&portfolioTableData{
		ID:       portfolioInfo.ID,
		Name:     portfolioInfo.Name,
		CreateTm: timeNow,
		UpdateTm: timeNow,
	}, &keyItemInfo{
		HashKeyName: portfolioTableHashKeyName,
	}, portfolioTableName)
}

// GetPortfolioInfo get a portfolio with DynamoDB
func (portfolioDynamoDB *PortfolioDynamoDB) GetPortfolioInfo(portfolioID string) (*databases.PortfolioInfo, error) {
	result, err := getItem(&keyItemInfo{
		HashKeyName:  portfolioTableHashKeyName,
		HashKeyValue: portfolioID,
	}, portfolioTableName)
	if err != nil {
		return nil, err
	}
	if result.Item == nil {
		return nil, nil
	}

	item := portfolioTableData{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return item.toAbstract(), nil
}

// GetPortfolioInfoList get a portfolio list with DynamoDB
func (portfolioDynamoDB *PortfolioDynamoDB) GetPortfolioInfoList() ([]*databases.PortfolioInfo, error) {
	result, err := scan(portfolioTableName)
	if err != nil {
		return nil, err
	}
	if len(result.Items) == 0 {
		return []*databases.PortfolioInfo{}, nil
	}
	portfolioInfoList := make([]*databases.PortfolioInfo, len(result.Items))
	for i, item := range result.Items {
		portfolioData := portfolioTableData{}
		err = dynamodbattribute.UnmarshalMap(item, &portfolioData)
		if err != nil {
			return []*databases.PortfolioInfo{}, err
		}
		portfolioInfoList[i] = portfolioData.toAbstract()
	}
	return portfolioInfoList, nil
}

// DeletePortfolio delete a portfolio and the related portfolio name with DynamoDB
func (portfolioDynamoDB *PortfolioDynamoDB) DeletePortfolio(portfolioID string) error {
	// Specific for DynamoDB
	portfolioInfo, err := portfolioDynamoDB.GetPortfolioInfo(portfolioID)
	if err != nil {
		return err
	}
	if portfolioInfo == nil {
		return common.NewErrorCode(common.ErrorItemNotFoundPortfolioCode)
	}
	err = deleteItem(&keyItemInfo{
		HashKeyName:  portfolioTableHashKeyName,
		HashKeyValue: portfolioInfo.ID,
	}, portfolioTableName)
	if err != nil {
		return err
	}
	return deleteItem(&keyItemInfo{
		HashKeyName:  portfolioNameTableHashKeyName,
		HashKeyValue: portfolioInfo.Name,
	}, portfolioNameTableName)
}
