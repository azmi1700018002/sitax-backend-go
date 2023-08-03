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

func (s *UpdateKewenanganService) UpdateKewenangan(ctx *gin.Context, groupId, menuId string, kewenangan m_kewenangan.Kewenangan) (*m_kewenangan.Kewenangan, error) {
	// Set GroupID and MenuID received from the route parameters into the kewenangan struct
	kewenangan.GroupID = groupId
	kewenangan.MenuID = menuId

	// Call the repository to update the kewenangan
	updatedKewenangan, err := s.updateKewenanganRepo.UpdateKewenangan(&kewenangan)
	if err != nil {
		return nil, err
	}

	return updatedKewenangan, nil
}
