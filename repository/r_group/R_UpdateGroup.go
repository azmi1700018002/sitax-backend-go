package r_group

import (
	"sitax/config/db"
	"sitax/model/m_group"
)

type UpdateGroupRepository interface {
	UpdateGroup(group *m_group.Group) (*m_group.Group, error)
}

type updateGroupRepository struct{}

func NewUpdateGroupRepository() UpdateGroupRepository {
	return &updateGroupRepository{}
}

func (r *updateGroupRepository) UpdateGroup(group *m_group.Group) (*m_group.Group, error) {
	// Check if the user exists
	var existingGroup m_group.Group
	if err := db.Server().Where("group_id = ?", group.GroupID).First(&existingGroup).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_group.Group{}).
		Where("group_id = ?", group.GroupID).
		Updates(map[string]interface{}{
			"group_nama":      group.GroupNama,
			"group_deskripsi": group.GroupDeskripsi,
		}).Error; err != nil {
		return nil, err
	}
	return group, nil
}
