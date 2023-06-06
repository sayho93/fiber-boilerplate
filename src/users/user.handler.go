package users

import (
	"fiber/src/common/database"
	"github.com/gofiber/fiber/v2"
)

func CreateOne(c *fiber.Ctx) error {
	user := new(User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Connection.Create(&user)

	return c.Status(201).JSON(user)
}

func FindMany(c *fiber.Ctx) error {
	var users []User

	database.Connection.Find(&users)
	return c.Status(200).JSON(users)
}

func FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User

	result := database.Connection.Find(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(&user)
}

func UpdateOne(c *fiber.Ctx) error {
	user := new(User)
	id := c.Params("id")

	if err := c.BodyParser(user); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	database.Connection.Where("id = ?", id).Updates(&user)
	return c.Status(200).JSON(user)
}

func DeleteOne(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User

	result := database.Connection.Delete(&user, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}
