package s_pajak

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak"
)

type GetPajakService interface {
	GetAllPajak() ([]m_pajak.Pajak, error)
	GetPajakByID(PajakID int) (*m_pajak.Pajak, error)
}

type getPajakService struct {
	getPajakRepository r_pajak.GetPajakRepository
}

func NewGetPajakService(getPajakRepository r_pajak.GetPajakRepository) GetPajakService {
	return &getPajakService{
		getPajakRepository: getPajakRepository,
	}
}

func (s *getPajakService) GetAllPajak() ([]m_pajak.Pajak, error) {
	return s.getPajakRepository.GetAllPajak()
}

func (s *getPajakService) GetPajakByID(PajakID int) (*m_pajak.Pajak, error) {
	return s.getPajakRepository.GetPajakByID(PajakID)
}
