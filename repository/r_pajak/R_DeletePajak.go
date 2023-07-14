package r_pajak

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type DeletePajakRepository interface {
	DeletePajakByID(PajakID int) error
}

type pajakDeleteRepository struct{}

func NewDeletePajakRepository() DeletePajakRepository {
	return &pajakDeleteRepository{}
}

func (*pajakDeleteRepository) DeletePajakByID(PajakID int) (err error) {
	// Menghapus data pajak dari tabel Pajak
	if err = db.Server().Unscoped().Where("pajak_id = ?", PajakID).Delete(&m_pajak.Pajak{}).Error; err != nil {
		return err
	}

	return nil
}
