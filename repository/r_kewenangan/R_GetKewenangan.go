package r_kewenangan

import (
	"sitax/config/db"
	"sitax/model/m_kewenangan"
)

type GetKewenanganRepository interface {
	GetAllKewenangan() ([]m_kewenangan.Kewenangan, error)
	GetKewenanganByID(GroupID string) (*m_kewenangan.Kewenangan, error)
}

type getKewenanganRepository struct{}

func NewGetKewenanganRepository() GetKewenanganRepository {
	return &getKewenanganRepository{}
}

func (r *getKewenanganRepository) GetAllKewenangan() ([]m_kewenangan.Kewenangan, error) {
	var kewenangans []m_kewenangan.Kewenangan
	result := db.Server().Preload("GroupIDfk").Find(&kewenangans)
	// result := db.Server().Find(&kewenangans)
	if result.Error != nil {
		return nil, result.Error
	}
	return kewenangans, nil
}

func (r *getKewenanganRepository) GetKewenanganByID(GroupID string) (*m_kewenangan.Kewenangan, error) {
	var kewenangan m_kewenangan.Kewenangan
	result := db.Server().Where("group_id = ?", GroupID).First(&kewenangan)
	if result.Error != nil {
		return nil, result.Error
	}
	return &kewenangan, nil
}
