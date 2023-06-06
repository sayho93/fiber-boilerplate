package main

import (
	"fiber/src/common"
	"fiber/src/users"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber v1",
	})

	app.Static("/static", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/users", users.FindMany)
	v1.Get("/users/:id", users.FindOne)
	v1.Post("users", users.CreateOne)
	v1.Patch("/users", users.UpdateOne)
	v1.Delete("/users", users.DeleteOne)

	log.Fatal(app.Listen(fmt.Sprintf("localhost:%d", common.Config.Port)))
}
