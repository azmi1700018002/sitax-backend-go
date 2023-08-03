package s_kewenangan

import "sitax/repository/r_kewenangan"

type DeleteKewenanganService interface {
	DeleteKewenanganByID(GroupID string, MenuID string) error
}

type deleteKewenanganService struct {
	deleteKewenanganRepo r_kewenangan.DeleteKewenanganRepository
}

func NewDeleteKewenanganService(deleteKewenanganRepo r_kewenangan.DeleteKewenanganRepository) DeleteKewenanganService {
	return &deleteKewenanganService{deleteKewenanganRepo}
}

func (s *deleteKewenanganService) DeleteKewenanganByID(GroupID string, MenuID string) error {
	err := s.deleteKewenanganRepo.DeleteKewenanganByID(GroupID, MenuID)
	if err != nil {
		return err
	}
	return nil
}
