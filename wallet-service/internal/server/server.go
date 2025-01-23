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