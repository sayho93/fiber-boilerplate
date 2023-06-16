package users

import (
	"github.com/gofiber/fiber/v2"
)

func NewRouter(router fiber.Router, handler UserHandler) {
	router.Post("/users", handler.CreateOne)
	router.Get("/users", handler.FindMany)
	router.Get("/users/:id", handler.FindOne)
	router.Patch("/users/:id", handler.UpdateOne)
	router.Delete("/users", handler.DeleteOne)
}
