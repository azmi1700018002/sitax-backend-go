package c_kantor

import (
	"net/http"
	"sitax/service/s_kantor"

	"github.com/gin-gonic/gin"
)

type deleteKantorController struct {
	deleteKantorService s_kantor.DeleteKantorService
}

func NewKantorDeleteController(deleteKantorService s_kantor.DeleteKantorService) *deleteKantorController {
	return &deleteKantorController{deleteKantorService}
}

func (c *deleteKantorController) DeleteKantor(ctx *gin.Context) {
	KdKantor := ctx.Param("kd_kantor")
	err := c.deleteKantorService.DeleteKantorByID(KdKantor)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid KdKantor" + err.Error()})
		return
	}
	err = c.deleteKantorService.DeleteKantorByID(KdKantor)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Kantor berhasil dihapus"})
}
