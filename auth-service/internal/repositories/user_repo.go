package repo

import (
	"log/slog"

	"github.com/LavaJover/dronwallet/auth/internal/middleware/password"
	"github.com/LavaJover/dronwallet/auth/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct{
	*gorm.DB
}

func (repo *UserRepo) AddUser (user *models.User){
	repo.Create(&models.User{
		Email: user.Email,
		Password: password.HashPassword(user.Password),
	})
}

func (repo *UserRepo) DeleteUser (id uint){
	repo.DB.Delete(&models.User{}, id)
}

func (repo *UserRepo) FindUserByCredentials (email string, raw_password string){
	user := &models.User{}
	repo.DB.Where("email = ? AND password = ?", email, password.HashPassword(raw_password)).First(&user)

	slog.Info("User found: ", "User: ", user)
}