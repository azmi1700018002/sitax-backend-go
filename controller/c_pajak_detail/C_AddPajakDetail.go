package c_pajak_detail

import (
	"net/http"
	"sitax/model/m_pajak"
	"sitax/service/s_pajak_detail"

	"github.com/gin-gonic/gin"
)

type addPajakDetailController struct {
	addPajakDetailService *s_pajak_detail.AddPajakDetailService
}

func NewPajakDetailAddController(addPajakDetailService *s_pajak_detail.AddPajakDetailService) *addPajakDetailController {
	return &addPajakDetailController{addPajakDetailService}
}

func (c *addPajakDetailController) AddPajakDetail(ctx *gin.Context) {
	var pajak m_pajak.PajakDetail
	if err := ctx.ShouldBind(&pajak); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addPajakDetail, err := c.addPajakDetailService.AddPajakDetail(ctx, pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addPajakDetail, "message": "Pajak Detail Telah Berhasil Ditambahkan"})
}
