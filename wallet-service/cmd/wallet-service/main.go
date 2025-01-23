package main

import (
	"log"
	"net"
	"log/slog"

	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/db"
	repo "github.com/LavaJover/DronCryptoWallet/wallet-service/internal/repositories"
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/server"
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/service"
	walletpb "github.com/LavaJover/DronCryptoWallet/wallet-service/proto/gen"
	"google.golang.org/grpc"
)

func main(){

	dsn := "host=localhost user=postgres password=admin dbname=dronwallet port=5432 sslmode=disable"

	db, err := db.InitDB(dsn)

	if err != nil{
		log.Fatalf("failed to connect to db: %v", err)
	}

	slog.Info("successfully connected to database")

	walletRepo := repo.PrivateKeyRepo{DB: db}

	walletService := &service.WalletService{&walletRepo}

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