package controllers

import (
	"agmc/lib/database"
	"agmc/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByID(c echo.Context) error {
	id := c.Param("userID")

	user, err := database.GetUserByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func AddUser(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := database.AddUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User saved",
	})
}

func UpdateUser(c echo.Context) error {
	var user models.User
	id := c.Param("userID")
	c.Bind(&user)

	err := database.UpdateUser(id, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "User updated",
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("userID")

	err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status":  "success",
		"message": "User deleted",
	})
}
