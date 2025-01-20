package server

import (
	"context"

	"github.com/LavaJover/DronCryptoWallet/wallet/internal/service"
	walletpb "github.com/LavaJover/DronCryptoWallet/wallet/proto/gen"
)

type WalletServer struct{
	walletpb.UnimplementedWalletServiceServer
	WalletService *service.WalletService
}

func (walletServer *WalletServer) GetWalletBalance (ctx context.Context, req *walletpb.GetBalanceRequest) (*walletpb.GetBalanceResponse, error){
	balance, err := walletServer.WalletService.GetWalletBalance(req.Address)

	if err != nil{
		return nil, err
	}

	return &walletpb.GetBalanceResponse{Balance: balance}, nil
}