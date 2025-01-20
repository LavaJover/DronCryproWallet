package repo

import (
	"github.com/LavaJover/dronwallet/auth/internal/models"
	"log/slog"
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

func (repo *UserRepo) FindUserByCredentials (email string, raw_password string){
	user := &models.User{}
	repo.DB.Where("email = ? AND password = ?", email, raw_password).First(&user)

	slog.Info("User found: ", "User: ", user)
}