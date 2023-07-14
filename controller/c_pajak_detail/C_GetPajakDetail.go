package c_pajak_detail

import (
	"net/http"
	"sitax/service/s_pajak_detail"

	"github.com/gin-gonic/gin"
)

type PajakDetailController struct {
	getPajakDetailService s_pajak_detail.GetPajakDetailService
}

func NewGetPajakDetailController(getPajakDetailService s_pajak_detail.GetPajakDetailService) *PajakDetailController {
	return &PajakDetailController{getPajakDetailService}
}

func (c *PajakDetailController) GetAllPajakDetail(ctx *gin.Context) {
	pajak, err := c.getPajakDetailService.GetAllPajakDetail()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data pajak", "data": pajak})
}
