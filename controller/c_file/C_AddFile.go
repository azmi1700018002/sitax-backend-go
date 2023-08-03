package c_file

import (
	"net/http"
	"sitax/model/m_file"
	"sitax/service/s_file"

	"github.com/gin-gonic/gin"
)

type addFileController struct {
	addFileService *s_file.AddFileService
}

func NewFileAddController(addFileService *s_file.AddFileService) *addFileController {
	return &addFileController{addFileService}
}

func (c *addFileController) AddFile(ctx *gin.Context) {
	var file m_file.File
	if err := ctx.ShouldBind(&file); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileID := ctx.Param("FileID") // Mengambil FileID dari permintaan HTTP

	// Use two variables to handle the return values of ctx.FormFile
	dataFile, err := ctx.FormFile("file_judul")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	addFile, err := c.addFileService.AddFile(ctx, file, fileID, dataFile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": addFile, "message": "File Telah Berhasil Ditambahkan"})
}
