package s_file

import (
	"mime/multipart"
	"sitax/model/m_file"
	"sitax/repository/r_file"

	"github.com/gin-gonic/gin"
)

type UpdateFileService struct {
	updateFileRepo r_file.UpdateFileRepository
}

func NewUpdateFileService(updateFileRepo r_file.UpdateFileRepository) *UpdateFileService {
	return &UpdateFileService{updateFileRepo}
}

func (s *UpdateFileService) UpdateFile(ctx *gin.Context, file m_file.File, dataFile *multipart.FileHeader) (*m_file.File, error) {
	// Mendapatkan ID grup dari parameter rute
	idFileStr := ctx.Param("file_id")

	// Set ID grup ke dalam struct grup
	file.FileID = idFileStr

	// Memanggil repository untuk memperbarui grup
	updatedFile, err := s.updateFileRepo.UpdateFile(&file, dataFile)
	if err != nil {
		return nil, err
	}

	return updatedFile, nil
}
