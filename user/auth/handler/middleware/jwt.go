package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Kelompok-4-Capstone-Alterra/go_women_center/helper"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JWTUser() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_USER")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomUserClaims)
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return jwt.ParseWithClaims(auth, new(helper.JwtCustomUserClaims), func(token *jwt.Token) (interface{}, error) {
				if token.Claims.(*helper.JwtCustomUserClaims).ExpiresAt.Time.Before(time.Now()) {
					return nil, fmt.Errorf("token expired")
				}
				return []byte(os.Getenv("JWT_SECRET_USER")), nil
			})
		},
		SuccessHandler: func(c echo.Context) {
			data := c.Get("user").(*jwt.Token).Claims.(*helper.JwtCustomUserClaims)
			c.Set("user", data)
		},
		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Println(err.Error())
			if err.Error() == "token expired" {
				return c.JSON(http.StatusUnauthorized, helper.ResponseData(err.Error(), http.StatusUnauthorized, nil))
			}
			return c.JSON(http.StatusUnauthorized, helper.ResponseData("invalid token", http.StatusUnauthorized, nil))
		},
	})
}
