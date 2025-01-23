package server

import (
	"context"

	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/service"
	walletpb "github.com/LavaJover/DronCryptoWallet/wallet-service/proto/gen"
)

type WalletServer struct{
	walletpb.UnimplementedWalletServiceServer
	WalletService *service.WalletService
}

func (walletServer *WalletServer) GetPrivateKey (ctx context.Context, req *walletpb.GetPrivateKeyRequest) (*walletpb.GetPrivateKeyResponse, error){
	privateKey, err :=  walletServer.WalletService.GetPrivateKey()

	if err != nil{
		return nil, err
	}

	return &walletpb.GetPrivateKeyResponse{
		PrivateKey: privateKey,
	}, nil
}

func (walletServer *WalletServer) RegisterPrivateKey (ctx context.Context, req *walletpb.RegisterPrivateKeyRequest) (*walletpb.RegisterPrivateKeyResponse, error){

	pk := req.PrivateKey
	jwtToken := req.Token

	err := walletServer.WalletService.RegisterPrivateKey(jwtToken, pk)

	if err != nil{
		return nil, err
	}

	return &walletpb.RegisterPrivateKeyResponse{
		Status: "OK",
	}, nil
}

func (walletServer *WalletServer) GetUserPrivateKeys (ctx context.Context, req *walletpb.GetUserPrivateKeysRequest) (*walletpb.GetUserPrivateKeysResponse, error){
	jwtToken := req.Token

	keys, err := walletServer.WalletService.GetUserPrivateKeys(jwtToken)

	if err != nil{
		return nil, err
	}

	return &walletpb.GetUserPrivateKeysResponse{
		PrivateKeys: *keys,
	}, err

}