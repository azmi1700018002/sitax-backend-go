package c_panduan_pajak

import (
	"net/http"
	"sitax/service/s_panduan_pajak"
	"strconv"

	"github.com/gin-gonic/gin"
)

type deletePanduanPajakController struct {
	deletePanduanPajakService s_panduan_pajak.DeletePanduanPajakService
}

func NewPanduanPajakDeleteController(deletePanduanPajakService s_panduan_pajak.DeletePanduanPajakService) *deletePanduanPajakController {
	return &deletePanduanPajakController{deletePanduanPajakService}
}

func (c *deletePanduanPajakController) DeletePanduanPajak(ctx *gin.Context) {
	PanduanPajakID := ctx.Param("panduan_pajak_id")

	// Konversi PanduanPajakID menjadi tipe data int
	idPanduanPajak, err := strconv.Atoi(PanduanPajakID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = c.deletePanduanPajakService.DeletePanduanPajakByID(idPanduanPajak)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": "PanduanPajak berhasil dihapus"})
}
