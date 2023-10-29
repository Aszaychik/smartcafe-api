package jwt

import (
	"aszaychik/smartcafe-api/domain/web"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(adminLoginResponse *web.AdminLoginResponse, id uint) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["username"] = adminLoginResponse.Username
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}