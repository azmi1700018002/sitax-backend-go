package r_panduan_pajak

import (
	"sitax/config/db"
	"sitax/model/m_panduan_pajak"
)

type UpdatePanduanPajakRepository interface {
	UpdatePanduanPajak(panduan_pajak *m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error)
}

type updatePanduanPajakRepository struct{}

func NewUpdatePanduanPajakRepository() UpdatePanduanPajakRepository {
	return &updatePanduanPajakRepository{}
}

func (r *updatePanduanPajakRepository) UpdatePanduanPajak(panduan_pajak *m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error) {
	// Check if the user exists
	var existingPanduanPajak m_panduan_pajak.PanduanPajak
	if err := db.Server().Where("panduan_pajak_id = ?", panduan_pajak.PanduanPajakID).First(&existingPanduanPajak).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_panduan_pajak.PanduanPajak{}).
		Where("panduan_pajak_id = ?", panduan_pajak.PanduanPajakID).
		Updates(map[string]interface{}{
			"nama_panduan_pajak":   panduan_pajak.NamaPanduanPajak,
			"parent_panduan_pajak": panduan_pajak.ParentPanduanPajak,
			"sts_panduan_pajak":    panduan_pajak.StsPanduanPajak,
			"sts_parent":           panduan_pajak.StsParent,
			"file_id":              panduan_pajak.FileID,
		}).Error; err != nil {
		return nil, err
	}
	return panduan_pajak, nil
}
