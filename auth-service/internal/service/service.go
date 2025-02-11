package service

import (
	"errors"

	"github.com/LavaJover/DronCryptoWallet/auth/internal/middleware"
	"github.com/LavaJover/DronCryptoWallet/auth/internal/middleware/password"
	"github.com/LavaJover/DronCryptoWallet/auth/internal/models"
	repo "github.com/LavaJover/DronCryptoWallet/auth/internal/repositories"
)

type AuthService struct{
	*repo.UserRepo
}

func (authService *AuthService) Register (email string, raw_password string) error{

	newUser := models.User{
		Email: email,
		Password: raw_password,
	}


	authService.UserRepo.AddUser(&newUser)

	return nil
}

func (authService *AuthService) Login (email string, rawRassword string) (string, error){

	user := authService.UserRepo.FindUserByEmail(email)

	if user.Email == ""{
		return "", errors.New("user " + email + " not found!")
	}

	if !password.CheckPassword(rawRassword, user.Password){
		return "", errors.New("wrong password for user " + email)
	}

	return middleware.GenerateJWT(int(user.ID))
}

func (authService *AuthService) ValidateJWT (token string) error{
	_, err := middleware.ValidateJWT(token)

	return err
}