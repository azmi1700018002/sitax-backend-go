package c_panduan_pajak

import (
	"net/http"
	"sitax/service/s_panduan_pajak"

	"github.com/gin-gonic/gin"
)

type PanduanPajakController struct {
	getPanduanPajakService s_panduan_pajak.GetPanduanPajakService
}

func NewGetPanduanPajakController(getPanduanPajakService s_panduan_pajak.GetPanduanPajakService) *PanduanPajakController {
	return &PanduanPajakController{getPanduanPajakService}
}

func (c *PanduanPajakController) GetAllPanduanPajak(ctx *gin.Context) {
	panduan_pajak, err := c.getPanduanPajakService.GetAllPanduanPajak()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data panduan_pajak", "data": panduan_pajak})
}
