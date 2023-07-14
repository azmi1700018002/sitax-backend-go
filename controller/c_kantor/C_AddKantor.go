package c_kantor

import (
	"net/http"
	"sitax/model/m_kantor"
	"sitax/service/s_kantor"

	"github.com/gin-gonic/gin"
)

type addKantorController struct {
	addKantorService *s_kantor.AddKantorService
}

func NewKantorAddController(addKantorService *s_kantor.AddKantorService) *addKantorController {
	return &addKantorController{addKantorService}
}

func (c *addKantorController) AddKantor(ctx *gin.Context) {
	var kantor m_kantor.Kantor
	if err := ctx.ShouldBind(&kantor); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	KdKantor := ctx.Param("kd_kantor") // Mengambil KantorID dari permintaan HTTP
	addKantor, err := c.addKantorService.AddKantor(ctx, kantor, KdKantor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addKantor, "message": "Kantor Telah Berhasil Ditambahkan"})
}
