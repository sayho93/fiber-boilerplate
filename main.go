package main

import (
	"fiber/src"
	"log"
	"os"
)

func init() {
	//TODO
}

func main() {
	server, _ := src.New()
	port := os.Getenv("PORT")
	log.Fatal(server.Listen(port))
}
