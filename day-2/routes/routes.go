package routes

import (
	"agmc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
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
