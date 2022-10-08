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
	newBookJSON    string = `{"title":"Book 1","summary":"Book about economy","author":"Author 1","category":"economy","year":2000}`
	updateBookJSON string = `{"title": "Book updated"}`
)

func setupTestBookAPI(method string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, "/users", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestGetBooks(t *testing.T) {
	// Setup
	c, rec := setupTestBookAPI(http.MethodGet, nil)

	// Assertions
	if assert.NoError(t, GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetBookByID(t *testing.T) {
	// Setup
	c, rec := setupTestBookAPI(http.MethodGet, nil)
	c.SetPath(":bookID")
	c.SetParamNames("bookID")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, GetBookByID(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAddBook(t *testing.T) {
	// Setup
	c, _ := setupTestBookAPI(http.MethodPost, strings.NewReader(newBookJSON))

	// Assertions
	assert.NoError(t, AddBook(c))
	// assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestUpdateBook(t *testing.T) {
	// Setup
	c, rec := setupTestBookAPI(http.MethodPut, strings.NewReader(updateBookJSON))
	c.SetPath(":bookID")
	c.SetParamNames("bookID")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, UpdateBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestDeleteBook(t *testing.T) {
	// Setup
	c, rec := setupTestBookAPI(http.MethodDelete, strings.NewReader(updateBookJSON))
	c.SetPath(":bookID")
	c.SetParamNames("bookID")
	c.SetParamValues("1")

	// Assertions
	if assert.NoError(t, DeleteBook(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
