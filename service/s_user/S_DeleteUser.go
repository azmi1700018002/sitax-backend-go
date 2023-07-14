package s_user

import (
	r_user "sitax/repository/r_user"
)

type DeleteUserService interface {
	DeleteUserByID(username string) error
}

type deleteUserService struct {
	deleteUserRepo r_user.DeleteUserRepository
}

func NewDeleteUserService(deleteUserRepo r_user.DeleteUserRepository) DeleteUserService {
	return &deleteUserService{deleteUserRepo}
}

func (s *deleteUserService) DeleteUserByID(username string) error {
	err := s.deleteUserRepo.DeleteUserByID(username)
	if err != nil {
		return err
	}
	return nil
}
