package handlers

import "github.com/adrienLamoureux/go-aircraft-lambda/src/services"

// AircraftService the global aircraft service
var AircraftService services.IAircraftService

// PortfolioService the global portfolio service
var PortfolioService services.IPortfolioService

// FlightService the global flight service
var FlightService services.IFlightService

// MetricService the global metric service
var MetricService services.IMetricService

// AirportService the global airport service
var AirportService services.IAirportService
