package s_panduan_pajak

import (
	"sitax/model/m_panduan_pajak"
	"sitax/repository/r_panduan_pajak"

	"github.com/gin-gonic/gin"
)

type AddPanduanPajakService struct {
	addPanduanPajakRepo r_panduan_pajak.AddPanduanPajakRepository
}

func NewAddPanduanPajakService(addPanduanPajakRepo r_panduan_pajak.AddPanduanPajakRepository) *AddPanduanPajakService {
	return &AddPanduanPajakService{addPanduanPajakRepo}
}

func (s *AddPanduanPajakService) AddPanduanPajak(ctx *gin.Context, panduan_pajak m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error) {
	addPanduanPajak, err := s.addPanduanPajakRepo.AddPanduanPajak(&panduan_pajak)
	if err != nil {
		return nil, err
	}

	return addPanduanPajak, nil
}
