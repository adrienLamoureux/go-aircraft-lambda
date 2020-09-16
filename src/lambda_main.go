// +build lambda

package main

import (
	"errors"

	"github.com/adrienLamoureux/go-aircraft-lambda/src/handlers"
	"github.com/adrienLamoureux/go-aircraft-lambda/src/setup"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davyzhang/agw"
	"github.com/gorilla/mux"
)

func main() {
	err := setup.SetupDynamoDB()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter()
	handlers.CreateRouter(router)
	if handler == nil {
		return errors.New("The request handler cannot be nil")
	}
	lambda.Start(agw.Handler(router))
	if err != nil {
		panic(err)
	}
}
