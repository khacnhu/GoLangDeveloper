package internal

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Reading environment variables
	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")

	db_ssl_mode := os.Getenv("SSL_MODE")
	db_port := os.Getenv("DB_PORT")
	db_timezone := os.Getenv("TIME_ZONE")

	// fake to test local database
	dsn := "host=localhost user=postgres password=Trankhacnhu132! dbname=godb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	fmt.Println("check dsn old = ", dsn)

	dsnTest := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, db_user, db_password, db_name, db_port, db_ssl_mode, db_timezone)
	// fmt.Println("check dsn new = ", dsnTest)

	db, err := gorm.Open(postgres.Open(dsnTest), &gorm.Config{})

	if err != nil {
		fmt.Println("Connect database failed huhuuuuuuuuuuuuuuu :(((")
		return nil

	}

	return db

}
