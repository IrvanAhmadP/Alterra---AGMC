package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title" gorm:"size:200;not null"`
	Author string `json:"author" gorm:"not null"`
	Year   string `json:"year" gorm:"not null"`
}
