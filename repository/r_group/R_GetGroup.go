package r_group

import (
	"sitax/config/db"
	"sitax/model/m_group"
)

type GetGroupRepository interface {
	GetAllGroup() ([]m_group.Group, error)
	GetGroupByID(groupID int) (*m_group.Group, error)
}

type getGroupRepository struct{}

func NewGetGroupRepository() GetGroupRepository {
	return &getGroupRepository{}
}

func (r *getGroupRepository) GetAllGroup() ([]m_group.Group, error) {
	var groups []m_group.Group
	result := db.Server().Preload("GroupIDfk").Preload("GroupIDfk2").Find(&groups)
	// result := db.Server().Find(&groups)
	if result.Error != nil {
		return nil, result.Error
	}
	return groups, nil
}

func (r *getGroupRepository) GetGroupByID(groupID int) (*m_group.Group, error) {
	var group m_group.Group
	result := db.Server().Where("group_id = ?", groupID).First(&group)
	if result.Error != nil {
		return nil, result.Error
	}
	return &group, nil
}
