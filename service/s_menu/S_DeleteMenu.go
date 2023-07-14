package s_menu

import "sitax/repository/r_menu"

type DeleteMenuService interface {
	DeleteMenuByID(id string) error
}

type deleteMenuService struct {
	deleteMenuRepo r_menu.DeleteMenuRepository
}

func NewDeleteMenuService(deleteMenuRepo r_menu.DeleteMenuRepository) DeleteMenuService {
	return &deleteMenuService{deleteMenuRepo}
}

func (s *deleteMenuService) DeleteMenuByID(id string) error {
	err := s.deleteMenuRepo.DeleteMenuByID(id)
	if err != nil {
		return err
	}
	return nil
}
