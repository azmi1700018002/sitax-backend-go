package s_menu

import (
	"sitax/model/m_menu"
	"sitax/repository/r_menu"

	"github.com/gin-gonic/gin"
)

type UpdateMenuService struct {
	updateMenuRepo r_menu.UpdateMenuRepository
}

func NewUpdateMenuService(updateMenuRepo r_menu.UpdateMenuRepository) *UpdateMenuService {
	return &UpdateMenuService{updateMenuRepo}
}

func (s *UpdateMenuService) UpdateMenu(ctx *gin.Context, menu m_menu.Menu) (*m_menu.Menu, error) {
	// Mendapatkan ID grup dari parameter rute
	idMenuStr := ctx.Param("menu_id")

	// Set ID grup ke dalam struct grup
	menu.MenuID = idMenuStr

	// Memanggil repository untuk memperbarui grup
	updatedMenu, err := s.updateMenuRepo.UpdateMenu(&menu)
	if err != nil {
		return nil, err
	}

	return updatedMenu, nil
}
