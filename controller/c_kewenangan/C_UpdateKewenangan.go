package c_kewenangan

import (
	"net/http"
	"sitax/model/m_kewenangan"
	"sitax/service/s_kewenangan"

	"github.com/gin-gonic/gin"
)

type UpdateKewenanganController struct {
	updateKewenanganService *s_kewenangan.UpdateKewenanganService
}

func NewUpdateKewenanganController(updateKewenanganService *s_kewenangan.UpdateKewenanganService) *UpdateKewenanganController {
	return &UpdateKewenanganController{updateKewenanganService}
}

func (c *UpdateKewenanganController) UpdateKewenangan(ctx *gin.Context) {
	// Mendapatkan data kewenangan dari request body
	var kewenangan m_kewenangan.Kewenangan
	err := ctx.ShouldBind(&kewenangan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update kewenangan
	updatedKewenangan, err := c.updateKewenanganService.UpdateKewenangan(ctx, kewenangan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedKewenangan, "message": "Kewenangan berhasil diupdate"})
}
