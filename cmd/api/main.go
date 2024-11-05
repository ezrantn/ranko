package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	app := application{
		config: config{
			addr: os.Getenv("PORT"),
		},
	}

	web := app.mount()

	if err := app.run(web); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
