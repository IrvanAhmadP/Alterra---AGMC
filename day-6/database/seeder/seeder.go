package seeder

import "agmc-day-6/database"

func Seed() {
	conn := database.GetConnection()

	bookTableSeeder(conn)
	userTableSeeder(conn)
}
