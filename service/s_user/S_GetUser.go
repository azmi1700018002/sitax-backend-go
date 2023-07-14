package s_user

import (
	"sitax/model/m_user"
	r_user "sitax/repository/r_user"
)

type GetUserService interface {
	GetAllUser() ([]m_user.User, error)
	GetUserByID(username string) (*m_user.User, error)
}

type getUserService struct {
	getUserRepository r_user.GetUserRepository
}

func NewGetUserService(getUserRepository r_user.GetUserRepository) GetUserService {
	return &getUserService{
		getUserRepository: getUserRepository,
	}
}

func (s *getUserService) GetAllUser() ([]m_user.User, error) {
	return s.getUserRepository.GetAllUser()
}

func (s *getUserService) GetUserByID(username string) (*m_user.User, error) {
	return s.getUserRepository.GetUserByID(username)
}
