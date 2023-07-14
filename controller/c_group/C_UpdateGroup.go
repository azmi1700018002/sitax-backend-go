package c_group

import (
	"net/http"
	"sitax/model/m_group"
	"sitax/service/s_group"

	"github.com/gin-gonic/gin"
)

type UpdateGroupController struct {
	updateGroupService *s_group.UpdateGroupService
}

func NewUpdateGroupController(updateGroupService *s_group.UpdateGroupService) *UpdateGroupController {
	return &UpdateGroupController{updateGroupService}
}

func (c *UpdateGroupController) UpdateGroup(ctx *gin.Context) {
	// Mendapatkan data group dari request body
	var group m_group.Group
	err := ctx.ShouldBind(&group)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update group
	updatedGroup, err := c.updateGroupService.UpdateGroup(ctx, group)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedGroup, "message": "Group berhasil diupdate"})
}
