package c_pajak_detail

import (
	"net/http"
	"sitax/model/m_pajak"
	"sitax/service/s_pajak_detail"

	"github.com/gin-gonic/gin"
)

type UpdatePajakDetailController struct {
	updatePajakDetailService *s_pajak_detail.UpdatePajakDetailService
}

func NewUpdatePajakDetailController(updatePajakDetailService *s_pajak_detail.UpdatePajakDetailService) *UpdatePajakDetailController {
	return &UpdatePajakDetailController{updatePajakDetailService}
}

func (c *UpdatePajakDetailController) UpdatePajakDetail(ctx *gin.Context) {
	// Mendapatkan data pajak dari request body
	var pajak m_pajak.PajakDetail
	err := ctx.ShouldBind(&pajak)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update pajak
	updatedPajakDetail, err := c.updatePajakDetailService.UpdatePajakDetail(ctx, pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedPajakDetail, "message": "PajakDetail berhasil diupdate"})
}
