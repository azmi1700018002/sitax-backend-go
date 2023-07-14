package r_pajak

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type AddPajakRepository interface {
	AddPajak(pajak *m_pajak.Pajak) (*m_pajak.Pajak, error)
}

type addPajakRepository struct{}

func NewAddPajakRepository() AddPajakRepository {
	return &addPajakRepository{}
}

func (r *addPajakRepository) AddPajak(pajak *m_pajak.Pajak) (*m_pajak.Pajak, error) {
	err := db.Server().Create(&pajak).Error
	if err != nil {
		return nil, err
	}
	return pajak, nil
}
