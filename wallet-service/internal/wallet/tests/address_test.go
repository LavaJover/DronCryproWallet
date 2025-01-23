package wallet_test

import (
	"testing"
	"fmt"

	"github.com/LavaJover/DronCryptoWallet/wallet-service/internal/wallet"
)

func TestGenerateTronAddressFromPrivateKey(t *testing.T){
	privateKeys := []string{
		"053008d68b698b83f3fda83fcc35ff1ddeda1d3a3a3ca13222f0a80540dc31bb",
		"e35d65f888bdef3941b98ebda483496e38c83b774137dfcd1580e7b0984044ce",
		"e35d65f888bdef3941b98ebda483496e38c83b774137dfcd1580e7b0984044ce",
		"6c3767ef998c3b24b03dd9c25dbd126449f53d0f6b812c1f8e6932d2736fec8e",
		"1e06189b76c0cca9f88a18f487da79d9b176a634a32cd38f84a8801a8fd8942e",
		"ea2c087586f50c7a5302dd6fa89b8cfd46bc2302fa06e4f2fa1ff9733adfe5d7",
	}

	expectedWalletAddresses := []string{
		"TBV6W7x1tg2HBv8cfpjsuAcZEGv61U9UBR",
		"TAucKjZuH5DBk6GNePMvgu5cuWXVC6cvRx",
		"TAucKjZuH5DBk6GNePMvgu5cuWXVC6cvRx",
		"TEzXJYUyKbbrBABVYuCA1XmS4o6PYNJT5r",
		"TFRqr69AeSbSAXVJtyTjgyw5hnTqdCpV3R",
		"TKiwwfF2CXR3nDRRaQhCkBFgYU1o4Fp5wH",
	}

	for i, pk := range privateKeys{
		walletAddress, err := wallet.GenerateTronAddress(pk)
		fmt.Println()
		if err != nil{
			t.Errorf("Failed to generate address: %v", err)
		}

		if expectedWalletAddresses[i] != walletAddress{
			t.Errorf("Wrong wallet address %s for pk %s", walletAddress, pk)
		}
		
	}
}