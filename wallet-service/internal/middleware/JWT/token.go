package jwttoken

import(
	"fmt"
	
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("your-secret-key")

// Структура для пользовательских данных внутри токена
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func DecodeJWT(tokenString string) (uint, error) {
	// Расшифровываем токен
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи совпадает с ожиданиями
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return 0, err
	}

	// Извлекаем и возвращаем поле UserID из токена
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UserID, nil
	}

	return 0, fmt.Errorf("invalid token")
}