package c_group

import (
	"net/http"
	"sitax/service/s_group"

	"github.com/gin-gonic/gin"
)

type deleteGroupController struct {
	deleteGroupService s_group.DeleteGroupService
}

func NewGroupDeleteController(deleteGroupService s_group.DeleteGroupService) *deleteGroupController {
	return &deleteGroupController{deleteGroupService}
}

func (c *deleteGroupController) DeleteGroup(ctx *gin.Context) {
	id := ctx.Param("group_id")
	err := c.deleteGroupService.DeleteGroupByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteGroupService.DeleteGroupByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Group berhasil dihapus"})
}
