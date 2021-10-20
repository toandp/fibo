package repository

import (
	"gorm.io/gorm"
	"team1.asia/fibo/db"
	"team1.asia/fibo/db/entity"
)

type UserRepository struct {
	ORM *gorm.DB
}

type UserRepositoryInterface interface {
	GetByUsername(username string) (*entity.User, error)
	Create(user *entity.User) (*entity.User, error)
}

func New(orm *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		ORM: orm,
	}
}

func (repo *UserRepository) GetByUsername(username string) (*entity.User, error) {
	var user entity.User

	if err := repo.ORM.Where(&entity.User{Username: username}).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user *entity.User) (*entity.User, error) {
	db.ORM.Create(&user)

	return user, nil
}
