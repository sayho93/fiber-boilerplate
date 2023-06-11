package middlewares

import (
	"errors"
	Errors "fiber/src/common/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"os"
)

var GeneralErrorHandler = func(ctx *fiber.Ctx, err error) error {
	logrus.SetOutput(os.Stderr)
	code := fiber.StatusInternalServerError

	var exception *Errors.Error
	if errors.As(err, &exception) {
		code = exception.Code
	}

	logrus.Errorf("%+v", exception)
	logrus.SetOutput(colorable.NewColorableStdout())

	return ctx.Status(code).JSON(exception)
}
