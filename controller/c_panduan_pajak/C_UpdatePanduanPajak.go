package c_panduan_pajak

import (
	"net/http"
	"sitax/model/m_panduan_pajak"
	"sitax/service/s_panduan_pajak"

	"github.com/gin-gonic/gin"
)

type UpdatePanduanPajakController struct {
	updatePanduanPajakService *s_panduan_pajak.UpdatePanduanPajakService
}

func NewUpdatePanduanPajakController(updatePanduanPajakService *s_panduan_pajak.UpdatePanduanPajakService) *UpdatePanduanPajakController {
	return &UpdatePanduanPajakController{updatePanduanPajakService}
}

func (c *UpdatePanduanPajakController) UpdatePanduanPajak(ctx *gin.Context) {
	// Mendapatkan data panduan_pajak dari request body
	var panduan_pajak m_panduan_pajak.PanduanPajak
	err := ctx.ShouldBind(&panduan_pajak)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update panduan_pajak
	updatedPanduanPajak, err := c.updatePanduanPajakService.UpdatePanduanPajak(ctx, panduan_pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedPanduanPajak, "message": "PanduanPajak berhasil diupdate"})
}
