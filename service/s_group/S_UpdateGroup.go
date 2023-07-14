package s_group

import (
	"sitax/model/m_group"
	"sitax/repository/r_group"

	"github.com/gin-gonic/gin"
)

type UpdateGroupService struct {
	updateGroupRepo r_group.UpdateGroupRepository
}

func NewUpdateGroupService(updateGroupRepo r_group.UpdateGroupRepository) *UpdateGroupService {
	return &UpdateGroupService{updateGroupRepo}
}

func (s *UpdateGroupService) UpdateGroup(ctx *gin.Context, group m_group.Group) (*m_group.Group, error) {
	// Mendapatkan ID grup dari parameter rute
	idGroupStr := ctx.Param("group_id")

	// Set ID grup ke dalam struct grup
	group.GroupID = idGroupStr

	// Memanggil repository untuk memperbarui grup
	updatedGroup, err := s.updateGroupRepo.UpdateGroup(&group)
	if err != nil {
		return nil, err
	}

	return updatedGroup, nil
}
