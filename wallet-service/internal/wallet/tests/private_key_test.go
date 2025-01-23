package wallet_test

import (
	"testing"
	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/wallet"
)

func TestGeneratePrivateKey(t *testing.T){
	for range 100{
		if _, err := wallet.GeneratePrivateKey(); err != nil{
			t.Error("failed to generate private key")
		}
	}
}