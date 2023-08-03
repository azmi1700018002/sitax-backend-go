package r_kewenangan

import (
	"sitax/config/db"
	"sitax/model/m_kewenangan"
)

type DeleteKewenanganRepository interface {
	DeleteKewenanganByID(GroupID string, MenuID string) error
}

type kewenanganDeleteRepository struct{}

func NewDeleteKewenanganRepository() DeleteKewenanganRepository {
	return &kewenanganDeleteRepository{}
}

func (*kewenanganDeleteRepository) DeleteKewenanganByID(GroupID string, MenuID string) (err error) {
	// Menghapus data kewenangan dari tabel Kewenangan berdasarkan group_id dan menu_id
	if err = db.Server().Unscoped().Where("group_id = ? AND menu_id = ?", GroupID, MenuID).Delete(&m_kewenangan.Kewenangan{}).Error; err != nil {
		return err
	}

	return nil
}
