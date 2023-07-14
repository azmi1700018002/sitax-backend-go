package c_menu

import (
	"net/http"
	"sitax/model/m_menu"
	"sitax/service/s_menu"

	"github.com/gin-gonic/gin"
)

type UpdateMenuController struct {
	updateMenuService *s_menu.UpdateMenuService
}

func NewUpdateMenuController(updateMenuService *s_menu.UpdateMenuService) *UpdateMenuController {
	return &UpdateMenuController{updateMenuService}
}

func (c *UpdateMenuController) UpdateMenu(ctx *gin.Context) {
	// Mendapatkan data menu dari request body
	var menu m_menu.Menu
	err := ctx.ShouldBind(&menu)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update menu
	updatedMenu, err := c.updateMenuService.UpdateMenu(ctx, menu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedMenu, "message": "Menu berhasil diupdate"})
}
