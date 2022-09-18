package models

type Book struct {
	ID       int64  `json:"id"`
	Title    string `json:"title" validate:"required"`
	Summary  string `json:"summary"`
	Author   string `json:"author" validate:"required"`
	Category string `json:"category" validate:"required"`
	Year     int    `json:"year" validate:"required"`
}
