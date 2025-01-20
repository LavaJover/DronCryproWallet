package uservalid

import (
	"errors"
	"net/mail"

	"github.com/LavaJover/DronCryptoWallet/api-gateway/models"
)

func ValidateUserRequest(user *models.User) error{
	if user.Email == "" || user.Password == ""{
		return errors.New("empty fields")
	}

	if len(user.Password) < 8{
		return errors.New("too short password")
	}

	_, err := mail.ParseAddress(user.Email)

	if err != nil{
		return err
	}

	return nil
}