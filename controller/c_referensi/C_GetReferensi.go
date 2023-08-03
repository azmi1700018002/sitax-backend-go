package c_referensi

import (
	"net/http"
	"sitax/service/s_referensi"

	"github.com/gin-gonic/gin"
)

type ReferensiController struct {
	getReferensiService s_referensi.GetReferensiService
}

func NewGetReferensiController(getReferensiService s_referensi.GetReferensiService) *ReferensiController {
	return &ReferensiController{getReferensiService}
}

func (c *ReferensiController) GetAllReferensi(ctx *gin.Context) {
	// Mendapatkan nilai parameter query dari URL
	grpID := ctx.Query("grp_id")
	ref := ctx.Query("ref")
	ket := ctx.Query("ket")

	referensi, err := c.getReferensiService.GetAllReferensi(grpID, ref, ket)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data referensi", "data": referensi})
}
