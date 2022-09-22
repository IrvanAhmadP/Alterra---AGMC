package middleware

import (
	res "agmc-day-6/pkg/util/response"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

	return func(c echo.Context) error {
		authToken := c.Request().Header.Get("Authorization")
		if authToken == "" {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, nil).Send(c)
		}

		if !strings.Contains(authToken, "Bearer ") {
			return res.CustomErrorBuilder(http.StatusBadRequest, "failed", "invalid token").Send(c)
		}

		tokenString := strings.Replace(authToken, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method :%v", token.Header["alg"])
			}

			return []byte(JWT_SECRET_KEY), nil
		})

		if !token.Valid || err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err).Send(c)
		}
		return next(c)
	}
}
