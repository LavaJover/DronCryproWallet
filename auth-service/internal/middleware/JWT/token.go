package token

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your-secret-key") // Храните ключ в переменной окружения!

// Claims — структура для данных токена.
type Claims struct {
    UserID int `json:"user_id"`
    jwt.RegisteredClaims
}

// GenerateJWT — функция для генерации токена.
func GenerateJWT(userID int) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour) // Токен будет действителен 1 час

    claims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    // Создаём токен с подписью
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// ValidateJWT — проверка токена.
func ValidateJWT(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtSecret, nil
    })

    if err != nil || !token.Valid {
        return nil, err
    }
    return claims, nil
}
