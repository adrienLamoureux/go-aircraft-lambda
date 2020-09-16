package setup

import (
	"os"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases/dynamodatabase"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/handlers"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
)

// SetupDynamoDB setup the architecture to work with DynamoDB
func SetupDynamoDB() error {
	region := os.Getenv("DYNAMO_REGION")
	if len(region) == 0 {
		region = "us-west-2"
	}
	endpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(endpoint) == 0 {
		endpoint = "http://localhost:8000"
	}
	err := dynamodatabase.InitializeClient(region, endpoint)
	if err != nil {
		return err
	}
	handlers.AircraftService = &services.AircraftDefaultService{
		AircraftDatabase:      &dynamodatabase.AircraftDynamoDB{},
		AircraftModelDatabase: &dynamodatabase.AircraftModelDynamoDB{},
	}
	handlers.PortfolioService = &services.PortfolioDefaultService{
		PortfolioDatabase:         &dynamodatabase.PortfolioDynamoDB{},
		PortfolioAircraftDatabase: &dynamodatabase.PortfolioAircraftDynamoDB{},
		AircraftDatabase:          &dynamodatabase.AircraftDynamoDB{},
	}
	handlers.FlightService = &services.FlightDefaultService{
		FlightDatabase:   &dynamodatabase.FlightDynamoDB{},
		AircraftDatabase: &dynamodatabase.AircraftDynamoDB{},
		AirportDatabase:  &dynamodatabase.AirportDynamoDB{},
	}
	handlers.MetricService = &services.MetricDefaultService{
		AircraftDatabase:          &dynamodatabase.AircraftDynamoDB{},
		FlightDatabase:            &dynamodatabase.FlightDynamoDB{},
		PortfolioDatabase:         &dynamodatabase.PortfolioDynamoDB{},
		PortfolioAircraftDatabase: &dynamodatabase.PortfolioAircraftDynamoDB{},
	}
	handlers.AirportService = &services.AirportDefaultService{
		AirportDatabase: &dynamodatabase.AirportDynamoDB{},
	}
	return nil
}
