package r_pajak_detail

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type AddPajakDetailRepository interface {
	AddPajakDetail(pajak_detail *m_pajak.PajakDetail) (*m_pajak.PajakDetail, error)
}

type addPajakDetailRepository struct{}

func NewAddPajakDetailRepository() AddPajakDetailRepository {
	return &addPajakDetailRepository{}
}

func (r *addPajakDetailRepository) AddPajakDetail(pajak_detail *m_pajak.PajakDetail) (*m_pajak.PajakDetail, error) {
	err := db.Server().Create(&pajak_detail).Error
	if err != nil {
		return nil, err
	}
	return pajak_detail, nil
}
