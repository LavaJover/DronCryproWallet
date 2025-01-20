package repo

import (
	"github.com/LavaJover/dronwallet/auth/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct{
	*gorm.DB
}

func (repo *UserRepo) AddUser (user *models.User){
	repo.Create(&user)
}

func (repo *UserRepo) DeleteUser (id uint){
	repo.DB.Delete(&models.User{}, id)
}