package c_kewenangan

import (
	"net/http"
	"sitax/model/m_kewenangan"
	"sitax/service/s_kewenangan"

	"github.com/gin-gonic/gin"
)

type UpdateKewenanganController struct {
	updateKewenanganService *s_kewenangan.UpdateKewenanganService
}

func NewUpdateKewenanganController(updateKewenanganService *s_kewenangan.UpdateKewenanganService) *UpdateKewenanganController {
	return &UpdateKewenanganController{updateKewenanganService}
}

func (c *UpdateKewenanganController) UpdateKewenangan(ctx *gin.Context) {
	// Get the group_id and menu_id from the route parameters
	groupID := ctx.Param("group_id")
	menuID := ctx.Param("menu_id")

	// Get the kewenangan data from the request body
	var kewenangan m_kewenangan.Kewenangan
	err := ctx.ShouldBind(&kewenangan)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service to update the kewenangan
	updatedKewenangan, err := c.updateKewenanganService.UpdateKewenangan(ctx, groupID, menuID, kewenangan)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Display the JSON response
	ctx.JSON(http.StatusOK, gin.H{"data": updatedKewenangan, "message": "Kewenangan berhasil diupdate"})
}
