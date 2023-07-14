package c_group

import (
	"net/http"
	"sitax/model/m_group"
	"sitax/service/s_group"

	"github.com/gin-gonic/gin"
)

type addGroupController struct {
	addGroupService *s_group.AddGroupService
}

func NewGroupAddController(addGroupService *s_group.AddGroupService) *addGroupController {
	return &addGroupController{addGroupService}
}

func (c *addGroupController) AddGroup(ctx *gin.Context) {
	var group m_group.Group
	if err := ctx.ShouldBind(&group); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupID := ctx.Param("GroupID") // Mengambil GroupID dari permintaan HTTP
	addGroup, err := c.addGroupService.AddGroup(ctx, group, groupID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addGroup, "message": "Group Telah Berhasil Ditambahkan"})
}
