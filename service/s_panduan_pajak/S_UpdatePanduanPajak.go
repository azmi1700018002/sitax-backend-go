package s_panduan_pajak

import (
	"sitax/model/m_panduan_pajak"
	"sitax/repository/r_panduan_pajak"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePanduanPajakService struct {
	updatePanduanPajakRepo r_panduan_pajak.UpdatePanduanPajakRepository
}

func NewUpdatePanduanPajakService(updatePanduanPajakRepo r_panduan_pajak.UpdatePanduanPajakRepository) *UpdatePanduanPajakService {
	return &UpdatePanduanPajakService{updatePanduanPajakRepo}
}

func (s *UpdatePanduanPajakService) UpdatePanduanPajak(ctx *gin.Context, panduan_pajak m_panduan_pajak.PanduanPajak) (*m_panduan_pajak.PanduanPajak, error) {
	// Mendapatkan ID grup dari parameter rute
	idPanduanPajakStr := ctx.Param("panduan_pajak_id")

	// Konversi string ke int menggunakan strconv.Atoi
	idPanduanPajak, err := strconv.Atoi(idPanduanPajakStr)
	if err != nil {
		return nil, err
	}

	// Set ID grup ke dalam struct grup
	panduan_pajak.PanduanPajakID = idPanduanPajak

	// Memanggil repository untuk memperbarui grup
	updatedPanduanPajak, err := s.updatePanduanPajakRepo.UpdatePanduanPajak(&panduan_pajak)
	if err != nil {
		return nil, err
	}

	return updatedPanduanPajak, nil
}
