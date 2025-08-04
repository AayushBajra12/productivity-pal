package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"

	"productivity-pal/backend/internal/ai/gemma"
	"productivity-pal/backend/internal/db"
	"productivity-pal/backend/internal/server"
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

func NewGemmaClient() *gemma.GemmaClient {
	baseURL := os.Getenv("GEMMA_BASE_URL")
	if baseURL == "" {
		baseURL = "http://gemma-cpu:11434"
	}
	return &gemma.GemmaClient{
		BaseURL: baseURL,
		HttpClient: &http.Client{
			Timeout: 100 * time.Second,
		},
	}
}
