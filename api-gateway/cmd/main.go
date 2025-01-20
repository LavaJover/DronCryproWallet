package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/LavaJover/DronCryptoWallet/api-gateway/models"
	uservalid "github.com/LavaJover/DronCryptoWallet/api-gateway/validation/user"
	authpb "github.com/LavaJover/DronCryptoWallet/auth-service/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


func main(){
	authServiceConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil{
		log.Fatalf("Failed to connect to auth-service: %v", err)
	}

	defer authServiceConn.Close()
	authServiceClient := authpb.NewAuthClient(authServiceConn)

	
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

	// Запуск сервера
	log.Println("API Gateway запущен на порту :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}