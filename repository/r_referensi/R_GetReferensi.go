package r_referensi

import (
	"sitax/config/db"
	"sitax/model/m_referensi"
)

type GetReferensiRepository interface {
	GetAllReferensi(grpID, ref, ket string) ([]m_referensi.Referensi, error)
	GetReferensiByID(referensiID int) (*m_referensi.Referensi, error)
}

type getReferensiRepository struct{}

func NewGetReferensiRepository() GetReferensiRepository {
	return &getReferensiRepository{}
}

func (r *getReferensiRepository) GetAllReferensi(grpID, ref, ket string) ([]m_referensi.Referensi, error) {
	var referensis []m_referensi.Referensi
	db := db.Server()

	// Memulai query dengan menambahkan WHERE jika parameter query tidak kosong
	if grpID != "" {
		db = db.Where("grp_id = ?", grpID)
	}
	if ref != "" {
		db = db.Where("ref = ?", ref)
	}
	if ket != "" {
		db = db.Where("ket = ?", ket)
	}

	result := db.Find(&referensis)
	if result.Error != nil {
		return nil, result.Error
	}
	return referensis, nil
}

func (r *getReferensiRepository) GetReferensiByID(referensiID int) (*m_referensi.Referensi, error) {
	var referensi m_referensi.Referensi
	result := db.Server().Where("referensi_id = ?", referensiID).First(&referensi)
	if result.Error != nil {
		return nil, result.Error
	}
	return &referensi, nil
}
