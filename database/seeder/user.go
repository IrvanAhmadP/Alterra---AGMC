package seeder

import (
	"agmc-day-6/internal/model"
	"log"

	"gorm.io/gorm"
)

func userTableSeeder(conn *gorm.DB) {
	var users = []*model.User{
		{Name: "Irvan A.", Username: "irvana", Email: "no-reply@gmail.com", Password: "1234567"},
	}

	for _, user := range users {
		user.HashPassword()
	}

	if err := conn.Create(&users).Error; err != nil {
		log.Printf("cannot seed data users, with error %v\n", err)
	}
	log.Printf("success seed data users")
}
