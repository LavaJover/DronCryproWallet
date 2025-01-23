package service

import "github.com/LavaJover/DronCryptoWallet/wallet-service/internal/wallet"

type WalletService struct{
	
}

func (walletService *WalletService) GetPrivateKey() (string, error){
	return wallet.GeneratePrivateKey()
}