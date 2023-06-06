package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"strconv"
)

type IUserService interface {
	CreateOne(User) User
	FindMany() []User
	FindOne(id int) User
	UpdateOne(id int, user User) User
	DeleteOne(id int)
}

type UserService struct {
	Repository IUserRepository
}

func (service *UserService) CreateOne(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	service.Repository.Create(*user)

	return c.Status(201).JSON(user)
}

func (service *UserService) FindMany(c *fiber.Ctx) error {
	return c.Status(200).JSON(service.Repository.Find())
}

func (service *UserService) FindOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	result, err := service.Repository.FindOne(id)
	if err != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(result)
}

func (service *UserService) UpdateOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	service.Repository.UpdateOne(id, *user)
	return c.Status(200).JSON(user)
}

func (service *UserService) DeleteOne(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user, err := service.Repository.DeleteOne(id)
	if err != nil {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(user)
}

func NewUserService(userRepository IUserRepository) *UserService {
	return &UserService{
		Repository: userRepository,
	}
}

var SetService = wire.NewSet(NewUserService)
