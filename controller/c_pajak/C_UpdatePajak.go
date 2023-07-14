package c_pajak

import (
	"net/http"
	"sitax/model/m_pajak"
	"sitax/service/s_pajak"

	"github.com/gin-gonic/gin"
)

type UpdatePajakController struct {
	updatePajakService *s_pajak.UpdatePajakService
}

func NewUpdatePajakController(updatePajakService *s_pajak.UpdatePajakService) *UpdatePajakController {
	return &UpdatePajakController{updatePajakService}
}

func (c *UpdatePajakController) UpdatePajak(ctx *gin.Context) {
	// Mendapatkan data pajak dari request body
	var pajak m_pajak.Pajak
	err := ctx.ShouldBind(&pajak)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update pajak
	updatedPajak, err := c.updatePajakService.UpdatePajak(ctx, pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedPajak, "message": "Pajak berhasil diupdate"})
}
