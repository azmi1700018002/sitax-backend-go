package c_kantor

import (
	"net/http"
	"sitax/model/m_kantor"
	"sitax/service/s_kantor"

	"github.com/gin-gonic/gin"
)

type UpdateKantorController struct {
	updateKantorService *s_kantor.UpdateKantorService
}

func NewUpdateKantorController(updateKantorService *s_kantor.UpdateKantorService) *UpdateKantorController {
	return &UpdateKantorController{updateKantorService}
}

func (c *UpdateKantorController) UpdateKantor(ctx *gin.Context) {
	// Mendapatkan data kantor dari request body
	var kantor m_kantor.Kantor
	err := ctx.ShouldBind(&kantor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update kantor
	updatedKantor, err := c.updateKantorService.UpdateKantor(ctx, kantor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedKantor, "message": "Kantor berhasil diupdate"})
}
