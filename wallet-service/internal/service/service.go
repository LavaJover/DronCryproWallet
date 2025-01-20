package service

import (
	"github.com/LavaJover/DronCryptoWallet/wallet/internal/config"
	"github.com/LavaJover/DronCryptoWallet/wallet/internal/wallet"
)

type WalletService struct{

}

func (walletService *WalletService) GetWalletBalance (address string) (float64, error){

	apiKey := config.MustLoad("/home/bodya/Рабочий стол/dronwallet/wallet-service/config/config.yaml").APIKey

	balance, err := wallet.GetUSDTBalance(address, apiKey)

	return balance, err
}