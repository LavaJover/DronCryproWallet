package repo

import (
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/models"
	"gorm.io/gorm"
)

type PrivateKeyRepo struct{
	*gorm.DB
}

func (pkRepo *PrivateKeyRepo) RegisterPK (key string, userID uint) error{
	pkRepo.DB.Create(&models.PrivateKey{
		Key: key,
		UserID: userID,
	})

	return nil
}