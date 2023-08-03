package r_panduan_pajak

import (
	"sitax/config/db"
	"sitax/model/m_panduan_pajak"
)

type AddPanduanPajakRepository interface {
	AddPanduanPajak(panduan_pajak *m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error)
}

type addPanduanPajakRepository struct{}

func NewAddPanduanPajakRepository() AddPanduanPajakRepository {
	return &addPanduanPajakRepository{}
}

func (r *addPanduanPajakRepository) AddPanduanPajak(panduan_pajak *m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error) {
	err := db.Server().Create(&panduan_pajak).Error
	if err != nil {
		return nil, err
	}
	return panduan_pajak, nil
}
