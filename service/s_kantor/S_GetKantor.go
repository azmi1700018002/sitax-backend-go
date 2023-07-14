package s_kantor

import (
	"sitax/model/m_kantor"
	"sitax/repository/r_kantor"
)

type GetKantorService interface {
	GetAllKantor() ([]m_kantor.Kantor, error)
	GetKantorByID(KdKantor string) (*m_kantor.Kantor, error)
}

type getKantorService struct {
	getKantorRepository r_kantor.GetKantorRepository
}

func NewGetKantorService(getKantorRepository r_kantor.GetKantorRepository) GetKantorService {
	return &getKantorService{
		getKantorRepository: getKantorRepository,
	}
}

func (s *getKantorService) GetAllKantor() ([]m_kantor.Kantor, error) {
	return s.getKantorRepository.GetAllKantor()
}

func (s *getKantorService) GetKantorByID(KdKantor string) (*m_kantor.Kantor, error) {
	return s.getKantorRepository.GetKantorByID(KdKantor)
}
