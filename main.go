package main

import (
	"App/database"
	"App/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Init()

	route := routes.Loader()

	route.Run(os.Getenv("APP_PORT"))
}
