package s_group

import "sitax/repository/r_group"

type DeleteGroupService interface {
	DeleteGroupByID(id string) error
}

type deleteGroupService struct {
	deleteGroupRepo r_group.DeleteGroupRepository
}

func NewDeleteGroupService(deleteGroupRepo r_group.DeleteGroupRepository) DeleteGroupService {
	return &deleteGroupService{deleteGroupRepo}
}

func (s *deleteGroupService) DeleteGroupByID(id string) error {
	err := s.deleteGroupRepo.DeleteGroupByID(id)
	if err != nil {
		return err
	}
	return nil
}
