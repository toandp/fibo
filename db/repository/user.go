package repository

import (
	"gorm.io/gorm"
	"team1.asia/fibo/db/entity"
	"team1.asia/fibo/log"
)

type UserRepository struct {
	ORM *gorm.DB
}

type UserRepositoryInterface interface {
	GetByUsername(username string) *entity.User
	Create(user *entity.User) *entity.User
}

// Create a new user repository instance.
func New(orm *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		ORM: orm,
	}
}

// Get the user by username.
func (repo *UserRepository) GetByUsername(username string) *entity.User {
	var user entity.User

	if err := repo.ORM.Where(&entity.User{Username: username}).First(&user).Error; err != nil {
		log.Error(err.Error())
		panic(err)
	}

	return &user
}

// Create a new user.
func (repo *UserRepository) Create(user *entity.User) *entity.User {
	repo.ORM.Create(&user)

	return user
}
