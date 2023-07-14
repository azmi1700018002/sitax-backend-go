package s_menu

import (
	"sitax/model/m_menu"
	"sitax/repository/r_menu"
)

type GetMenuService interface {
	GetAllMenu() ([]m_menu.Menu, error)
	GetMenuByID(menuID int) (*m_menu.Menu, error)
}

type getMenuService struct {
	getMenuRepository r_menu.GetMenuRepository
}

func NewGetMenuService(getMenuRepository r_menu.GetMenuRepository) GetMenuService {
	return &getMenuService{
		getMenuRepository: getMenuRepository,
	}
}

func (s *getMenuService) GetAllMenu() ([]m_menu.Menu, error) {
	return s.getMenuRepository.GetAllMenu()
}

func (s *getMenuService) GetMenuByID(menuID int) (*m_menu.Menu, error) {
	return s.getMenuRepository.GetMenuByID(menuID)
}
