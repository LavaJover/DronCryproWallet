package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"log/slog"
	"net/http"
	"os"
	"bytes"

	"github.com/LavaJover/DronCryptoWallet/api-gateway/internal/config"
	"github.com/LavaJover/DronCryptoWallet/api-gateway/models"
	authpb "github.com/LavaJover/DronCryptoWallet/auth-service/proto/gen"
	walletpb "github.com/LavaJover/DronCryptoWallet/wallet-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO:
// - Затестить пакет validator

const(
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"
)

func setupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env{
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func main(){

	cfg := config.MustLoad()
	fmt.Println(cfg)

	myLog := setupLogger(cfg.Env)
	myLog.Info("starting api-gateway")

	// Auth service init
	authServiceConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	defer authServiceConn.Close()
	authServiceClient := authpb.NewAuthClient(authServiceConn)

	// Wallet service init
	walletServiceConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect to wallet-service: %v", err)
	}

	defer walletServiceConn.Close()
	walletServiceClient := walletpb.NewWalletServiceClient(walletServiceConn)

	// Ручка регистрации нового пользователя
	http.HandleFunc("/api/auth/reg", func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		myLog.Info("/api/auth/reg", "method", r.Method)
	
		// Handle OPTIONS requests
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}
	
		// Ensure the request is POST
		if r.Method != http.MethodPost {
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}
	
		// Parse and decode the JSON payload
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "Ошибка при парсинге JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
	
		// Send registration request to AuthService
		response, err := authServiceClient.Register(context.Background(), &authpb.RegisterRequest{
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "application/json")
			http.Error(w, "Ошибка при соединении с AuthService: "+err.Error(), http.StatusBadRequest)
			return
		}
	
		// Respond with success
		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
	

	// Ручка логина пользователя
	http.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {

		myLog.Info("/api/auth/login", "method", r.Method)

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost{
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}

		var user models.User
	
		err := json.NewDecoder(r.Body).Decode(&user)
	
		if err != nil{
			http.Error(w, "Ошибка при парсинге JSON: " + err.Error(), http.StatusBadRequest)
			return
		}

		response, err := authServiceClient.Login(context.Background(), &authpb.LoginRequest{
			Email: user.Email,
			Password: user.Password,
		})

		if err != nil{
			http.Error(w, "login failed" + err.Error(), http.StatusBadRequest)
		}

		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	})

	// Ручка получения приватного ключа
	http.HandleFunc("/api/wallet/privatekey", func(w http.ResponseWriter, r *http.Request) {

		myLog.Info("/api/wallet/privatekey", "method", r.Method)

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodGet{
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}

		response, err := walletServiceClient.GetPrivateKey(context.Background(), &walletpb.GetPrivateKeyRequest{Token: "some_token"})

		if err != nil{
			myLog.Error("failed to generate private key", "error", err)
			http.Error(w, "failed to generate private key", http.StatusNotFound)
		}

		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	// Ручка создания нового кошелька пользователя

	// Ручка проверки баланса кошелька пользователя

	// Ручка получения списка транзакций кошелька пользоватея

	// Ручка создания транзакции TRX кошелька пользователя

	// Ручка получения списка кошельков пользователя
	http.HandleFunc("/api/wallet/process-user-wallets", func(w http.ResponseWriter, r *http.Request) {
		myLog.Info("/api/wallet/process-user-wallets", "method", r.Method)

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Чтение тела запроса
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		type RequestPayload struct{
			Token string `json:"token"`
		}

		// Парсим входящий запрос
		var requestPayload RequestPayload
		if err := json.Unmarshal(body, &requestPayload); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		response, err := walletServiceClient.GetUserPrivateKeys(context.Background(), &walletpb.GetUserPrivateKeysRequest{Token: requestPayload.Token})

		if err != nil{
			myLog.Error("failed to get user private keys")
			http.Error(w, "Failed to get user private keys", http.StatusBadRequest)
			return
		}

		privateKeys := response.PrivateKeys

		// Структура запроса
		type WalletRequest struct {
			PrivateKeys []string `json:"privateKeys"`
		}
		
		// Структура для получения ответа от /generate-wallets
		type WalletResponse struct {
			Addresses []struct {
				PrivateKey string `json:"privateKey"`
				Address    string `json:"address"`
				Error      string `json:"error,omitempty"`
			} `json:"addresses"`
		}

		walletRequest := WalletRequest{
			PrivateKeys: privateKeys,
		}
		requestBody, err := json.Marshal(walletRequest)
		if err != nil {
			http.Error(w, "Failed to marshal request", http.StatusInternalServerError)
			return
		}

		resp, err := http.Post("http://localhost:3000/generate-addresses", "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			http.Error(w, "Failed to send request to wallet service", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Чтение ответа от /generate-wallets
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response from wallet service", http.StatusInternalServerError)
			return
		}

		// Парсим ответ от wallet service
		var walletResponse WalletResponse
		if err := json.Unmarshal(respBody, &walletResponse); err != nil {
			myLog.Error("failed to get response from node wallet service: " + err.Error())
			http.Error(w, "Invalid response from wallet service", http.StatusInternalServerError)
			return
		}

		// Сохраняем список адресов
		var walletAddresses []string
		for _, wallet := range walletResponse.Addresses {
			if wallet.Error == "" {
				walletAddresses = append(walletAddresses, wallet.Address)
			} else {
				log.Printf("Error generating address for private key %s: %s", wallet.PrivateKey, wallet.Error)
			}
		}

		// Возвращаем список адресов в ответ
		http_response := map[string]interface{}{
			"walletAddresses": walletAddresses,
		}


		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(http_response)
	})


	// Запуск сервера
	myLog.Info("api gateway serving", "address", cfg.Address)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}