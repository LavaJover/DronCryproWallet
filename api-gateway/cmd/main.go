package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	authpb "github.com/LavaJover/DronCryptoWallet/auth-service/proto/gen"

)

func main(){
	authServiceConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	defer authServiceConn.Close()

	authpb.NewAuthClient(authServiceConn)

}