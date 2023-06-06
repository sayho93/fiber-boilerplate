package src

import (
	"fiber/src/common"
	"fiber/src/common/database"
	"fiber/src/users"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/wire"
	"time"
)

var AppSet = wire.NewSet(
	NewApp,
	common.ConfigSet,
	database.DBSet,
	users.SetRepository,
	users.SetService,
)

func NewApp(config *common.Config, service *users.UserService) *fiber.App {
	app := fiber.New(config.Fiber)

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(helmet.New())
	app.Use(csrf.New(config.Csrf))
	app.Use(logger.New())

	app.Static("/static", "./public", fiber.Static{
		Compress:      true,
		ByteRange:     true,
		Browse:        true,
		Index:         "index.html",
		CacheDuration: 10 * time.Second,
		MaxAge:        3600,
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Get("/users", service.FindMany)
	v1.Get("/users/:id", service.FindOne)
	v1.Post("users", service.CreateOne)
	v1.Patch("/users", service.UpdateOne)
	v1.Delete("/users", service.DeleteOne)

	return app
}
