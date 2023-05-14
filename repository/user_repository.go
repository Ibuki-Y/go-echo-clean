package repository

import (
	"github.com/Ibuki-Y/go-echo-clean/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (u *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := u.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *userRepository) CreateUser(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
