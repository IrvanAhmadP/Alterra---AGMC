package controllers

import (
	"agmc/models"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

var books = []models.Book{
	{
		ID:       1,
		Title:    "Book 1",
		Summary:  "Book about economy",
		Author:   "Author 1",
		Category: "economy",
		Year:     2000,
	},
	{
		ID:       2,
		Title:    "Book 2",
		Summary:  "Book about economy",
		Author:   "Author 2",
		Category: "economy",
		Year:     2001,
	},
	{
		ID:       3,
		Title:    "Book 3",
		Summary:  "Book about economy",
		Author:   "Author 3",
		Category: "economy",
		Year:     2002,
	},
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   books,
	})
}

func GetBookByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("bookID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": "Book ID is invalid",
		})
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status": "success",
				"data":   book,
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status":  "failed",
		"message": "Book not found",
	})
}

func AddBook(c echo.Context) error {
	var book models.Book
	c.Bind(&book)

	if err := c.Validate(book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	book.ID = time.Now().Unix()
	books = append(books, book)

	return c.JSON(http.StatusCreated, map[string]string{
		"status":  "success",
		"message": "Book saved",
	})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("bookID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": "Book ID is invalid",
		})
	}

	req := models.Book{}
	c.Bind(&req)

	for i, book := range books {
		if book.ID == id {
			if req.Title != "" {
				book.Title = req.Title
			}

			if req.Summary != "" {
				book.Summary = req.Summary
			}

			if req.Author != "" {
				book.Author = req.Author
			}

			if req.Year != 0 {
				book.Year = req.Year
			}

			books[i] = book

			return c.JSON(http.StatusOK, map[string]string{
				"status":  "success",
				"message": "Book updated",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status":  "failed",
		"message": "Book not found",
	})
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("bookID"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  "failed",
			"message": "Book ID is invalid",
		})
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			return c.JSON(http.StatusOK, map[string]string{
				"status":  "success",
				"message": "Book deleted",
			})
		}
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"status":  "failed",
		"message": "Book not found",
	})
}
