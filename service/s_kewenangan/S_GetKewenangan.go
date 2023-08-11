package s_kewenangan

import (
	"sitax/model/m_kewenangan"
	"sitax/repository/r_kewenangan"
)

type GetKewenanganService interface {
	GetAllKewenangan() ([]m_kewenangan.Kewenangan, error)
	GetKewenanganByID(GroupID string) ([]*m_kewenangan.Kewenangan, error)
}

type getKewenanganService struct {
	getKewenanganRepository r_kewenangan.GetKewenanganRepository
}

func NewGetKewenanganService(getKewenanganRepository r_kewenangan.GetKewenanganRepository) GetKewenanganService {
	return &getKewenanganService{
		getKewenanganRepository: getKewenanganRepository,
	}
}

func (s *getKewenanganService) GetAllKewenangan() ([]m_kewenangan.Kewenangan, error) {
	return s.getKewenanganRepository.GetAllKewenangan()
}

func (s *getKewenanganService) GetKewenanganByID(GroupID string) ([]*m_kewenangan.Kewenangan, error) {
	return s.getKewenanganRepository.GetKewenanganByID(GroupID)
}
