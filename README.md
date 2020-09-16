# go-aircraft-lambda

## Getting Started

### Software to install

- Install Go 1.13+ (https://golang.org/dl/)
- Install AWS CLI (https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
- Configure AWS (https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html#cli-configure-quickstart-config) with any credentials
- Download DynamoDBLocal (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html)

### Setup DynamoDB

#### Run DynamoDB

Go to the DynamoDB folder to start it with (Unix): https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html

```
java -Djava.library.path=./DynamoDBLocal_lib -jar DynamoDBLocal.jar -sharedDb
```

#### Run dynamodb-admin (Optional)

A nice tool that helps to vizualize the dynamodb tables

```
npm install -g dynamodb-admin
```

then

```
dynamodb-admin
```

#### Override default dynamoDB location with environment variables (Optional)

Default expected dynamoDB endpoint: http://localhost:8000 on region us-west-2 but you can change it with

```
export DYNAMO_REGION=my_region
```

```
export DYNAMO_ENDPOINT=my_endpoint
```

## Dependencies installation

```
go mod tidy
```

## Run Tests

```
go test ./...
```

## Run Local Server

The local server will be accessible at http://localhost:7200

### Without VSCode

```
go run src/dev_main.go
```

### With VSCode

Start debugging (F5 shortcut by default)

## Play with the API

### Import the API

Import the API specification, written in api-v1.json, into PostMan (for example)

### Setup the tables

Call the specific endpoint (not in the api-v1.json) POST http://localhost:7200/dev/createDynamoTables in order to create the tables with the aircraft models, aircrafts and allowed airports pre-populated in DynamoDB

### Run a scenario

Since airports, aircraft models and aircrafts are already here, you can then run a classic scenario :
- create a porfolio
- associate some aircraft to this portfolio
- create a flight for an aircraft to some airport
- get the flight metric for the portfolio
- get the flight metric for all the aircraft model
- unlink an aircraft from a portfolio
- delete a portfolio