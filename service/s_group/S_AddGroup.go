package s_group

import (
	"sitax/model/m_group"
	"sitax/repository/r_group"

	"github.com/gin-gonic/gin"
)

type AddGroupService struct {
	addGroupRepo r_group.AddGroupRepository
}

func NewAddGroupService(addGroupRepo r_group.AddGroupRepository) *AddGroupService {
	return &AddGroupService{addGroupRepo}
}

func (s *AddGroupService) AddGroup(ctx *gin.Context, group m_group.Group, groupID string) (*m_group.Group, error) {
	addGroup, err := s.addGroupRepo.AddGroup(&group, groupID)
	if err != nil {
		return nil, err
	}

	return addGroup, nil
}
