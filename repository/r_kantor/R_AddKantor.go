package r_kantor

import (
	"sitax/config/db"
	"sitax/model/m_kantor"
)

type AddKantorRepository interface {
	AddKantor(kantor *m_kantor.Kantor, kantorID string) (*m_kantor.Kantor, error)
}

type addKantorRepository struct{}

func NewAddKantorRepository() AddKantorRepository {
	return &addKantorRepository{}
}

func (r *addKantorRepository) AddKantor(kantor *m_kantor.Kantor, kantorID string) (*m_kantor.Kantor, error) {
	err := db.Server().Create(&kantor).Error
	if err != nil {
		return nil, err
	}
	return kantor, nil
}
