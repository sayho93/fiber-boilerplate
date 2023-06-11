package common

import (
	"fiber/src/common/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/google/wire"
	"github.com/mattn/go-colorable"
	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
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
	Port  int
	Fiber fiber.Config
	DB    DB
	Csrf  csrf.Config
}

func fiberConfig() fiber.Config {
	return fiber.Config{
		//Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Fiber v1",
		ErrorHandler:  middlewares.GeneralErrorHandler,
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

type LumberjackHook struct {
	Writer *lumberjack.Logger
}

func (hook *LumberjackHook) Fire(entry *log.Entry) error {
	msg, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(msg))
	return err
}

func (hook *LumberjackHook) Levels() []log.Level {
	return log.AllLevels
}

func NewLumberjackHook(writer *lumberjack.Logger) *LumberjackHook {
	return &LumberjackHook{
		Writer: writer,
	}
}

func loggerConfig() {
	logPath := "./logs/logs.log"
	maxSize := 100
	maxBackups := 90
	maxAge := 1

	logRotation := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    maxSize,    // 파일 최대 크기 (MB)
		MaxBackups: maxBackups, // 보관할 백업 파일의 최대 개수
		MaxAge:     maxAge,     // 보관 기간 (일)
		Compress:   true,       // 압축 여부
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})
	log.SetLevel(log.DebugLevel)
	//log.SetReportCaller(true)
	log.SetOutput(colorable.NewColorableStdout())

	logHook := NewLumberjackHook(logRotation)
	log.AddHook(logHook)
}

func csrfConfig() csrf.Config {
	return csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "csrf_",               // name of the session cookie
		CookieSameSite: "Lax",                 // indicates if CSRF cookie is requested by SameSite
		Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUID,            // creates a new CSRF token
	}
}

func NewConfig() *Config {
	port, parseErr := strconv.Atoi(os.Getenv("PORT"))
	if parseErr != nil {
		panic(parseErr)
	}

	loggerConfig()

	var config = Config{
		Port:  port,
		Fiber: fiberConfig(),
		DB:    dbConfig(),
		Csrf:  csrfConfig(),
	}

	return &config
}

var ConfigSet = wire.NewSet(NewConfig)
