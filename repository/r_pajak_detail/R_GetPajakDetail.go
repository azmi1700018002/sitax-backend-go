package r_pajak_detail

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type GetPajakDetailRepository interface {
	GetAllPajakDetail() ([]m_pajak.PajakDetail, error)
	GetPajakDetailByID(PajakDetailID int) (*m_pajak.PajakDetail, error)
}

type getPajakDetailRepository struct{}

func NewGetPajakDetailRepository() GetPajakDetailRepository {
	return &getPajakDetailRepository{}
}

func (r *getPajakDetailRepository) GetAllPajakDetail() ([]m_pajak.PajakDetail, error) {
	var pajak_details []m_pajak.PajakDetail
	// result := db.Server().Preload("PajakDetailIDfk").Find(&pajak_details)
	result := db.Server().Find(&pajak_details)
	if result.Error != nil {
		return nil, result.Error
	}
	return pajak_details, nil
}

func (r *getPajakDetailRepository) GetPajakDetailByID(PajakDetailID int) (*m_pajak.PajakDetail, error) {
	var pajak m_pajak.PajakDetail
	result := db.Server().Where("pajak_detail_id = ?", PajakDetailID).First(&pajak)
	if result.Error != nil {
		return nil, result.Error
	}
	return &pajak, nil
}
