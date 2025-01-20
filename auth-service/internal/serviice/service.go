package service

import (
	// token "github.com/LavaJover/dronwallet/auth/internal/middleware/JWT"
	"errors"

	token "github.com/LavaJover/dronwallet/auth/internal/middleware/JWT"
	"github.com/LavaJover/dronwallet/auth/internal/middleware/password"
	// "github.com/LavaJover/dronwallet/auth/internal/middleware/password"
	"github.com/LavaJover/dronwallet/auth/internal/models"
	repo "github.com/LavaJover/dronwallet/auth/internal/repositories"
)


type AuthService struct{
	*repo.UserRepo
}


func (authService *AuthService) Register (email string, raw_password string) (uint, error){

	newUser := models.User{
		Email: email,
		Password: raw_password,
	}

	authService.UserRepo.AddUser(&newUser)

	return newUser.ID, nil
}

func (authService *AuthService) Login (email string, rawRassword string) (string, error){

	user := authService.UserRepo.FindUserByEmail(email)

	if user.Email == ""{
		return "", errors.New("user " + email + " not found!")
	}

	if !password.CheckPassword(rawRassword, user.Password){
		return "", errors.New("wrong password for user " + email)
	}

	return token.GenerateJWT(int(user.ID))
}