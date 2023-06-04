package helper

import (
	"errors"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/constant"
	"github.com/golang-jwt/jwt/v4"
)

type AuthJWT interface {
	GenerateUserToken(id string, email string, username string, authBy constant.AuthBy) (string, error)
	GenerateAdminToken(email string) (string, error)
	// IsAdmin(token *jwt.Token) (error)
}

type authJWT struct {
	secretUser string
	secretAdmin string
}

func NewAuthJWT(secretUser, secretAdmin string) AuthJWT {
	return &authJWT{
		secretUser: secretUser,
		secretAdmin: secretAdmin,
	}
}

func (aj *authJWT) GetUserSecret() string {
	return aj.secretUser
}

func (aj *authJWT) GetAdminSecret() string {
	return aj.secretAdmin
}

func (aj *authJWT) GenerateUserToken(id string, email string, username string, authBy constant.AuthBy) (string, error) {
	claims := &JwtCustomUserClaims{
		id,
		email,
		username,
		authBy,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(aj.secretUser))
}

type JwtCustomUserClaims struct {
	ID     string          `json:"id"`
	Email  string          `json:"email"`
	Username string        `json:"username"`
	AuthBy constant.AuthBy `json:"auth_by"`
	jwt.RegisteredClaims
}

func (aj *authJWT) GenerateAdminToken(email string) (string, error) {
	claims := &JwtCustomAdminClaims{
		email,
		"admin",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	return token.SignedString([]byte(aj.secretAdmin))
}

type JwtCustomAdminClaims struct {
	Email   string `json:"email"`
	Username string        `json:"username"`	
	jwt.RegisteredClaims
}

var (
	ErrInvalidCredential = errors.New("invalid credential")
)