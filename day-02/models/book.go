package models

type Book struct {
	ID       int64  `json:"id" form:"id"`
	Title    string `json:"title" form:"title"`
	Summary  string `json:"summary" form:"summary"`
	Author   string `json:"author" form:"author"`
	Category string `json:"category" form:"category"`
	Year     int    `json:"year" form:"year"`
}
