package r_user

import (
	"sitax/config/db"
	"sitax/config/helper"
	"sitax/model/m_user"
)

type RegisterUserRepository interface {
	RegisterUser(user *m_user.User) error
}

type registerUserRepository struct{}

func NewUserRepository() RegisterUserRepository {
	return &registerUserRepository{}
}

func (r *registerUserRepository) RegisterUser(user *m_user.User) error {

	// Hash password before save
	hashedPassword, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// Menyimpan user ke dalam database
	if err := db.Server().Create(user).Error; err != nil {
		return err
	}

	return nil
}
