package service

import (
	"context"
	"go-grpc/commons"
	"go-grpc/commons/models"
	block "go-grpc/commons/pb"
	"log"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

// GetBlockById create the grpc client connection and calls the rpc function GetBlockById on proto
//
// the function creates a request message passing the given id and retrieve the grpc response block data returning it
func GetBlockById(id string) (models.Block, error) {

	conn, err := grpc.Dial(*commons.PORTClient, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("error dial create connection", err.Error())
		return models.Block{}, err
	}

	client := block.NewBlocksClient(conn)
	req := block.RequestID{
		Id: id,
	}

	res, err := client.GetBlockById(context.Background(), &req)
	if err != nil {
		log.Println("err call rpc function getblockbyid", err.Error())
		return models.Block{}, err
	}

	blockDTO := models.TransformResponseIntoBlockDTO(res)

	return blockDTO, nil
}
