package seeder

import "github.com/Final-Project-Kelompok-3/users/database"


func Seed() {

	conn := database.GetConnection()

	roleSeeder(conn)
	// otherTableSeeder(conn)
}