package s_file

import (
	"mime/multipart"
	"sitax/model/m_file"
	"sitax/repository/r_file"

	"github.com/gin-gonic/gin"
)

type AddFileService struct {
	addFileRepo r_file.AddFileRepository
}

func NewAddFileService(addFileRepo r_file.AddFileRepository) *AddFileService {
	return &AddFileService{addFileRepo}
}

func (s *AddFileService) AddFile(ctx *gin.Context, file m_file.File, fileID string, dataFile *multipart.FileHeader) (*m_file.File, error) {

	addFile, err := s.addFileRepo.AddFile(&file, fileID, dataFile)
	if err != nil {
		return nil, err
	}

	return addFile, nil
}
