package service

import (
	jwttoken "github.com/LavaJover/DronCryptoWallet/wallet-service/internal/middleware/JWT"
	repo "github.com/LavaJover/DronCryptoWallet/wallet-service/internal/repositories"
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/wallet"
)

type WalletService struct{
	*repo.PrivateKeyRepo
}

func (walletService *WalletService) GetPrivateKey() (string, error){
	return wallet.GeneratePrivateKey()
}

func (walletService *WalletService) RegisterPrivateKey(token string, key string) error{

	userID, err := jwttoken.DecodeJWT(token)

	if err != nil{
		return err
	}

	return walletService.PrivateKeyRepo.RegisterPK(key, userID)
}

func (walletService *WalletService) GetUserPrivateKeys(token string) (*[]string, error){
	userID, err := jwttoken.DecodeJWT(token)

	if err != nil{
		return nil, err
	}

	keys, err := walletService.PrivateKeyRepo.GetUserPrivateKeys(userID)

	if err != nil{
		return nil, err
	}

	return keys, nil

}