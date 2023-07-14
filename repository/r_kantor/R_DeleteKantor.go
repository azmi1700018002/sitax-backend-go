package r_kantor

import (
	"sitax/config/db"
	"sitax/model/m_kantor"
	"sitax/model/m_user"
)

type DeleteKantorRepository interface {
	DeleteKantorByID(KdKantor string) error
}

type kantorDeleteRepository struct{}

func NewDeleteKantorRepository() DeleteKantorRepository {
	return &kantorDeleteRepository{}
}

func (*kantorDeleteRepository) DeleteKantorByID(KdKantor string) (err error) {
	// Mengubah nilai KdKantor pada entitas pengguna yang memiliki KdKantor yang sesuai menjadi null
	if err = db.Server().Model(&m_user.User{}).Where("kd_kantor = ?", KdKantor).Update("kd_kantor", nil).Error; err != nil {
		return err
	}

	// Menghapus data kantor dari tabel Kantor
	if err = db.Server().Unscoped().Where("kd_kantor = ?", KdKantor).Delete(&m_kantor.Kantor{}).Error; err != nil {
		return err
	}

	return nil
}
