package r_pajak_detail

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type DeletePajakDetailRepository interface {
	DeletePajakDetailByID(PajakDetailID int) error
}

type pajakDeleteRepository struct{}

func NewDeletePajakDetailRepository() DeletePajakDetailRepository {
	return &pajakDeleteRepository{}
}

func (*pajakDeleteRepository) DeletePajakDetailByID(PajakDetailID int) (err error) {
	// Menghapus data pajak dari tabel PajakDetail
	if err = db.Server().Unscoped().Where("pajak_detail_id = ?", PajakDetailID).Delete(&m_pajak.PajakDetail{}).Error; err != nil {
		return err
	}

	return nil
}
