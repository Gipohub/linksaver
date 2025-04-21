package main

import (
	"log"
	"net"

	linksaver "github.com/Gipohub/linksaver/proto"
	"github.com/Gipohub/linksaver/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	linksaver.RegisterLinkSaverServer(s, server.New())

	log.Println("LinkSaver gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
