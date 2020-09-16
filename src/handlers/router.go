package handlers

import (
	"github.com/gorilla/mux"
)

func CreateRouter(router *mux.Router) {
	router.HandleFunc("/ping", ping).Methods("GET")

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/portfolios", handleGetPortfolioList).Methods("GET")
	apiRouter.HandleFunc("/portfolio", handleCreatePortfolio).Methods("POST")
	apiRouter.HandleFunc("/aircraftModel", handleCreateAircraftModel).Methods("POST")
	apiRouter.HandleFunc("/aircraft", handleCreateAircraft).Methods("POST")
	apiRouter.HandleFunc("/airport", handleCreateAirport).Methods("POST")
	apiRouter.HandleFunc("/flight", handleCreateFlight).Methods("POST")
	apiRouter.HandleFunc("/flightMetrics", handleGetAircraftModelFlightMetricList).Methods("GET")

	portfolioRouter := apiRouter.PathPrefix("/portfolio").Subrouter()
	portfolioRouter.HandleFunc("/{portfolioId}", handleGetPortfolio).Methods("GET")
	portfolioRouter.HandleFunc("/{portfolioId}", handleDeletePortfolio).Methods("DELETE")

	portfolioDetailsRouter := portfolioRouter.PathPrefix("/{portfolioId}").Subrouter()
	portfolioDetailsRouter.HandleFunc("/aircraft", handleCreatePortfolioAircraft).Methods("POST")
	portfolioDetailsRouter.HandleFunc("/flightMetrics", handleGetPortfolioFlightMetricList).Methods("GET")

	portfolioDetailsAircraftRouter := portfolioDetailsRouter.PathPrefix("/aircraft").Subrouter()
	portfolioDetailsAircraftRouter.HandleFunc("/{aircraftId}", handleDeletePortfolioAircraft).Methods("DELETE")
}
