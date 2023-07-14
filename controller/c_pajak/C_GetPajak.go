package c_pajak

import (
	"net/http"
	"sitax/service/s_pajak"

	"github.com/gin-gonic/gin"
)

type PajakController struct {
	getPajakService s_pajak.GetPajakService
}

func NewGetPajakController(getPajakService s_pajak.GetPajakService) *PajakController {
	return &PajakController{getPajakService}
}

func (c *PajakController) GetAllPajak(ctx *gin.Context) {
	pajak, err := c.getPajakService.GetAllPajak()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data pajak", "data": pajak})
}
