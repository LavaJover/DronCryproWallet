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

func (pkRepo *PrivateKeyRepo) GetUserPrivateKeys (userID uint) (*[] string, error){
	var keys []string

	err := pkRepo.DB.Model(&models.PrivateKey{}).Where("user_id = ?", userID).Pluck("key", &keys).Error

	if err != nil {
		return nil, err
	}

	return &keys, nil

}