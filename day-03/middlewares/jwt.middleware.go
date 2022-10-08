package middlewares

import (
	"agmc/lib/helpers"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Claims struct {
	UserID int
	jwt.StandardClaims
}

var LOGIN_DURATION = time.Duration(1) * time.Hour
var JWT_SECRET_KEY = []byte(helpers.GetSecretKeyConfig())
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		_, err := ExtractTokenUserID(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"status":  "failed",
				"message": err.Error(),
			})
		}
		return next(c)
	}
}

func CreateToken(userID int) (string, error) {
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(LOGIN_DURATION).Unix(),
		},
	}
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	return token.SignedString(JWT_SECRET_KEY)
}

func ExtractTokenUserID(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != JWT_SIGNING_METHOD {
			return nil, fmt.Errorf("signing method invalid")
		}
		return JWT_SECRET_KEY, nil
	})
	if err != nil {
		return 0, err
	}

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userID := claims["UserID"]
		return userID, nil
	}
	return 0, fmt.Errorf("invalid token")
}
