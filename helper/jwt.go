package helper

import (
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/golang-jwt/jwt/v4"
)

type AuthJWT interface {
	GenerateToken(id string, email string, authBy constant.AuthBy) (string, error)
}

type authJWT struct {
	secret string
}

func NewAuthJWT(secret string) *authJWT {
	return &authJWT{
		secret: secret,
	}
}

func (aj *authJWT) GenerateToken(id string, email string, authBy constant.AuthBy) (string, error) {
	claims := &JwtCustomClaims{
		id,
		email,
		authBy,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(aj.secret))
}

type JwtCustomClaims struct {
	Id     string          `json:"id"`
	Email  string          `json:"email"`
	AuthBy constant.AuthBy `json:"auth_by"`
	jwt.RegisteredClaims
}
