package s_menu

import (
	"sitax/model/m_menu"
	"sitax/repository/r_menu"

	"github.com/gin-gonic/gin"
)

type AddMenuService struct {
	addMenuRepo r_menu.AddMenuRepository
}

func NewAddMenuService(addMenuRepo r_menu.AddMenuRepository) *AddMenuService {
	return &AddMenuService{addMenuRepo}
}

func (s *AddMenuService) AddMenu(ctx *gin.Context, menu m_menu.Menu, menuID string) (*m_menu.Menu, error) {
	addMenu, err := s.addMenuRepo.AddMenu(&menu, menuID)
	if err != nil {
		return nil, err
	}

	return addMenu, nil
}
