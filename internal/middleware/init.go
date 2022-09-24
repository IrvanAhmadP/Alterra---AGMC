package middleware

import (
	"agmc-day-6/pkg/util/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) {
	e.Use(
		middleware.Recover(),
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, error=${error} method=${method}, uri=${uri}, status=${status}\n",
		}),
	)

	e.HTTPErrorHandler = ErrorHandler
	e.Validator = &validator.CustomValidator{Validator: validator.NewValidator()}
}
