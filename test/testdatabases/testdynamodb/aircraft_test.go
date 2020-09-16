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

// TestAircraftDynamoDB is the test suite on DynamoDB
func TestAircraftDynamoDB(t *testing.T) {
	err := setup.SetupDynamoDB()
	if err != nil {
		panic(err)
	}
	dynamoHelper := &testdynamodb.DynamoDBHelper{
		TableSuffix: fmt.Sprintf("test_%d", time.Now().UnixNano()/int64(time.Millisecond)),
		Client:      dynamodatabase.GetClient(),
	}
	dynamoHelper.SetAircraftTableName(dynamodatabase.GetAircraftTableName())
	suite.Run(t, &testdatabases.AircraftTestSuite{
		Database: &dynamodatabase.AircraftDynamoDB{},
		Helper:   dynamoHelper,
	})
}
