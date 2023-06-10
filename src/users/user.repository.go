package users

import (
	"fmt"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type RepositoryError struct {
	status  int
	message string
}

func (e *RepositoryError) Error() string {
	return fmt.Sprintf("Status: %d, Message: %s", e.status, e.message)
}

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
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(&RepositoryError{503, "db error"}, "not affected")
	}
	return &user, nil
}

func (repository *UserRepository) Find() ([]User, error) {
	var users []User
	if err := repository.DB.Where("deleted_at != ?", "null").Find(&users).Error; err != nil {
		return nil, errors.Wrap(&RepositoryError{503, "aa"}, "getUsers failure")
	}

	return users, nil
}

func (repository *UserRepository) FindOne(id int) (*User, error) {
	var user User
	result := repository.DB.Find(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(&RepositoryError{404, "not found"}, "not affected")
	}

	return &user, nil
}

func (repository *UserRepository) UpdateOne(id int, user User) (*User, error) {
	result := repository.DB.Where("id = ?", id).Updates(&user)
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(&RepositoryError{404, "not found"}, "not affected")
	}
	return &user, nil
}

func (repository *UserRepository) DeleteOne(id int) (*User, error) {
	var user User
	result := repository.DB.Delete(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.Wrap(&RepositoryError{404, "not found"}, "not affected")
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
