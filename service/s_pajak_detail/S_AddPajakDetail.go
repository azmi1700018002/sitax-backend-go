package s_pajak_detail

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak_detail"

	"github.com/gin-gonic/gin"
)

type AddPajakDetailService struct {
	addPajakDetailRepo r_pajak_detail.AddPajakDetailRepository
}

func NewAddPajakDetailService(addPajakDetailRepo r_pajak_detail.AddPajakDetailRepository) *AddPajakDetailService {
	return &AddPajakDetailService{addPajakDetailRepo}
}

func (s *AddPajakDetailService) AddPajakDetail(ctx *gin.Context, pajak m_pajak.PajakDetail) (*m_pajak.PajakDetail, error) {
	addPajakDetail, err := s.addPajakDetailRepo.AddPajakDetail(&pajak)
	if err != nil {
		return nil, err
	}

	return addPajakDetail, nil
}
