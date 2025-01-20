package usertest

import (
	"testing"

	"github.com/LavaJover/DronCryptoWallet/api-gateway/models"
	uservalid "github.com/LavaJover/DronCryptoWallet/api-gateway/validation/user"
)

func TestValidateUserRequest(t *testing.T){
	users := []*models.User{
		&models.User{
			Email: "test@gmail.com",
			Password: "asd87JJyaks",
		},
		&models.User{
			Email: "testgmail.com",
			Password: "IIyyabhsus",
		},
		&models.User{
			Email: "test@gmail.com",
			Password: "123",
		},
	}

	if uservalid.ValidateUserRequest(users[0]) != nil || uservalid.ValidateUserRequest(users[1]) == nil || uservalid.ValidateUserRequest(users[2]) == nil{
		t.Error("wrong validation")
	}
	
}