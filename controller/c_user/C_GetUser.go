package c_user

import (
	"net/http"
	"sitax/service/s_user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	getUserService s_user.GetUserService
}

func NewGetUserController(getUserService s_user.GetUserService) *UserController {
	return &UserController{getUserService}
}

func (c *UserController) GetAllUser(ctx *gin.Context) {
	users, err := c.getUserService.GetAllUser()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data user", "data": users})
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := c.getUserService.GetUserByID(username)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, user)
}
