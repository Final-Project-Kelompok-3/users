package database

import (
	"os"
	"sync"

	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB

	// sync.Once digunakan untuk menghindari konfigurasi database yang berulang kali
	once   sync.Once
)

func CreateConnection() {
	
	conf := dbConfig{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName: os.Getenv("DB_NAME"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	
	// jika menggunakan postgres, anda dapat uncomment kode di bawah ini
	//postgres := postgresqlConfig{dbConfig: conf}

	once.Do(func() {
		mysql.Connect()
		//postgres.Connect()
	})
}
	
func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}