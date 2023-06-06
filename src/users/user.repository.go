package users

import (
	"fmt"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type RepositoryError struct {
	status  int
	message string
}

func (r *RepositoryError) Error() string {
	return fmt.Sprintf("%d-%s", r.status, r.message)
}

type IUserRepository interface {
	Create(user User)
	Find() []User
	FindOne(id int) (User, error)
	UpdateOne(id int, user User) User
	DeleteOne(id int) (User, error)
}

type UserRepository struct {
	DB gorm.DB
}

func (repository *UserRepository) Create(user User) error {
	return repository.DB.Create(user)
}

func (repository *UserRepository) Find() []User {
	var users []User
	repository.DB.Find(&users)
	return users
}

func (repository *UserRepository) FindOne(id int) (User, error) {
	var user User

	result := repository.DB.Find(&user, id)
	if result.RowsAffected == 0 {
		return user, &RepositoryError{404, "not found"}
	}
	return user, nil
}

func (repository *UserRepository) UpdateOne(id int, user User) {
	repository.DB.Where("id = ?", id).Updates(user)
}

func (repository *UserRepository) DeleteOne(id int) (User, error) {
	var user User
	result := repository.DB.Delete(&user, id)
	if result.RowsAffected == 0 {
		return user, &RepositoryError{404, "not found"}
	}
	return user, nil
}

func NewUserRepository(db gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

var SetRepository = wire.NewSet(NewUserRepository)
