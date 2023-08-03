package r_panduan_pajak

import (
	"sitax/config/db"
	"sitax/model/m_panduan_pajak"
)

type GetPanduanPajakRepository interface {
	GetAllPanduanPajak() ([]m_panduan_pajak.PanduanPajak, error)
	GetPanduanPajakByID(PanduanPajakID int) (*m_panduan_pajak.PanduanPajak, error)
}

type getPanduanPajakRepository struct{}

func NewGetPanduanPajakRepository() GetPanduanPajakRepository {
	return &getPanduanPajakRepository{}
}

func (r *getPanduanPajakRepository) GetAllPanduanPajak() ([]m_panduan_pajak.PanduanPajak, error) {
	var panduan_pajaks []m_panduan_pajak.PanduanPajak
	// result := db.Server().Preload("PanduanPajakIDfk").Find(&panduan_pajaks)
	result := db.Server().Find(&panduan_pajaks)
	if result.Error != nil {
		return nil, result.Error
	}
	return panduan_pajaks, nil
}

func (r *getPanduanPajakRepository) GetPanduanPajakByID(PanduanPajakID int) (*m_panduan_pajak.PanduanPajak, error) {
	var panduan_pajak m_panduan_pajak.PanduanPajak
	result := db.Server().Where("id_panduan_pajak = ?", PanduanPajakID).First(&panduan_pajak)
	if result.Error != nil {
		return nil, result.Error
	}
	return &panduan_pajak, nil
}
