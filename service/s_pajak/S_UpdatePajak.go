package s_pajak

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePajakService struct {
	updatePajakRepo r_pajak.UpdatePajakRepository
}

func NewUpdatePajakService(updatePajakRepo r_pajak.UpdatePajakRepository) *UpdatePajakService {
	return &UpdatePajakService{updatePajakRepo}
}

func (s *UpdatePajakService) UpdatePajak(ctx *gin.Context, pajak m_pajak.Pajak) (*m_pajak.Pajak, error) {
	// Mendapatkan ID grup dari parameter rute
	idPajakStr := ctx.Param("pajak_id")

	// Konversi string ke int menggunakan strconv.Atoi
	idPajak, err := strconv.Atoi(idPajakStr)
	if err != nil {
		return nil, err
	}

	// Set ID grup ke dalam struct grup
	pajak.PajakID = idPajak

	// Memanggil repository untuk memperbarui grup
	updatedPajak, err := s.updatePajakRepo.UpdatePajak(&pajak)
	if err != nil {
		return nil, err
	}

	return updatedPajak, nil
}
