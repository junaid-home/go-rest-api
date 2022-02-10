package helpers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabaseInstance() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("[ ERROR ] Unable to connect with mysql!\n", err)
	}

	fmt.Println("[ OK ] Connected to the DB!")

	return db
}
