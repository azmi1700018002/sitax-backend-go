package c_kewenangan

import (
	"net/http"
	"sitax/service/s_kewenangan"

	"github.com/gin-gonic/gin"
)

type deleteKewenanganController struct {
	deleteKewenanganService s_kewenangan.DeleteKewenanganService
}

func NewKewenanganDeleteController(deleteKewenanganService s_kewenangan.DeleteKewenanganService) *deleteKewenanganController {
	return &deleteKewenanganController{deleteKewenanganService}
}

func (c *deleteKewenanganController) DeleteKewenangan(ctx *gin.Context) {
	GroupID := ctx.Param("group_id")
	MenuID := ctx.Param("menu_id")
	err := c.deleteKewenanganService.DeleteKewenanganByID(GroupID, MenuID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid GroupID"})
		return
	}
	err = c.deleteKewenanganService.DeleteKewenanganByID(GroupID, MenuID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Kewenangan berhasil dihapus"})
}
