package c_panduan_pajak

import (
	"net/http"
	"sitax/model/m_panduan_pajak"
	"sitax/service/s_panduan_pajak"

	"github.com/gin-gonic/gin"
)

type addPanduanPajakController struct {
	addPanduanPajakService *s_panduan_pajak.AddPanduanPajakService
}

func NewPanduanPajakAddController(addPanduanPajakService *s_panduan_pajak.AddPanduanPajakService) *addPanduanPajakController {
	return &addPanduanPajakController{addPanduanPajakService}
}

func (c *addPanduanPajakController) AddPanduanPajak(ctx *gin.Context) {
	var panduan_pajak m_panduan_pajak.PanduanPajak
	if err := ctx.ShouldBind(&panduan_pajak); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addPanduanPajak, err := c.addPanduanPajakService.AddPanduanPajak(ctx, panduan_pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addPanduanPajak, "message": "PanduanPajak Telah Berhasil Ditambahkan"})
}
