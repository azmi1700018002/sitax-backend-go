package r_menu

import (
	"sitax/config/db"
	"sitax/model/m_menu"
)

type DeleteMenuRepository interface {
	DeleteMenuByID(id string) error
}

type menuDeleteRepository struct{}

func NewDeleteMenuRepository() DeleteMenuRepository {
	return &menuDeleteRepository{}
}

func (*menuDeleteRepository) DeleteMenuByID(id string) (err error) {
	// Menghapus data menu dari tabel Menu
	if err = db.Server().Unscoped().Where("menu_id = ?", id).Delete(&m_menu.Menu{}).Error; err != nil {
		return err
	}

	return nil
}
