package wallet

import (
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

func GeneratePrivateKey() (string, error){
	privateKey, err := secp256k1.GeneratePrivateKey()
	if err != nil {
		return "", err
	}

	// Приватный ключ в виде байтов
	privateKeyBytes := privateKey.Serialize()

	return string(privateKeyBytes), nil
}