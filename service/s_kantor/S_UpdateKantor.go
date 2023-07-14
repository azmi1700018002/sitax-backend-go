package s_kantor

import (
	"sitax/model/m_kantor"
	"sitax/repository/r_kantor"

	"github.com/gin-gonic/gin"
)

type UpdateKantorService struct {
	updateKantorRepo r_kantor.UpdateKantorRepository
}

func NewUpdateKantorService(updateKantorRepo r_kantor.UpdateKantorRepository) *UpdateKantorService {
	return &UpdateKantorService{updateKantorRepo}
}

func (s *UpdateKantorService) UpdateKantor(ctx *gin.Context, kantor m_kantor.Kantor) (*m_kantor.Kantor, error) {
	// Mendapatkan ID grup dari parameter rute
	KdKantorStr := ctx.Param("kd_kantor")

	// Set ID grup ke dalam struct grup
	kantor.KdKantor = KdKantorStr

	// Memanggil repository untuk memperbarui grup
	updatedKantor, err := s.updateKantorRepo.UpdateKantor(&kantor)
	if err != nil {
		return nil, err
	}

	return updatedKantor, nil
}
