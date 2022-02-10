package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"api/src/helpers"
	"api/src/models"
	"api/src/router"
)

func init() {
	if envLoadError := godotenv.Load(); envLoadError != nil {
		log.Fatal("[ ERROR ] Failed to load .env file")
	}

}

func main() {
	var PORT string
	db := helpers.CreateDatabaseInstance()

	router := router.RegisterRoutes(db)

	if migrateError := db.AutoMigrate(&models.Food{}, &models.User{}); migrateError != nil {
		log.Fatal("[ ERROR ] Couldn't migrate models!")
	}

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "9090"
	}

	fmt.Printf("[ OK ] Server is Started and Listening on port: %v", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
