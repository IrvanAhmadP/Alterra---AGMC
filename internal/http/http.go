package http

import (
	"agmc-day-6/internal/app/auth"
	"agmc-day-6/internal/app/book"
	"agmc-day-6/internal/app/user"
	"agmc-day-6/internal/factory"

	"github.com/labstack/echo/v4"
)

func NewHttp(e *echo.Echo, f *factory.Factory) {
	auth.NewHandler(f).Route(e.Group("/auth"))
	book.NewHandler(f).Route(e.Group("/books"))
	user.NewHandler(f).Route(e.Group("/users"))
}
