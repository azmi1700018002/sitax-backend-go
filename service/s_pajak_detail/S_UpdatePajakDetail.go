package s_pajak_detail

import (
	"sitax/model/m_pajak"
	"sitax/repository/r_pajak_detail"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePajakDetailService struct {
	updatePajakDetailRepo r_pajak_detail.UpdatePajakDetailRepository
}

func NewUpdatePajakDetailService(updatePajakDetailRepo r_pajak_detail.UpdatePajakDetailRepository) *UpdatePajakDetailService {
	return &UpdatePajakDetailService{updatePajakDetailRepo}
}

func (s *UpdatePajakDetailService) UpdatePajakDetail(ctx *gin.Context, pajak m_pajak.PajakDetail) (*m_pajak.PajakDetail, error) {
	// Mendapatkan ID grup dari parameter rute
	idPajakDetailStr := ctx.Param("pajak_detail_id")

	// Konversi string ke int menggunakan strconv.Atoi
	idPajakDetail, err := strconv.Atoi(idPajakDetailStr)
	if err != nil {
		return nil, err
	}

	// Set ID grup ke dalam struct grup
	pajak.PajakDetailID = idPajakDetail

	// Memanggil repository untuk memperbarui grup
	updatedPajakDetail, err := s.updatePajakDetailRepo.UpdatePajakDetail(&pajak)
	if err != nil {
		return nil, err
	}

	return updatedPajakDetail, nil
}
