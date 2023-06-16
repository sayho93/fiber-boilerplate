package src

import (
	"fiber/src/common"
	"fiber/src/common/database"
	"fiber/src/common/middlewares"
	"fiber/src/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"time"
)

var AppSet = wire.NewSet(
	NewApp,
	common.ConfigSet,
	database.DBSet,
	users.SetRepository,
	users.SetService,
	users.SetHandler,
)

func NewApp(config *common.Config, handler users.UserHandler) *fiber.App {
	app := fiber.New(config.Fiber)

	if !fiber.IsChild() {
		logrus.Debug("Master process init")
	} else {
		logrus.Debug("Child process init")
	}

	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))
	app.Use(helmet.New())
	//app.Use(csrf.New(config.Csrf))
	app.Use(requestid.New())
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(middlewares.LogMiddleware)

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

	users.NewRouter(v1, handler)

	return app
}
