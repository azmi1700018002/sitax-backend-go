package r_kewenangan

import (
	"sitax/config/db"
	"sitax/model/m_kewenangan"
)

type AddKewenanganRepository interface {
	AddKewenangan(kewenangan *m_kewenangan.Kewenangan, GroupID string) (*m_kewenangan.Kewenangan, error)
}

type addKewenanganRepository struct{}

func NewAddKewenanganRepository() AddKewenanganRepository {
	return &addKewenanganRepository{}
}

func (r *addKewenanganRepository) AddKewenangan(kewenangan *m_kewenangan.Kewenangan, GroupID string) (*m_kewenangan.Kewenangan, error) {
	err := db.Server().Create(&kewenangan).Error
	if err != nil {
		return nil, err
	}
	return kewenangan, nil
}
