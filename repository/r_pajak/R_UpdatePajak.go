package r_pajak

import (
	"sitax/config/db"
	"sitax/model/m_pajak"
)

type UpdatePajakRepository interface {
	UpdatePajak(pajak *m_pajak.Pajak) (*m_pajak.Pajak, error)
}

type updatePajakRepository struct{}

func NewUpdatePajakRepository() UpdatePajakRepository {
	return &updatePajakRepository{}
}

func (r *updatePajakRepository) UpdatePajak(pajak *m_pajak.Pajak) (*m_pajak.Pajak, error) {
	// Check if the user exists
	var existingPajak m_pajak.Pajak
	if err := db.Server().Where("pajak_id = ?", pajak.PajakID).First(&existingPajak).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_pajak.Pajak{}).
		Where("pajak_id = ?", pajak.PajakID).
		Updates(map[string]interface{}{
			"nama_pajak":   pajak.NamaPajak,
			"parent_pajak": pajak.ParentPajak,
			"sts_pajak":    pajak.StsPajak,
			"ket_pajak":    pajak.KetPajak,
			"sts_parent":   pajak.StsParent,
			"file_id":      pajak.FileID,
		}).Error; err != nil {
		return nil, err
	}
	return pajak, nil
}
