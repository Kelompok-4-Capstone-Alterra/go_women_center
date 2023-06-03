package helper

import (
	"errors"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/golang-jwt/jwt/v4"
)

type AuthJWT interface {
	GenerateUserToken(id string, email string, authBy constant.AuthBy) (string, error)
	GenerateAdminToken(email string) (string, error)
	IsAdmin(token *jwt.Token) (error)
}

type authJWT struct {
	secret string
}

func NewAuthJWT(secret string) *authJWT {
	return &authJWT{
		secret: secret,
	}
}

func (aj *authJWT) GetSecret() string {
	return aj.secret
}

func (aj *authJWT) GenerateUserToken(id string, email string, authBy constant.AuthBy) (string, error) {
	claims := &JwtCustomUserClaims{
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

type JwtCustomUserClaims struct {
	Id     string          `json:"id"`
	Email  string          `json:"email"`
	AuthBy constant.AuthBy `json:"auth_by"`
	jwt.RegisteredClaims
}

func (aj *authJWT) GenerateAdminToken(email string) (string, error) {
	claims := &JwtCustomAdminClaims{
		email,
		true,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(aj.secret))
}

type JwtCustomAdminClaims struct {
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

func (aj *authJWT) IsAdmin(token *jwt.Token) (error) {
	claims := token.Claims.(jwt.MapClaims)
	idPayload, ok := claims["is_admin"].(bool)
	if !ok || !idPayload {
		return ErrInvalidCredential
	}
	return nil
}
var (
	ErrInvalidCredential = errors.New("invalid credential")
)