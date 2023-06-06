package common

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type IConfig struct {
	Port            int
	MariadbHost     string
	MariadbUsername string
	MariadbPassword string
	MariadbDatabase string
	MariadbPort     string
}

var Config IConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port, parseErr := strconv.Atoi(os.Getenv("PORT"))
	if parseErr != nil {
		panic(parseErr)
	}

	Config = IConfig{
		Port:            port,
		MariadbHost:     os.Getenv("MARIADB_HOST"),
		MariadbUsername: os.Getenv("MARIADB_USERNAME"),
		MariadbPassword: os.Getenv("MARIADB_PASSWORD"),
		MariadbDatabase: os.Getenv("MARIADB_DATABASE"),
		MariadbPort:     os.Getenv("MARIADB_PORT"),
	}
}
