package r_menu

import (
	"sitax/config/db"
	"sitax/model/m_menu"
)

type GetMenuRepository interface {
	GetAllMenu() ([]m_menu.Menu, error)
	GetMenuByID(menuID int) (*m_menu.Menu, error)
}

type getMenuRepository struct{}

func NewGetMenuRepository() GetMenuRepository {
	return &getMenuRepository{}
}

func (r *getMenuRepository) GetAllMenu() ([]m_menu.Menu, error) {
	var menus []m_menu.Menu
	result := db.Server().Preload("MenuIDfk.GroupIDfk").Find(&menus)
	// result := db.Server().Find(&menus)
	if result.Error != nil {
		return nil, result.Error
	}
	return menus, nil
}

func (r *getMenuRepository) GetMenuByID(menuID int) (*m_menu.Menu, error) {
	var menu m_menu.Menu
	result := db.Server().Where("menu_id = ?", menuID).First(&menu)
	if result.Error != nil {
		return nil, result.Error
	}
	return &menu, nil
}
