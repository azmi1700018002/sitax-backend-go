package r_menu

import (
	"sitax/config/db"
	"sitax/model/m_menu"
)

type AddMenuRepository interface {
	AddMenu(menu *m_menu.Menu, menuID string) (*m_menu.Menu, error)
}

type addMenuRepository struct{}

func NewAddMenuRepository() AddMenuRepository {
	return &addMenuRepository{}
}

func (r *addMenuRepository) AddMenu(menu *m_menu.Menu, menuID string) (*m_menu.Menu, error) {
	err := db.Server().Create(&menu).Error
	if err != nil {
		return nil, err
	}
	return menu, nil
}
