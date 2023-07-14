package c_user

import (
	"net/http"
	"sitax/service/s_user"

	"github.com/gin-gonic/gin"
)

type deleteUserController struct {
	deleteUserService s_user.DeleteUserService
}

func NewUserDeleteController(deleteUserService s_user.DeleteUserService) *deleteUserController {
	return &deleteUserController{deleteUserService}
}

func (c *deleteUserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	err := c.deleteUserService.DeleteUserByID(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Username"})
		return
	}
	err = c.deleteUserService.DeleteUserByID(username)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "User berhasil dihapus"})
}
