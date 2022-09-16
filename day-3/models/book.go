package models

type Book struct {
	ID       int64  `json:"id" form:"id"`
	Title    string `json:"title" form:"title" validate:"required"`
	Summary  string `json:"summary" form:"summary"`
	Author   string `json:"author" form:"author" validate:"required"`
	Category string `json:"category" form:"category" validate:"required"`
	Year     int    `json:"year" form:"year" validate:"required"`
}
