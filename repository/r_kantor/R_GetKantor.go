package r_kantor

import (
	"sitax/config/db"
	"sitax/model/m_kantor"
)

type GetKantorRepository interface {
	GetAllKantor() ([]m_kantor.Kantor, error)
	GetKantorByID(KdKantor string) (*m_kantor.Kantor, error)
}

type getKantorRepository struct{}

func NewGetKantorRepository() GetKantorRepository {
	return &getKantorRepository{}
}

func (r *getKantorRepository) GetAllKantor() ([]m_kantor.Kantor, error) {
	var kantors []m_kantor.Kantor
	result := db.Server().Preload("KdKantorfk").Find(&kantors)
	// result := db.Server().Find(&kantors)
	if result.Error != nil {
		return nil, result.Error
	}
	return kantors, nil
}

func (r *getKantorRepository) GetKantorByID(KdKantor string) (*m_kantor.Kantor, error) {
	var kantor m_kantor.Kantor
	result := db.Server().Where("kd_kantor = ?", KdKantor).First(&kantor)
	if result.Error != nil {
		return nil, result.Error
	}
	return &kantor, nil
}
