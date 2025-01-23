package wallet

import (
	"encoding/hex"
	"fmt"
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/mr-tron/base58"
)

// GenerateTronAddress generates a Tron address from a private key
func GenerateTronAddress(privateKeyHex string) (string, error) {
	// Decode the private key from hex
	privateKeyBytes, err := hex.DecodeString(privateKeyHex)
	if err != nil {
		return "", fmt.Errorf("invalid private key: %v", err)
	}

	// Generate ECDSA private key
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate private key: %v", err)
	}

	// Generate the public key
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)

	// Compute Keccak-256 hash of the public key (excluding the first byte)
	hash := crypto.Keccak256(publicKeyBytes[1:])

	// Take the last 20 bytes of the hash
	addressBytes := hash[len(hash)-20:]

	// Add Tron MainNet address prefix (0x41)
	prefixedAddress := append([]byte{0x41}, addressBytes...)

	// Compute checksum (first 4 bytes of Keccak-256 hash of the prefixed address)
	checksum := crypto.Keccak256(prefixedAddress)[:4]

	// Append checksum to the address
	fullAddress := append(prefixedAddress, checksum...)

	// Encode to Base58
	address := base58.Encode(fullAddress)

	return address, nil
}