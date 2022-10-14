# GRPC Project

## Cloning the project

```sh
git clone https://github.com/arthur-rc18/GO-GRPC.git
```

## Project structure

Project based on a gRPC instance connected with a client and service. The client represents the Rest Api while the service make a connection to a **bigQuery** database and sends a hard coded query to it. This query will receive a specific Id and then will search for a block based on that Id, bringing temperature and precipitations data of that block.

## Compiling

To execute the program, both client and service has to be running.

    To run the client: 
        cd client
        go run main.go

    To run the service:
        cd service
        go run main.go

## Execution flow

When an Id is passed through the URL, is the service who receives and searchs for it in the database, in case of an existent blockId (Z0164), the response will be:

``` JSON
{
    "data_timestamp": "2022-06-28T23:05:24.708115Z",
    "created_timestamp": "2022-06-28T23:05:24.708115Z",
    "temperature_inst": "21.012310",
    "temperature_min": "20.148813",
    "temperature_max": "25.221354",
    "precipitation_inst": "0.000000"
}
```

## Tests

All the code are covered up with tests. To check when it receives an existed and no existed block, and the specifics status code. To run the tests:
 
        cd client/handlers
        go test

        cd service
        go test

## Block definition

The block is a struct with the specific fields:

``` go
type Block struct {
	DataTimestamp     time.Time 
	CreatedTimestamp  time.Time
	TemperatureInst   string    
	TemperatureMin    string    
	TemperatureMax    string    
	PrecipitationInst string    
}
```