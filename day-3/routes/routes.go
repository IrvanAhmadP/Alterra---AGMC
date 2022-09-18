package routes

import (
	"agmc/controllers"
	"agmc/middlewares"

	"github.com/go-playground/validator/v10"

	"github.com/labstack/echo/v4"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}

func New() *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	book := e.Group("books")
	book.GET("", controllers.GetBooks)
	book.GET("/:bookID", controllers.GetBookByID)
	book.POST("", controllers.AddBook, middlewares.Auth)
	book.PUT("/:bookID", controllers.UpdateBook, middlewares.Auth)
	book.DELETE("/:bookID", controllers.DeleteBook, middlewares.Auth)

	user := e.Group("users")
	user.POST("/login", controllers.LoginUser)
	user.GET("", controllers.GetUsers, middlewares.Auth)
	user.GET("/:userID", controllers.GetUserByID, middlewares.Auth)
	user.POST("", controllers.AddUser)
	user.PUT("/:userID", controllers.UpdateUser, middlewares.Auth)
	user.DELETE("/:userID", controllers.DeleteUser, middlewares.Auth)

	return e
}
