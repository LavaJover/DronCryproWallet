package password

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)


func HashPassword(rawPassword string) string{
	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)

	if err != nil{
		log.Fatalf("Failed to gen hash from password")
	}

	return string(hashedPassowrd)
}


func CheckPassword(rawPassword string, hashedPassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	return err == nil
}