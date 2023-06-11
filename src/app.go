package src

import (
	"encoding/json"
	"fiber/src/common"
	"fiber/src/common/database"
	"fiber/src/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"strings"
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

	app.Use(func(c *fiber.Ctx) error {
		logrus.Info(c)
		queryParams := c.Request().URI().QueryArgs().String()
		if strings.EqualFold(queryParams, "") {
			return c.Next()
		}
		logrus.Infof("Query: %s", queryParams)

		var prettyBodyParams interface{}
		if c.Body() == nil {
			return c.Next()
		}
		err := json.Unmarshal(c.Body(), &prettyBodyParams)
		if err != nil {
			logrus.Errorf("Failed to unmarshal body parameters: %v", err)
			return c.Next()
		}

		prettyBody, errIndent := json.MarshalIndent(prettyBodyParams, "", "  ")
		if errIndent != nil {
			logrus.Errorf("Failed to marshal body parameters: %v", err)
			return c.Next()
		}

		logrus.Infof("Body: %s", string(prettyBody))
		return c.Next()
	})
	//app.Use(logger.New(config.Logger))

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
	v1.Post("/users", handler.CreateOne)
	v1.Get("/users", handler.FindMany)
	v1.Get("/users/:id", handler.FindOne)
	v1.Patch("/users/:id", handler.UpdateOne)
	v1.Delete("/users", handler.DeleteOne)

	return app
}
