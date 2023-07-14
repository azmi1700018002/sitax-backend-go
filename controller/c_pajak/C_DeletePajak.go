package c_pajak

import (
	"net/http"
	"sitax/service/s_pajak"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deletePajakController struct {
	deletePajakService s_pajak.DeletePajakService
}

func NewPajakDeleteController(deletePajakService s_pajak.DeletePajakService) *deletePajakController {
	return &deletePajakController{deletePajakService}
}

func (c *deletePajakController) DeletePajak(ctx *gin.Context) {
	PajakID := ctx.Param("pajak_id")

	// Konversi PajakID menjadi tipe data int
	idPajak, err := strconv.Atoi(PajakID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.deletePajakService.DeletePajakByID(idPajak)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "Pajak berhasil dihapus"})
}
