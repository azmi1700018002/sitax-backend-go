package r_group

import (
	"sitax/config/db"
	"sitax/model/m_group"
)

type DeleteGroupRepository interface {
	DeleteGroupByID(id string) error
}

type groupDeleteRepository struct{}

func NewDeleteGroupRepository() DeleteGroupRepository {
	return &groupDeleteRepository{}
}

func (*groupDeleteRepository) DeleteGroupByID(id string) (err error) {
	// Menghapus data group dari tabel Group
	if err = db.Server().Unscoped().Where("group_id = ?", id).Delete(&m_group.Group{}).Error; err != nil {
		return err
	}

	return nil
}
