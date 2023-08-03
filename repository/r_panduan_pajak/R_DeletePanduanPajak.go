package r_panduan_pajak

import (
	"sitax/config/db"
	"sitax/model/m_panduan_pajak"
)

type DeletePanduanPajakRepository interface {
	DeletePanduanPajakByID(PanduanPajakID int) error
}

type panduan_pajakDeleteRepository struct{}

func NewDeletePanduanPajakRepository() DeletePanduanPajakRepository {
	return &panduan_pajakDeleteRepository{}
}

func (*panduan_pajakDeleteRepository) DeletePanduanPajakByID(PanduanPajakID int) (err error) {
	// Menghapus data panduan_pajak dari tabel PanduanPajak
	if err = db.Server().Unscoped().Where("panduan_pajak_id = ?", PanduanPajakID).Delete(&m_panduan_pajak.PanduanPajak{}).Error; err != nil {
		return err
	}

	return nil
}
