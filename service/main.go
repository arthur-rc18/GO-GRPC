package main

import (
	"go-grpc/commons"
	block "go-grpc/commons/pb"
	"go-grpc/service/database"
	"log"
	"net"

	"cloud.google.com/go/bigquery"
	"google.golang.org/grpc"
)

// Server defines an implementation of BlocksServer protobuf interface (on commons)
//
//its methods are located on model.go file in this package
type Server struct {
	block.UnimplementedBlocksServer
	//	adicionar bqClient
	BQClient *bigquery.Client
}

func main() {
	server := grpc.NewServer()
	srv := &Server{}
	var err error
	srv.BQClient, err = database.NewBQClient()
	if err != nil {
		log.Fatalln("bqclient instance error:", err)
	}
	block.RegisterBlocksServer(server, srv)
	listener, err := net.Listen("tcp", *commons.PORTService)
	if err != nil {
		log.Fatalln("listener instance error:", err)
	}
	err = server.Serve(listener)
	if err != nil {
		log.Fatalln("server.serve error:", err)
	}
}
