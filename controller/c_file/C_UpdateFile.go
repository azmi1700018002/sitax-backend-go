package c_file

import (
	"net/http"
	"sitax/model/m_file"
	"sitax/service/s_file"

	"github.com/gin-gonic/gin"
)

type UpdateFileController struct {
	updateFileService *s_file.UpdateFileService
}

func NewUpdateFileController(updateFileService *s_file.UpdateFileService) *UpdateFileController {
	return &UpdateFileController{updateFileService}
}

func (c *UpdateFileController) UpdateFile(ctx *gin.Context) {
	// Mendapatkan data file dari request body
	var file m_file.File
	err := ctx.ShouldBind(&file)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memanggil service untuk update file
	updatedFile, err := c.updateFileService.UpdateFile(ctx, file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Menampilkan respon JSON
	ctx.JSON(http.StatusOK, gin.H{"data": updatedFile, "message": "File berhasil diupdate"})
}
