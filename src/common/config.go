package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/wire"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type DB struct {
	MariadbHost     string
	MariadbUsername string
	MariadbPassword string
	MariadbDatabase string
	MariadbPort     string
}

type Config struct {
	Port   int
	Fiber  fiber.Config
	DB     DB
	Csrf   csrf.Config
	Logger logger.Config
}

func fiberConfig() fiber.Config {
	return fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber v1",
	}
}

func dbConfig() DB {
	return DB{
		MariadbHost:     os.Getenv("MARIADB_HOST"),
		MariadbUsername: os.Getenv("MARIADB_USERNAME"),
		MariadbPassword: os.Getenv("MARIADB_PASSWORD"),
		MariadbDatabase: os.Getenv("MARIADB_DATABASE"),
		MariadbPort:     os.Getenv("MARIADB_PORT"),
	}
}

func loggerConfig() logger.Config {
	file, err := os.OpenFile("./logs/my_logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	return logger.Config{
		Output: file, // add file to save output
	}
}

func csrfConfig() csrf.Config {
	return csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "my_csrf_",            // name of the session cookie
		CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
		Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUID,            // creates a new CSRF token
	}
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port, parseErr := strconv.Atoi(os.Getenv("PORT"))
	if parseErr != nil {
		panic(parseErr)
	}

	var config = Config{
		Port:   port,
		Fiber:  fiberConfig(),
		DB:     dbConfig(),
		Csrf:   csrfConfig(),
		Logger: loggerConfig(),
	}

	return &config
}

var ConfigSet = wire.NewSet(NewConfig)
