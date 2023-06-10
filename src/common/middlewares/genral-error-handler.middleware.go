package middlewares

//
//import (
//	"errors"
//	"github.com/gofiber/fiber/v2"
//)
//
//type CustomError struct {
//	Msg string
//	File  string
//	Line int
//}
//
//var GeneralErrorHandler = func(ctx *fiber.Ctx, err error) error {
//	code := fiber.StatusInternalServerError
//
//	var exception *fiber.Error
//	if errors.As(err, &exception) {
//		code = exception.Code
//	}
//
//	response := CustomError{
//		Msg: err.
//	}
//
//	return ctx.Status(code).JSON(response)
//}
