package main

import (
	"log"
	"log/slog"
	"net"

	"github.com/LavaJover/DronCryptoWallet/auth/internal/config"
	"github.com/LavaJover/DronCryptoWallet/auth/internal/db"
	repo "github.com/LavaJover/DronCryptoWallet/auth/internal/repositories"
	"github.com/LavaJover/DronCryptoWallet/auth/internal/server"
	service "github.com/LavaJover/DronCryptoWallet/auth/internal/serviice"
	authpb "github.com/LavaJover/DronCryptoWallet/auth/proto/gen"
	"google.golang.org/grpc"
)

func main(){

	cfg := config.MustLoad()

	grpcServer := grpc.NewServer()

	dsn := cfg.Dsn
	db, err := db.InitDB(dsn)

	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}

	slog.Info("Connected to database!")

	userRepo := repo.UserRepo{DB: db}
	
	authServer := server.AuthServer{AuthService: &service.AuthService{UserRepo: &userRepo}}

	authpb.RegisterAuthServer(grpcServer, &authServer)

	listener, err := net.Listen("tcp", ":"+cfg.Port)

	if err != nil{
		log.Fatalf("Failed to listen on port %s: %v", cfg.Port, err)
	}

	slog.Info("server running on port " + cfg.Port)
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}