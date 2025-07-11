package config

import (
	"fmt"
	"log"
	database "main/connection/db/mongo"

	"github.com/joho/godotenv"
)

func Execute() {
	loadEnvironment()
	database.ConnectionToMongoDB()

	fmt.Println("[Config] Loaded Succesful")
}

func loadEnvironment() {
	err := godotenv.Load("../../config/.env")
	if err != nil {
		log.Fatal("Error to load .env file", err)
	}
}
