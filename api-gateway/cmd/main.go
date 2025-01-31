package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	// "strings"

	"github.com/LavaJover/DronCryptoWallet/api-gateway/internal/config"
	"github.com/LavaJover/DronCryptoWallet/api-gateway/models"
	authpb "github.com/LavaJover/DronCryptoWallet/auth-service/proto/gen"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TODO:
// - Затестить пакет validator

// @title Simple Go API
// @version 1.0
// @description Example of a simple HTTP server with Swagger documentation.
// @host localhost:8080
// @BasePath /api/v1

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)

	myLog := setupLogger(cfg.Env)
	myLog.Info("starting api-gateway")

	// Auth service init
	authServiceConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	defer authServiceConn.Close()
	authServiceClient := authpb.NewAuthClient(authServiceConn)

	// @Summary Get user by ID
	// @Description Returns a user by their ID.
	// @Tags users
	// @Accept json
	// @Produce json
	// @Param id query int true "User ID"
	// @Success 200 {object} User
	// @Failure 404 {object} map[string]string
	// @Router /users [get]
	// Ручка регистрации нового пользователя
	http.HandleFunc("/api/auth/reg", func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		myLog.Info(r.URL.Path, "method", r.Method)

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
		myLog.Info(r.URL.Path, "status", http.StatusCreated)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	})

	// Ручка логина пользователя
	http.HandleFunc("/api/auth/login", func(w http.ResponseWriter, r *http.Request) {

		myLog.Info(r.URL.Path, "method", r.Method)

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}

		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)

		if err != nil {
			http.Error(w, "Ошибка при парсинге JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		response, err := authServiceClient.Login(context.Background(), &authpb.LoginRequest{
			Email:    user.Email,
			Password: user.Password,
		})

		if err != nil {
			http.Error(w, "login failed"+err.Error(), http.StatusBadRequest)
		}

		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	})

	// Ручка валидации токена
	http.HandleFunc("/api/auth/valid", func(w http.ResponseWriter, r *http.Request) {

		myLog.Info(r.URL.Path, "method", r.Method)

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		}

		if r.Method != http.MethodPost {
			myLog.Error("method is not supported", "method", r.Method)
			http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует заголовок Authorization", http.StatusUnauthorized)
			return
		}

		// Проверяем, что токен имеет префикс "Bearer "
		const bearerPrefix = "Bearer "
		if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			http.Error(w, "Некорректный формат токена", http.StatusUnauthorized)
			return
		}

		// Извлекаем сам токен
		token := authHeader[len(bearerPrefix):]

		myLog.Info("JWT", "token", token)

		validateResponse, err := authServiceClient.ValidateJWT(context.Background(), &authpb.ValidateJWTRequest{Token: token})

		if err != nil {
			myLog.Error("invalid JWT token", "err", err.Error())
			http.Error(w, "invalid JWT token", http.StatusBadRequest)
			return
		}

		myLog.Info(r.URL.Path, "status", http.StatusOK)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(validateResponse)

	})

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Запуск сервера
	myLog.Info("api gateway serving", "address", cfg.Address)
	if err := http.ListenAndServe(cfg.Address, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
