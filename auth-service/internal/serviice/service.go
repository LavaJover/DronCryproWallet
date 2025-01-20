package service

import (
	// token "github.com/LavaJover/dronwallet/auth/internal/middleware/JWT"
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

func (authService *AuthService) Login (email string, raw_password string) (string, error){

	// token.GenerateJWT()

	authService.UserRepo.FindUserByCredentials(email, raw_password)

	return "asdasd", nil

}