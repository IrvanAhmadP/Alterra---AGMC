package seeder

import (
	"agmc-day-6/internal/model"
	"log"

	"gorm.io/gorm"
)

func bookTableSeeder(conn *gorm.DB) {
	var books = []model.Book{
		{Title: "The Power of Habit", Author: "Charles Duhigg", Year: "2012"},
		{Title: "Atomic Habits", Author: "James Clear", Year: "2018"},
		{Title: "Biografi Utsman Bin Affan", Author: "Prof. Dr. Ali Muhammad Ash-Shalabi", Year: "2017"},
	}

	if err := conn.Create(&books).Error; err != nil {
		log.Printf("cannot seed data books, with error %v\n", err)
	}
	log.Printf("success seed data books")
}
