package c_kewenangan

import (
	"net/http"
	"sitax/service/s_kewenangan"

	"github.com/gin-gonic/gin"
)

type KewenanganController struct {
	getKewenanganService s_kewenangan.GetKewenanganService
}

func NewGetKewenanganController(getKewenanganService s_kewenangan.GetKewenanganService) *KewenanganController {
	return &KewenanganController{getKewenanganService}
}

func (c *KewenanganController) GetAllKewenangan(ctx *gin.Context) {
	kewenangan, err := c.getKewenanganService.GetAllKewenangan()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data kewenangan", "data": kewenangan})
}

func (c *KewenanganController) GetKewenanganByID(ctx *gin.Context) {
	groupID := ctx.Param("group_id")
	kewenangans, err := c.getKewenanganService.GetKewenanganByID(groupID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, kewenangans)
}
