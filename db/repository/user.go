package repository

import (
	"gorm.io/gorm"
	"team1.asia/fibo/db/entity"
)

type UserRepository struct {
	ORM *gorm.DB
}

type UserRepositoryInterface interface {
	GetByUsername(username string) *entity.User
	Create(user *entity.User) *entity.User
}

func New(orm *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		ORM: orm,
	}
}

// Gets the user by username.
// @param  username string
// @return *entity.User
func (repo *UserRepository) GetByUsername(username string) *entity.User {
	var user entity.User

	if err := repo.ORM.Where(&entity.User{Username: username}).First(&user).Error; err != nil {
		panic(err)
	}

	return &user
}

// Create the user.
// @param  user *entity.User
// @return *entity.User
func (repo *UserRepository) Create(user *entity.User) *entity.User {
	repo.ORM.Create(&user)

	return user
}
