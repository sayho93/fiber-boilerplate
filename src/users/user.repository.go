package users

import (
	"fiber/src/common/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user User) (*User, error)
	Find() ([]User, error)
	FindOne(id int) (*User, error)
	UpdateOne(id int, user User) (*User, error)
	DeleteOne(id int) (*User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (repository *UserRepository) Create(user User) (*User, error) {
	result := repository.DB.Create(&user)
	if result.Error != nil {
		return nil, errors.New(fiber.StatusServiceUnavailable, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fiber.StatusNotFound, "not affected")
	}

	return &user, nil
}

func (repository *UserRepository) Find() ([]User, error) {
	var users []User
	if err := repository.DB.Where("deletedAt != ?", "null").Find(&users).Error; err != nil {
		return nil, errors.New(fiber.StatusServiceUnavailable, err.Error())
	}

	return users, nil
}

func (repository *UserRepository) FindOne(id int) (*User, error) {
	var user User
	result := repository.DB.Find(&user, id)
	if result.Error != nil {
		return nil, errors.New(fiber.StatusServiceUnavailable, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fiber.StatusNotFound, "Not affected")
	}

	return &user, nil
}

func (repository *UserRepository) UpdateOne(id int, user User) (*User, error) {
	result := repository.DB.Where("id = ?", id).Updates(&user)
	if result.Error != nil {
		return nil, errors.New(fiber.StatusServiceUnavailable, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fiber.StatusNotFound, "not affected")
	}
	return &user, nil
}

func (repository *UserRepository) DeleteOne(id int) (*User, error) {
	var user User
	result := repository.DB.Delete(&user, id)
	if result.Error != nil {
		return nil, errors.New(fiber.StatusServiceUnavailable, result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return nil, errors.New(fiber.StatusNotFound, "not affected")
	}
	return &user, nil
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

var SetRepository = wire.NewSet(
	NewUserRepository,
	wire.Bind(new(IUserRepository), new(*UserRepository)),
)
