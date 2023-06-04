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

func JWTAdmin() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET_ADMIN")),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helper.JwtCustomAdminClaims)
		},
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			return jwt.ParseWithClaims(auth, new(helper.JwtCustomAdminClaims), func(token *jwt.Token) (interface{}, error) {
				if token.Claims.(*helper.JwtCustomAdminClaims).ExpiresAt.Time.Before(time.Now()) {
					return nil, fmt.Errorf("token expired")
				}
				return []byte(os.Getenv("JWT_SECRET_ADMIN")), nil
			})
		},
		ContextKey: "admin",
		SuccessHandler: func(c echo.Context) {
			data := c.Get("admin").(*jwt.Token).Claims.(*helper.JwtCustomAdminClaims)

			c.Set("admin", data)

		},
		ErrorHandler: func(c echo.Context, err error) error {
			if err.Error() == "token expired" {
				return c.JSON(http.StatusUnauthorized, helper.ResponseData(http.StatusUnauthorized, err.Error(), nil))
			}
			return c.JSON(http.StatusUnauthorized, helper.ResponseData(http.StatusUnauthorized, "invalid token", nil))
		},
	})
}