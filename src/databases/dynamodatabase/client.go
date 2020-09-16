package dynamodatabase

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var svc *dynamodb.DynamoDB

// InitializeClient is initializing the DynamoDB client
func InitializeClient(region, endpoint string) error {
	if len(region) == 0 {
		return errors.New("Missing region")
	}
	if len(endpoint) == 0 {
		return errors.New("Missing endpoint")
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   &region,
			Endpoint: &endpoint,
		},
	})
	if err != nil {
		return err
	}
	svc = dynamodb.New(sess)
	return nil
}

// GetClient return the current DynamoDB client
func GetClient() *dynamodb.DynamoDB {
	return svc
}
