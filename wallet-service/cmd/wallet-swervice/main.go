package main

import (
	"fmt"
	"log"

	"github.com/LavaJover/dronwallet/wallet/internal/wallet"
)

var(
	usdtAddress = "TXCwnDoSsAf1opVTitqdGxSKVE6uzP2DYN"
	apiKey = "b221365a-5a86-4d75-a1a3-1456c7f1864d"
)

func main(){
	balance, err := wallet.GetUSDTBalance(usdtAddress, apiKey)

	if err != nil{
		log.Fatalf("Error getting balance: %v", err)
	}

	fmt.Println(balance)
}