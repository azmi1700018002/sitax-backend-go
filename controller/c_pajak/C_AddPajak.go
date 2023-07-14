package c_pajak

import (
	"net/http"
	"sitax/model/m_pajak"
	"sitax/service/s_pajak"

	"github.com/gin-gonic/gin"
)

type addPajakController struct {
	addPajakService *s_pajak.AddPajakService
}

func NewPajakAddController(addPajakService *s_pajak.AddPajakService) *addPajakController {
	return &addPajakController{addPajakService}
}

func (c *addPajakController) AddPajak(ctx *gin.Context) {
	var pajak m_pajak.Pajak
	if err := ctx.ShouldBind(&pajak); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	addPajak, err := c.addPajakService.AddPajak(ctx, pajak)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addPajak, "message": "Pajak Telah Berhasil Ditambahkan"})
}
