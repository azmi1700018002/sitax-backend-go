package s_kantor

import (
	"sitax/model/m_kantor"
	"sitax/repository/r_kantor"

	"github.com/gin-gonic/gin"
)

type AddKantorService struct {
	addKantorRepo r_kantor.AddKantorRepository
}

func NewAddKantorService(addKantorRepo r_kantor.AddKantorRepository) *AddKantorService {
	return &AddKantorService{addKantorRepo}
}

func (s *AddKantorService) AddKantor(ctx *gin.Context, kantor m_kantor.Kantor, KdKantor string) (*m_kantor.Kantor, error) {
	addKantor, err := s.addKantorRepo.AddKantor(&kantor, KdKantor)
	if err != nil {
		return nil, err
	}

	return addKantor, nil
}
