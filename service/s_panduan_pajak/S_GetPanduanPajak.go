package s_panduan_pajak

import (
	"sitax/model/m_panduan_pajak"
	"sitax/repository/r_panduan_pajak"
)

type GetPanduanPajakService interface {
	GetAllPanduanPajak() ([]m_panduan_pajak.PanduanPajak, error)
	GetPanduanPajakByID(PanduanPajakID int) (*m_panduan_pajak.PanduanPajak, error)
}

type getPanduanPajakService struct {
	getPanduanPajakRepository r_panduan_pajak.GetPanduanPajakRepository
}

func NewGetPanduanPajakService(getPanduanPajakRepository r_panduan_pajak.GetPanduanPajakRepository) GetPanduanPajakService {
	return &getPanduanPajakService{
		getPanduanPajakRepository: getPanduanPajakRepository,
	}
}

func (s *getPanduanPajakService) GetAllPanduanPajak() ([]m_panduan_pajak.PanduanPajak, error) {
	return s.getPanduanPajakRepository.GetAllPanduanPajak()
}

func (s *getPanduanPajakService) GetPanduanPajakByID(PanduanPajakID int) (*m_panduan_pajak.PanduanPajak, error) {
	return s.getPanduanPajakRepository.GetPanduanPajakByID(PanduanPajakID)
}
