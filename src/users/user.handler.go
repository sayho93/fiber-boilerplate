package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
	"strconv"
)

type UserHandler interface {
	CreateOne(c *fiber.Ctx) error
	FindMany(c *fiber.Ctx) error
	FindOne(c *fiber.Ctx) error
	UpdateOne(c *fiber.Ctx) error
	DeleteOne(c *fiber.Ctx) error
}

type userHandler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &userHandler{service: service}
}

var SetHandler = wire.NewSet(NewUserHandler)

func (handler userHandler) CreateOne(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	result, err := handler.service.CreateOne(user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (handler userHandler) FindMany(c *fiber.Ctx) error {
	result, err := handler.service.FindMany()
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler userHandler) FindOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	result, err := handler.service.FindOne(id)
	if err != nil {
		return err
	}
	logrus.Info("test")
	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler userHandler) UpdateOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	result, err := handler.service.UpdateOne(id, user)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(result)
}

func (handler userHandler) DeleteOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := handler.service.DeleteOne(id)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(user)
}
