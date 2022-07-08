package seeder

import (
	"log"

	"github.com/Final-Project-Kelompok-3/users/internal/model"
	"gorm.io/gorm"
)

func roleSeeder(conn *gorm.DB) {

	var users = []model.Role{
		{Name: "Admin", Description: "Adminisrator"},
		{Name: "Siswa", Description: "Siswa pendaftar"},
		{Name: "Sekolah", Description: "Pihak sekolah"},
	}

	if err := conn.Create(&users).Error; err != nil {
		log.Printf("cannot seed data roles, with error %v\n", err)
	}
	log.Println("success seed data roles")
}