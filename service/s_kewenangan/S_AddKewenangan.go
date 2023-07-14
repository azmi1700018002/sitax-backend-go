package s_kewenangan

import (
	"sitax/model/m_kewenangan"
	"sitax/repository/r_kewenangan"

	"github.com/gin-gonic/gin"
)

type AddKewenanganService struct {
	addKewenanganRepo r_kewenangan.AddKewenanganRepository
}

func NewAddKewenanganService(addKewenanganRepo r_kewenangan.AddKewenanganRepository) *AddKewenanganService {
	return &AddKewenanganService{addKewenanganRepo}
}

func (s *AddKewenanganService) AddKewenangan(ctx *gin.Context, kewenangan m_kewenangan.Kewenangan, GroupID string) (*m_kewenangan.Kewenangan, error) {
	addKewenangan, err := s.addKewenanganRepo.AddKewenangan(&kewenangan, GroupID)
	if err != nil {
		return nil, err
	}

	return addKewenangan, nil
}
