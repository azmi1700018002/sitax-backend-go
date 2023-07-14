package s_pajak_detail

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak_detail"
)

type GetPajakDetailService interface {
	GetAllPajakDetail() ([]m_pajak.PajakDetail, error)
	GetPajakDetailByID(PajakDetailID int) (*m_pajak.PajakDetail, error)
}

type getPajakDetailService struct {
	getPajakDetailRepository r_pajak_detail.GetPajakDetailRepository
}

func NewGetPajakDetailService(getPajakDetailRepository r_pajak_detail.GetPajakDetailRepository) GetPajakDetailService {
	return &getPajakDetailService{
		getPajakDetailRepository: getPajakDetailRepository,
	}
}

func (s *getPajakDetailService) GetAllPajakDetail() ([]m_pajak.PajakDetail, error) {
	return s.getPajakDetailRepository.GetAllPajakDetail()
}

func (s *getPajakDetailService) GetPajakDetailByID(PajakDetailID int) (*m_pajak.PajakDetail, error) {
	return s.getPajakDetailRepository.GetPajakDetailByID(PajakDetailID)
}
