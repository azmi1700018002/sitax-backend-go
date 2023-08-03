package c_pajak

import (
	"net/http"
	"sitax/service/s_pajak"
	"strconv"

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

func (c *PajakController) GetPajakByID(ctx *gin.Context) {
	PajakIDStr := ctx.Param("pajak_id")
	PajakID, err := strconv.Atoi(PajakIDStr)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pajak, err := c.getPajakService.GetPajakByID(PajakID)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, pajak)
}
