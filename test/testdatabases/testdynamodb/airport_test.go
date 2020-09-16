package testdynamodb_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/setup"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases/dynamodatabase"
	"github.com/adrienLamoureux/go-aircraft-lambda/test/testdatabases"
	"github.com/adrienLamoureux/go-aircraft-lambda/test/testdatabases/testdynamodb"
	"github.com/stretchr/testify/suite"
)

// TestAirportDynamoDB is the test suite on DynamoDB
func TestAirportDynamoDB(t *testing.T) {
	err := setup.SetupDynamoDB()
	if err != nil {
		panic(err)
	}
	dynamoHelper := &testdynamodb.DynamoDBHelper{
		TableSuffix: fmt.Sprintf("test_%d", time.Now().UnixNano()/int64(time.Millisecond)),
		Client:      dynamodatabase.GetClient(),
	}
	dynamoHelper.SetAirportTableName(dynamodatabase.GetAirportTableName())
	suite.Run(t, &testdatabases.AirportTestSuite{
		Database: &dynamodatabase.AirportDynamoDB{},
		Helper:   dynamoHelper,
	})
}
