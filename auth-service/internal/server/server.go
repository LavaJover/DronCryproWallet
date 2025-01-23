package server

import (
	"context"

	service "github.com/LavaJover/DronCryptoWallet/auth/internal/service"
	authpb "github.com/LavaJover/DronCryptoWallet/auth/proto/gen"
)

type AuthServer struct{
	*service.AuthService
	authpb.UnimplementedAuthServer
}

func (authServer *AuthServer) Register (ctx context.Context, req *authpb.RegisterRequest) (*authpb.RegisterResponse, error){
	err := authServer.AuthService.Register(req.Email, req.Password)

	if err != nil{
		return nil, err
	}

	return &authpb.RegisterResponse{}, nil
}

func (authServer *AuthServer) Login (ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error){
	token, err := authServer.AuthService.Login(req.Email, req.Password)

	if err != nil{
		return nil, err
	}

	return &authpb.LoginResponse{
		Token: token,
	}, nil
}

func (authServer *AuthServer) ValidateJWT (ctx context.Context, req *authpb.ValidateJWTRequest) (*authpb.ValidateJWTResponse, error){
	err := authServer.AuthService.ValidateJWT(req.Token)

	if err != nil{
		return &authpb.ValidateJWTResponse{Valid: false}, err
	}

	return &authpb.ValidateJWTResponse{Valid: true}, nil
}