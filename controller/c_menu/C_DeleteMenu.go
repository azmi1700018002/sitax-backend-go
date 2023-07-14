package c_menu

import (
	"net/http"
	"sitax/service/s_menu"

	"github.com/gin-gonic/gin"
)

type deleteMenuController struct {
	deleteMenuService s_menu.DeleteMenuService
}

func NewMenuDeleteController(deleteMenuService s_menu.DeleteMenuService) *deleteMenuController {
	return &deleteMenuController{deleteMenuService}
}

func (c *deleteMenuController) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("menu_id")
	err := c.deleteMenuService.DeleteMenuByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = c.deleteMenuService.DeleteMenuByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "Menu berhasil dihapus"})
}
