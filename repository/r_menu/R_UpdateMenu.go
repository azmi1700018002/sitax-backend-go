package r_menu

import (
	"sitax/config/db"
	"sitax/model/m_menu"
)

type UpdateMenuRepository interface {
	UpdateMenu(menu *m_menu.Menu) (*m_menu.Menu, error)
}

type updateMenuRepository struct{}

func NewUpdateMenuRepository() UpdateMenuRepository {
	return &updateMenuRepository{}
}

func (r *updateMenuRepository) UpdateMenu(menu *m_menu.Menu) (*m_menu.Menu, error) {
	// Check if the user exists
	var existingMenu m_menu.Menu
	if err := db.Server().Where("menu_id = ?", menu.MenuID).First(&existingMenu).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_menu.Menu{}).
		Where("menu_id = ?", menu.MenuID).
		Updates(map[string]interface{}{
			"menu_nama":      menu.MenuNama,
			"menu_link":      menu.MenuLink,
			"menu_deskripsi": menu.MenuDeskripsi,
			"menu_status":    menu.MenuStatus,
			"menu_icon":      menu.MenuIcon,
			"menu_kategori":  menu.MenuKategori,
			"parent_id":      menu.ParentID,
			"parent_sts":     menu.ParentSts,
			"no_urut":        menu.NoUrut,
		}).Error; err != nil {
		return nil, err
	}
	return menu, nil
}
