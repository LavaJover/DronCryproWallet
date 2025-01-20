package db

import (
	"github.com/LavaJover/DronCryptoWallet/auth/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) (*gorm.DB, error){
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&models.User{})

	if err != nil{
		return nil, err
	}

	return db, nil
}