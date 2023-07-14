package r_group

import (
	"sitax/config/db"
	"sitax/model/m_group"
)

type AddGroupRepository interface {
	AddGroup(group *m_group.Group, groupID string) (*m_group.Group, error)
}

type addGroupRepository struct{}

func NewAddGroupRepository() AddGroupRepository {
	return &addGroupRepository{}
}

func (r *addGroupRepository) AddGroup(group *m_group.Group, groupID string) (*m_group.Group, error) {
	err := db.Server().Create(&group).Error
	if err != nil {
		return nil, err
	}
	return group, nil
}
