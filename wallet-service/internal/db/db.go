package db

import (
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.PrivateKey{})

	if err != nil{
		return nil, err
	}

	return db, nil
}