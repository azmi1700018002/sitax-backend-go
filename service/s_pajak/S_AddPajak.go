package s_pajak

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak"

	"github.com/gin-gonic/gin"
)

type AddPajakService struct {
	addPajakRepo r_pajak.AddPajakRepository
}

func NewAddPajakService(addPajakRepo r_pajak.AddPajakRepository) *AddPajakService {
	return &AddPajakService{addPajakRepo}
}

func (s *AddPajakService) AddPajak(ctx *gin.Context, pajak m_pajak.Pajak) (*m_pajak.Pajak, error) {
	addPajak, err := s.addPajakRepo.AddPajak(&pajak)
	if err != nil {
		return nil, err
	}

	return addPajak, nil
}
