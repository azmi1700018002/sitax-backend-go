package r_kewenangan

import (
	"errors"
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

// func (r *updateKewenanganRepository) UpdateKewenangan(kewenangan *m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error) {
// 	// Check if the user exists
// 	var existingKewenangan m_kewenangan.Kewenangan
// 	if err := db.Server().Where("group_id = ?", kewenangan.GroupID).First(&existingKewenangan).Error; err != nil {
// 		return nil, err
// 	}

// 	// Update user data in the database
// 	if err := db.Server().Model(&m_kewenangan.Kewenangan{}).
// 		Where("group_id = ?", kewenangan.GroupID).
// 		Updates(map[string]interface{}{
// 			"menu_id":    kewenangan.MenuID,
// 			"is_created": kewenangan.IsCreated,
// 			"is_updated": kewenangan.IsUpdated,
// 			"is_deleted": kewenangan.IsDeleted,
// 		}).Error; err != nil {
// 		return nil, err
// 	}
// 	return kewenangan, nil
// }

func (r *updateKewenanganRepository) UpdateKewenangan(kewenangan *m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error) {
	// Check if the user exists
	var existingKewenangan m_kewenangan.Kewenangan
	if err := db.Server().Where("group_id = ? AND menu_id = ?", kewenangan.GroupID, kewenangan.MenuID).First(&existingKewenangan).Error; err != nil {
		return nil, err
	}

	// Check if MenuID has changed
	if existingKewenangan.MenuID != kewenangan.MenuID {
		// Check if the new MenuID already exists in the database for the same GroupID
		var duplicateKewenangan m_kewenangan.Kewenangan
		if err := db.Server().Where("group_id = ? AND menu_id = ?", existingKewenangan.GroupID, kewenangan.MenuID).First(&duplicateKewenangan).Error; err == nil {
			// Combination already exists, handle the conflict (e.g., return an error or take appropriate action)
			return nil, errors.New("the new MenuID already exists for the same GroupID")
		}

		// Update existingKewenangan with the new MenuID
		existingKewenangan.MenuID = kewenangan.MenuID
	}

	// Update other attributes
	existingKewenangan.GroupID = kewenangan.GroupID
	existingKewenangan.IsCreated = kewenangan.IsCreated
	existingKewenangan.IsUpdated = kewenangan.IsUpdated
	existingKewenangan.IsDeleted = kewenangan.IsDeleted

	// Update user data in the database
	if err := db.Server().Model(&m_kewenangan.Kewenangan{}).
		Where("group_id = ? AND menu_id = ?", kewenangan.GroupID, kewenangan.MenuID).
		Updates(map[string]interface{}{
			"group_id":   existingKewenangan.GroupID,
			"menu_id":    existingKewenangan.MenuID,
			"is_created": existingKewenangan.IsCreated,
			"is_updated": existingKewenangan.IsUpdated,
			"is_deleted": existingKewenangan.IsDeleted,
		}).Error; err != nil {
		return nil, err
	}

	// If group_id and menu_id have changed, perform additional updates
	if existingKewenangan.GroupID != kewenangan.GroupID || existingKewenangan.MenuID != kewenangan.MenuID {
		// Perform additional updates using the new group_id and menu_id
		if err := db.Server().Model(&m_kewenangan.Kewenangan{}).
			Where("group_id = ? AND menu_id = ?", kewenangan.GroupID, kewenangan.MenuID).
			Updates(map[string]interface{}{
				"group_id":   kewenangan.GroupID,
				"menu_id":    kewenangan.MenuID,
				"is_created": existingKewenangan.IsCreated,
				"is_updated": existingKewenangan.IsUpdated,
				"is_deleted": existingKewenangan.IsDeleted,
			}).Error; err != nil {
			return nil, err
		}
	}

	return &existingKewenangan, nil
}
