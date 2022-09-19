package controllers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	loginUserJSON  = `{"email":"test@gmail.com","password":"1234"}`
	newUserJSON    = `{"name":"Test Account","username":"test","email":"test@gmail.com","password":"1234"}`
	updateUserJSON = `{"name":"Name updated"}`
)

func setupTestUserAPI(method string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, "/users", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestLoginUser(t *testing.T) {
	c, rec := setupTestUserAPI(http.MethodPost, strings.NewReader(loginUserJSON))

	// Assertions
	if assert.NoError(t, LoginUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUsers(t *testing.T) {
	c, rec := setupTestUserAPI(http.MethodGet, nil)

	// Assertions
	if assert.NoError(t, GetUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetUserByID(t *testing.T) {
	c, rec := setupTestUserAPI(http.MethodGet, nil)
	c.SetPath(":userID")
	c.SetParamNames("userID")
	c.SetParamValues("25")

	// Assertions
	if assert.NoError(t, GetUserByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAddUser(t *testing.T) {
	c, _ := setupTestUserAPI(http.MethodPost, strings.NewReader(newUserJSON))

	// Assertions
	assert.NoError(t, AddUser(c))
	// assert.Equal(t, http.StatusOK, rec.Code)
}

func TestUpdateUser(t *testing.T) {
	c, rec := setupTestBookAPI(http.MethodPut, strings.NewReader(updateUserJSON))
	c.SetPath(":userID")
	c.SetParamNames("userID")
	c.SetParamValues("25")

	// Assertions
	if assert.NoError(t, UpdateUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteUser(t *testing.T) {
	c, rec := setupTestBookAPI(http.MethodDelete, nil)
	c.SetPath(":userID")
	c.SetParamNames("userID")
	c.SetParamValues("24")

	// Assertions
	if assert.NoError(t, DeleteUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
