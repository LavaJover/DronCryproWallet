package repo

import (
	"github.com/LavaJover/DronCryptoWallet/auth/internal/middleware/password"
	"github.com/LavaJover/DronCryptoWallet/auth/internal/models"
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


func (repo *UserRepo) FindUserByEmail (email string) *models.User{
	user := &models.User{}

	repo.Where("email = ?", email).First(user)

	return user
}