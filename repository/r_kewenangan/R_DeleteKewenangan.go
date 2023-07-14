package r_kewenangan

import (
	"sitax/config/db"
	"sitax/model/m_kewenangan"
)

type DeleteKewenanganRepository interface {
	DeleteKewenanganByID(GroupID string) error
}

type kewenanganDeleteRepository struct{}

func NewDeleteKewenanganRepository() DeleteKewenanganRepository {
	return &kewenanganDeleteRepository{}
}

func (*kewenanganDeleteRepository) DeleteKewenanganByID(GroupID string) (err error) {
	// Menghapus data kewenangan dari tabel Kewenangan
	if err = db.Server().Unscoped().Where("group_id = ?", GroupID).Delete(&m_kewenangan.Kewenangan{}).Error; err != nil {
		return err
	}

	return nil
}
