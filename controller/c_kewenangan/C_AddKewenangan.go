package c_kewenangan

import (
	"net/http"
	"sitax/model/m_kewenangan"
	"sitax/service/s_kewenangan"

	"github.com/gin-gonic/gin"
)

type addKewenanganController struct {
	addKewenanganService *s_kewenangan.AddKewenanganService
}

func NewKewenanganAddController(addKewenanganService *s_kewenangan.AddKewenanganService) *addKewenanganController {
	return &addKewenanganController{addKewenanganService}
}

func (c *addKewenanganController) AddKewenangan(ctx *gin.Context) {
	var kewenangan m_kewenangan.Kewenangan
	if err := ctx.ShouldBind(&kewenangan); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupIdStr := ctx.Param("group_id") // Mengambil KewenanganID dari permintaan HTTP
	addKewenangan, err := c.addKewenanganService.AddKewenangan(ctx, kewenangan, groupIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addKewenangan, "message": "Kewenangan Telah Berhasil Ditambahkan"})
}
