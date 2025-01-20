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
	uservalid "github.com/LavaJover/DronCryptoWallet/api-gateway/validation/user"
	authpb "github.com/LavaJover/DronCryptoWallet/auth-service/proto/gen"
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

	authServiceConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	defer authServiceConn.Close()
	authServiceClient := authpb.NewAuthClient(authServiceConn)

	// Ручка регистрации нового пользователя
	http.HandleFunc("/api/auth/reg", func (w http.ResponseWriter, r *http.Request){
		if r.Method != http.MethodPost{
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
			return
		}
	
		var user models.User
	
		err := json.NewDecoder(r.Body).Decode(&user)
	
		if err != nil{
			http.Error(w, "Ошибка при парсинге JSON: " + err.Error(), http.StatusBadRequest)
			return
		}
	
		if err := uservalid.ValidateUserRequest(&user); err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		response, err := authServiceClient.Register(context.Background(), &authpb.RegisterRequest{
			Email: user.Email,
			Password: user.Password,
		})

		if err != nil{
			http.Error(w, "Ошибка при соединении с AuthService", http.StatusBadRequest)
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(response)
	
	})

	// Ручка логина пользователя
	http.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost{
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

		json.NewEncoder(w).Encode(response)

	})


	// Запуск сервера
	log.Println("API Gateway запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}