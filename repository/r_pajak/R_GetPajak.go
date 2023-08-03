package r_pajak

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type GetPajakRepository interface {
	GetAllPajak() ([]m_pajak.Pajak, error)
	GetPajakByID(PajakID int) (*m_pajak.Pajak, error)
}

type getPajakRepository struct{}

func NewGetPajakRepository() GetPajakRepository {
	return &getPajakRepository{}
}

func (r *getPajakRepository) GetAllPajak() ([]m_pajak.Pajak, error) {
	var pajaks []m_pajak.Pajak
	result := db.Server().Preload("PajakIDfk").Find(&pajaks)
	// result := db.Server().Find(&pajaks)
	if result.Error != nil {
		return nil, result.Error
	}
	return pajaks, nil
}

func (r *getPajakRepository) GetPajakByID(PajakID int) (*m_pajak.Pajak, error) {
	var pajak m_pajak.Pajak
	result := db.Server().Where("pajak_id = ?", PajakID).Preload("PajakIDfk").First(&pajak)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pajak, nil
}
