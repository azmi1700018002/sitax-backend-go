package c_menu

import (
	"net/http"
	"sitax/service/s_menu"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	getMenuService s_menu.GetMenuService
}

func NewGetMenuController(getMenuService s_menu.GetMenuService) *MenuController {
	return &MenuController{getMenuService}
}

func (c *MenuController) GetAllMenu(ctx *gin.Context) {
	menu, err := c.getMenuService.GetAllMenu()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data menu", "data": menu})
}
