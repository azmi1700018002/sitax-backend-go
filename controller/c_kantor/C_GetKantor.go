package c_kantor

import (
	"net/http"
	"sitax/service/s_kantor"

	"github.com/gin-gonic/gin"
)

type KantorController struct {
	getKantorService s_kantor.GetKantorService
}

func NewGetKantorController(getKantorService s_kantor.GetKantorService) *KantorController {
	return &KantorController{getKantorService}
}

func (c *KantorController) GetAllKantor(ctx *gin.Context) {
	kantor, err := c.getKantorService.GetAllKantor()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data kantor", "data": kantor})
}
