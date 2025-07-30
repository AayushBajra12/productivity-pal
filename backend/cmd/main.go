package main

import (
	"log"
	"productivity-pal/backend/internal/db"
	"productivity-pal/backend/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Shutting down application, unable to load env variables: ", err)
	}

	err = db.InitDB()
	if err != nil {
		log.Fatal("Shutting down application, unable to initialize db: ", err)
	}

	err = server.StartServer()
	if err != nil {
		log.Fatal("Shutting down application, unable to start the server: ", err)
	}
}
