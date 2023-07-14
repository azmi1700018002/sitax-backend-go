package r_pajak_detail

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type UpdatePajakDetailRepository interface {
	UpdatePajakDetail(pajak_detail *m_pajak.PajakDetail) (*m_pajak.PajakDetail, error)
}

type updatePajakDetailRepository struct{}

func NewUpdatePajakDetailRepository() UpdatePajakDetailRepository {
	return &updatePajakDetailRepository{}
}

func (r *updatePajakDetailRepository) UpdatePajakDetail(pajak_detail *m_pajak.PajakDetail) (*m_pajak.PajakDetail, error) {
	// Check if the user exists
	var existingPajakDetail m_pajak.PajakDetail
	if err := db.Server().Where("pajak_detail_id = ?", pajak_detail.PajakDetailID).First(&existingPajakDetail).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_pajak.PajakDetail{}).
		Where("pajak_detail_id = ?", pajak_detail.PajakDetailID).
		Updates(map[string]interface{}{
			"pajak_id":   pajak_detail.PajakID,
			"ppn":        pajak_detail.Ppn,
			"pasal23":    pajak_detail.Pasal23,
			"pph_final":  pajak_detail.PphFinal,
			"pajak_lain": pajak_detail.PajakLain,
			"keterangan": pajak_detail.Keterangan,
		}).Error; err != nil {
		return nil, err
	}
	return pajak_detail, nil
}
