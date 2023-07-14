package c_menu

import (
	"net/http"
	"sitax/model/m_menu"
	"sitax/service/s_menu"

	"github.com/gin-gonic/gin"
)

type addMenuController struct {
	addMenuService *s_menu.AddMenuService
}

func NewMenuAddController(addMenuService *s_menu.AddMenuService) *addMenuController {
	return &addMenuController{addMenuService}
}

func (c *addMenuController) AddMenu(ctx *gin.Context) {
	var menu m_menu.Menu
	if err := ctx.ShouldBind(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuID := ctx.Param("MenuID") // Mengambil MenuID dari permintaan HTTP
	addMenu, err := c.addMenuService.AddMenu(ctx, menu, menuID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addMenu, "message": "Menu Telah Berhasil Ditambahkan"})
}
