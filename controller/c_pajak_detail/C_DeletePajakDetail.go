package c_pajak_detail

import (
	"net/http"
	"sitax/service/s_pajak_detail"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deletePajakDetailController struct {
	deletePajakDetailService s_pajak_detail.DeletePajakDetailService
}

func NewPajakDetailDeleteController(deletePajakDetailService s_pajak_detail.DeletePajakDetailService) *deletePajakDetailController {
	return &deletePajakDetailController{deletePajakDetailService}
}

func (c *deletePajakDetailController) DeletePajakDetail(ctx *gin.Context) {
	PajakDetailID := ctx.Param("pajak_detail_id")

	// Konversi PajakDetailID menjadi tipe data int
	idPajakDetail, err := strconv.Atoi(PajakDetailID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.deletePajakDetailService.DeletePajakDetailByID(idPajakDetail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "PajakDetail berhasil dihapus"})
}
