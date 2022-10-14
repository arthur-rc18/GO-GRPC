package tests

import (
	"go-grpc/commons"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SetupGrpcTestClient setup a grpc client connection for unity test purpose
func SetupGrpcTestClient() block.BlocksClient {
	conn, err := grpc.Dial(*commons.PORTClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("error dial create test connection", err.Error())
	}
	client := block.NewBlocksClient(conn)
	return client
}
