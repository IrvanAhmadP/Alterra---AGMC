package seeder

import (
	"agmc-day-6/database"

	"gorm.io/gorm"
)

type seed struct {
	DB *gorm.DB
}

func NewSeeder() *seed {
	return &seed{database.GetConnection()}
}

func (s *seed) SeedAll() {
	bookTableSeeder(s.DB)
	userTableSeeder(s.DB)
}

func (s *seed) DeleteAll() {
	s.DB.Exec("DELETE FROM books")
	s.DB.Exec("DELETE FROM users")
}
