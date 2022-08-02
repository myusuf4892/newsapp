package middlewares

import (
	"errors"
	_config "news/databases/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(_config.JWT()),
	})
}
func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix() //Token expires after 5 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(_config.JWT()))
}

func ExtractToken(e echo.Context) (int, error) {
	token := e.Get("user")
	if token == nil {
		return 0, errors.New("not authorized")
	}
	user := token.(*jwt.Token)
	if user.Valid {
		claims := user.Claims.(jwt.MapClaims)
		userID := claims["userID"].(float64)
		return int(userID), nil
	}
	return 0, errors.New("invalid token")
}
