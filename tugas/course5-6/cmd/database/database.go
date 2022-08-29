package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func NewConnection() *gorm.DB {
	// load .env file
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		// log.Fatalf("Error loading .env file")
		panic(errEnv)
	}

	env := os.Getenv("ENV")

	fmt.Println("ENV : " + env)
	var dsn string = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=30",
		os.Getenv("NAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DB_NAME"),
	)

	fmt.Println("sql string : " + dsn)
	var err error
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
