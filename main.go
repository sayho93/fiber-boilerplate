package main

import (
	"fiber/src"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func main() {
	server, _ := src.New()
	port := os.Getenv("PORT")
	log.Fatal(server.Listen(fmt.Sprintf("localhost:%s", port)))
}
