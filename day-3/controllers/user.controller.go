package controllers

import (
	"agmc/lib/database"
	"agmc/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	users, err := database.LoginUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

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

	if err := c.Validate(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
	}

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

	affectedRows, err := database.UpdateUser(id, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	message := fmt.Sprintf("%d users updated", affectedRows)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": message,
	})
}

func DeleteUser(c echo.Context) error {
	id := c.Param("userID")

	affectedRows, err := database.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	message := fmt.Sprintf("%d users deleted", affectedRows)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": message,
	})
}
