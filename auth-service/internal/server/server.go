package server

import (
	"context"
	"log"

	service "github.com/LavaJover/DronCryptoWallet/auth/internal/serviice"
	authpb "github.com/LavaJover/DronCryptoWallet/auth/proto/gen"
)

type AuthServer struct{
	*service.AuthService
	authpb.UnimplementedAuthServer
}

func (authServer *AuthServer) Register (ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error){
	id, err := authServer.AuthService.Register(req.Email, req.Password)

	if err != nil{
		return nil, err
	}

	return &authpb.RegisterResponse{
		UserId: int64(id),
	}, nil
}

func (authServer *AuthServer) Login (ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error){
	token, err := authServer.AuthService.Login(req.Email, req.Password)

	if err != nil{
		log.Fatalf("Failed to log user in: %v", err)
	}

	return &authpb.LoginResponse{
		Token: token,
	}, nil
}