package passwordTest

import (
	"testing"

	"github.com/LavaJover/dronwallet/auth/internal/middleware/password"
)

func TestCheckPassword(t *testing.T){
	rawPasswords := []string{"absdaj", "bodya", "bagir", "8Uha_012"}
	hashedPasswords := make([]string, len(rawPasswords))

	for i, rawPassword := range rawPasswords{
		hashedPasswords[i] = password.HashPassword(rawPassword)
	}

	for i, hashedPassword := range hashedPasswords{
		if !password.CheckPassword(rawPasswords[i], hashedPassword){
			t.Error("Incorrect result!")
		}
	}

}