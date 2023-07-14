package s_kewenangan

import (
	"sitax/model/m_kewenangan"
	"sitax/repository/r_kewenangan"

	"github.com/gin-gonic/gin"
)

type UpdateKewenanganService struct {
	updateKewenanganRepo r_kewenangan.UpdateKewenanganRepository
}

func NewUpdateKewenanganService(updateKewenanganRepo r_kewenangan.UpdateKewenanganRepository) *UpdateKewenanganService {
	return &UpdateKewenanganService{updateKewenanganRepo}
}

func (s *UpdateKewenanganService) UpdateKewenangan(ctx *gin.Context, kewenangan m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error) {
	// Mendapatkan ID grup dari parameter rute
	groupIdStr := ctx.Param("group_id")

	// Set ID grup ke dalam struct grup
	kewenangan.GroupID = groupIdStr

	// Memanggil repository untuk memperbarui grup
	updatedKewenangan, err := s.updateKewenanganRepo.UpdateKewenangan(&kewenangan)
	if err != nil {
		return nil, err
	}

	return updatedKewenangan, nil
}
