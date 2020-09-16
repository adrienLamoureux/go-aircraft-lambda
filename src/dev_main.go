// +build !lambda

package main

import (
	"fmt"
	"net/http"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/databases/dynamodatabase"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/services"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/setup"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/handlers"
	"github.com/gorilla/mux"
)

func createDevRouter() *mux.Router {
	router := mux.NewRouter()
	handlers.CreateRouter(router)
	devRouter := router.PathPrefix("/dev").Subrouter()
	devRouter.HandleFunc("/createDynamoTables", createDynamoTables).Methods("POST")
	return router
}

func main() {
	err := setup.SetupDynamoDB()
	if err != nil {
		panic(err)
	}
	port := "7200"
	fmt.Printf("Starting server at %s\n", port)
	err = http.ListenAndServe(":"+port, createDevRouter())
	if err != nil {
		fmt.Printf("Failed to start server at port %s\nError: %s", port, err.Error())
		panic(err)
	}
}

func createDynamoTables(w http.ResponseWriter, r *http.Request) {
	err := dynamodatabase.CreateAircraftModelTable(nil)
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraftModel(&services.AircraftModelVO{
		ID: "A320-200",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraftModel(&services.AircraftModelVO{
		ID: "B737-800",
	})
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreateAircraftTable(nil)
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "ZS-GAO",
		Model: "A320-200",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "D-AIUO",
		Model: "A320-200",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "B-6636",
		Model: "A320-200",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "LY-BFM",
		Model: "B737-800",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "G-GDFW",
		Model: "B737-800",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AircraftService.CreateAircraft(&services.AircraftVO{
		ID:    "XA-AMZ",
		Model: "B737-800",
	})
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreatePortfolioTable(nil)
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreatePortfolioNameTable(nil)
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreatePortfolioAircraftTable(nil)
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreateFlightTable(nil)
	if err != nil {
		panic(err)
	}
	err = dynamodatabase.CreateAirportTable(nil)
	if err != nil {
		panic(err)
	}
	_, err = handlers.AirportService.CreateAirport(&services.AirportVO{
		ID: "DUB",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AirportService.CreateAirport(&services.AirportVO{
		ID: "STN",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AirportService.CreateAirport(&services.AirportVO{
		ID: "MXP",
	})
	if err != nil {
		panic(err)
	}
	_, err = handlers.AirportService.CreateAirport(&services.AirportVO{
		ID: "CDG",
	})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
}
