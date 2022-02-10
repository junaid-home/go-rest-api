package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"api/src/router"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}
}

func main() {
	router := router.RegisterRoutes()

	fmt.Printf("[ OK ] Server is Started and Listening on port: %v", PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
