package main

import (
	"log"
	"log/slog"
	"net"

	"github.com/LavaJover/dronwallet/auth/internal/db"
	repo "github.com/LavaJover/dronwallet/auth/internal/repositories"
	"github.com/LavaJover/dronwallet/auth/internal/server"
	service "github.com/LavaJover/dronwallet/auth/internal/serviice"
	authpb "github.com/LavaJover/dronwallet/auth/proto/gen"
	"google.golang.org/grpc"
)

func main(){
	grpcServer := grpc.NewServer()

	dsn := "host=localhost user=postgres password=admin dbname=dronwallet port=5432 sslmode=disable"
	db, err := db.InitDB(dsn)

	if err != nil{
		log.Fatalf("Failed to connect to database: %v", err)
	}

	slog.Info("Connected to database!")

	userRepo := repo.UserRepo{DB: db}
	
	authServer := server.AuthServer{AuthService: &service.AuthService{UserRepo: &userRepo}}

	authpb.RegisterAuthServer(grpcServer, &authServer)

	listener, err := net.Listen("tcp", ":50052")

	if err != nil{
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	log.Println("gRPC server is running on port 50052")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}