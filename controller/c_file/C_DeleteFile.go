package c_file

import (
	"net/http"
	"sitax/service/s_file"

	"github.com/gin-gonic/gin"
)

type deleteFileController struct {
	deleteFileService s_file.DeleteFileService
}

func NewFileDeleteController(deleteFileService s_file.DeleteFileService) *deleteFileController {
	return &deleteFileController{deleteFileService}
}

func (c *deleteFileController) DeleteFile(ctx *gin.Context) {
	id := ctx.Param("file_id")
	err := c.deleteFileService.DeleteFileByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID" + err.Error()})
		return
	}
	err = c.deleteFileService.DeleteFileByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": "File berhasil dihapus"})
}
