package c_user

import (
	"net/http"
	"path/filepath"
	"sitax/model/m_user"
	"sitax/repository/r_user"

	"github.com/gin-gonic/gin"
)

type registerUserController struct {
	registerUserRepo r_user.RegisterUserRepository
}

func NewUserController(registerUserRepo r_user.RegisterUserRepository) *registerUserController {
	return &registerUserController{
		registerUserRepo: registerUserRepo,
	}
}

func (c *registerUserController) RegisterUser(ctx *gin.Context) {
	var user m_user.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cek apakah file foto profil di-upload
	file, err := ctx.FormFile("profile_picture")
	if err == nil {
		// Simpan file di server local
		newFileName := filepath.Base(filepath.Clean(file.Filename))
		savePath := "../../../../../../../laragon/www/Sitax/assets/img/profile/" + newFileName

		if err := ctx.SaveUploadedFile(file, savePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		user.ProfilePicture = newFileName
	} else {
		// jika tidak, set avatar CDN sebagai profil gambar
		user.ProfilePicture = "default_profile.png"
	}

	if err := c.registerUserRepo.RegisterUser(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user})
}
