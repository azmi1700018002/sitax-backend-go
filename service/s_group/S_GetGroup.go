package s_group

import (
	"sitax/model/m_group"
	"sitax/repository/r_group"
)

type GetGroupService interface {
	GetAllGroup() ([]m_group.Group, error)
	GetGroupByID(groupID int) (*m_group.Group, error)
}

type getGroupService struct {
	getGroupRepository r_group.GetGroupRepository
}

func NewGetGroupService(getGroupRepository r_group.GetGroupRepository) GetGroupService {
	return &getGroupService{
		getGroupRepository: getGroupRepository,
	}
}

func (s *getGroupService) GetAllGroup() ([]m_group.Group, error) {
	return s.getGroupRepository.GetAllGroup()
}

func (s *getGroupService) GetGroupByID(groupID int) (*m_group.Group, error) {
	return s.getGroupRepository.GetGroupByID(groupID)
}
