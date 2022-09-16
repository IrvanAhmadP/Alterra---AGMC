package routes

import (
	"agmc/controllers"

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
	v1 := e.Group("v1/")

	book := v1.Group("books")
	book.GET("", controllers.GetBooks)
	book.GET("/:bookID", controllers.GetBookByID)
	book.POST("", controllers.AddBook)
	book.PUT("/:bookID", controllers.UpdateBook)
	book.DELETE("/:bookID", controllers.DeleteBook)

	user := v1.Group("users")
	user.GET("", controllers.GetUsers)
	user.GET("/:userID", controllers.GetUserByID)
	user.POST("", controllers.AddUser)
	user.PUT("/:userID", controllers.UpdateUser)
	user.DELETE("/:userID", controllers.DeleteUser)

	return e
}
