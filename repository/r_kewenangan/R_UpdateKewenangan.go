package r_kewenangan

import (
	"sitax/config/db"
	"sitax/model/m_kewenangan"
)

type UpdateKewenanganRepository interface {
	UpdateKewenangan(kewenangan *m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error)
}

type updateKewenanganRepository struct{}

func NewUpdateKewenanganRepository() UpdateKewenanganRepository {
	return &updateKewenanganRepository{}
}

func (r *updateKewenanganRepository) UpdateKewenangan(kewenangan *m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error) {
	// Check if the user exists
	var existingKewenangan m_kewenangan.Kewenangan
	if err := db.Server().Where("group_id = ?", kewenangan.GroupID).First(&existingKewenangan).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_kewenangan.Kewenangan{}).
		Where("group_id = ?", kewenangan.GroupID).
		Updates(map[string]interface{}{
			"menu_id":    kewenangan.MenuID,
			"is_created": kewenangan.IsCreated,
			"is_updated": kewenangan.IsUpdated,
			"is_deleted": kewenangan.IsDeleted,
		}).Error; err != nil {
		return nil, err
	}
	return kewenangan, nil
}
