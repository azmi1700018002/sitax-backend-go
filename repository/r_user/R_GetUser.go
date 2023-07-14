package r_user

import (
	"sitax/config/db"
	"sitax/model/m_user"
)

type GetUserRepository interface {
	GetAllUser() ([]m_user.User, error)
	GetUserByID(username string) (*m_user.User, error)
}

type getUserRepository struct{}

func NewGetUserRepository() GetUserRepository {
	return &getUserRepository{}
}

func (r *getUserRepository) GetAllUser() ([]m_user.User, error) {
	var users []m_user.User
	result := db.Server().Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *getUserRepository) GetUserByID(username string) (*m_user.User, error) {
	var user m_user.User
	result := db.Server().Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
