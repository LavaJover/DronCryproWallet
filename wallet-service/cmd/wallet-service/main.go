package main

import (
	"log"
	"net"

	"github.com/LavaJover/dronwallet/wallet/internal/server"
	"github.com/LavaJover/dronwallet/wallet/internal/service"
	walletpb "github.com/LavaJover/dronwallet/wallet/proto/gen"
	"google.golang.org/grpc"
)

func main(){
	walletService := &service.WalletService{}

	grpcServer := grpc.NewServer()
	walletServer := &server.WalletServer{WalletService: walletService}
	walletpb.RegisterWalletServiceServer(grpcServer, walletServer)

	listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen on port 50051: %v", err)
    }

    log.Println("gRPC server is running on port 50051")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }

}