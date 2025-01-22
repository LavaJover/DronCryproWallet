package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

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
		myLog.Info("register handler", "URL", "/api/auth/reg")
	
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
			http.Error(w, "Ошибка при парсинге JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
	
		// Send registration request to AuthService
		response, err := authServiceClient.Register(context.Background(), &authpb.RegisterRequest{
			Email:    user.Email,
			Password: user.Password,
		})
		if err != nil {
			http.Error(w, "Ошибка при соединении с AuthService: "+err.Error(), http.StatusBadRequest)
			return
		}
	
		// Respond with success
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})
	

	// Ручка логина пользователя
	http.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {

		myLog.Info("login handler", "URL", "/api/auth/login")

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

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	})

	http.HandleFunc("/api/wallet/balance", func(w http.ResponseWriter, r *http.Request) {
		response, err := walletServiceClient.GetWalletBalance(context.Background(), &walletpb.GetBalanceRequest{
			Address: "TQUurKqa9dpZcoCV7QwdQMJxueycFipFbh",
		})

		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(response)
	})

	// Запуск сервера
	myLog.Info("api gateway serving", "address", cfg.Address)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}