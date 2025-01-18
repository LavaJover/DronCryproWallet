package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/ripemd160"
)

// Base58 символы
const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Base58Check encoding
func encodeBase58(input []byte) string {
	var result []byte
	x := new(big.Int).SetBytes(input)

	// Преобразование в Base58
	base := big.NewInt(int64(len(base58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod)
		result = append([]byte{base58Alphabet[mod.Int64()]}, result...)
	}

	// Добавляем '1' для каждого ведущего нуля
	for _, b := range input {
		if b != 0 {
			break
		}
		result = append([]byte{base58Alphabet[0]}, result...)
	}

	return string(result)
}

// Base58Check encoding для Tron
func encode58Check(input []byte) string {
	// Добавляем хеш
	hash1 := sha256.Sum256(input)
	hash2 := sha256.Sum256(hash1[:])

	// Добавляем контрольные 4 байта
	extended := append(input, hash2[:4]...)

	// Кодируем в Base58
	return encodeBase58(extended)
}

// Генерация Tron-адреса
func generateTronAddress() (string, string, error) {
	// Генерация приватного ключа secp256k1
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate private key: %v", err)
	}

	// Приватный ключ в формате HEX
	privateKeyBytes := privateKey.D.Bytes()
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Публичный ключ
	publicKey := append(privateKey.PublicKey.X.Bytes(), privateKey.PublicKey.Y.Bytes()...)

	// RIPEMD160 (после SHA256)
	shaHash := sha256.Sum256(publicKey)
	ripemdHasher := ripemd160.New()
	_, _ = ripemdHasher.Write(shaHash[:])
	publicKeyHash := ripemdHasher.Sum(nil)

	// Добавляем префикс для адреса Tron (0x41)
	addressBytes := append([]byte{0x41}, publicKeyHash...)

	// Генерируем адрес в формате Base58Check
	address := encode58Check(addressBytes)

	return address, privateKeyHex, nil
}

func Run(){
	address, privateKey, err := generateTronAddress()
	if err != nil {
		log.Fatalf("Error generating Tron address: %v", err)
	}

	fmt.Printf("Generated Tron Address: %s\n", address)
	fmt.Printf("Private Key (HEX): %s\n", privateKey)
}