package c_file

import (
	"net/http"
	"sitax/service/s_file"

	"github.com/gin-gonic/gin"
)

type FileController struct {
	getFileService s_file.GetFileService
}

func NewGetFileController(getFileService s_file.GetFileService) *FileController {
	return &FileController{getFileService}
}

func (c *FileController) GetAllFile(ctx *gin.Context) {
	file, err := c.getFileService.GetAllFile()
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Hasil data file", "data": file})
}
