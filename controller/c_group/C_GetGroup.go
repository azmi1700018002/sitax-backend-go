package c_group

import (
	"net/http"
	"sitax/service/s_group"

	"github.com/gin-gonic/gin"
)

type GroupController struct {
	getGroupService s_group.GetGroupService
}

func NewGetGroupController(getGroupService s_group.GetGroupService) *GroupController {
	return &GroupController{getGroupService}
}

func (c *GroupController) GetAllGroup(ctx *gin.Context) {
	group, err := c.getGroupService.GetAllGroup()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data group", "data": group})
}
