package models

import "gorm.io/gorm"

type PrivateKey struct{
	gorm.Model
	Key string `gorm:"not null;unique"`
	UserID uint
}