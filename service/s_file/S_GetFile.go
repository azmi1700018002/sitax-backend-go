package s_file

import (
	"sitax/model/m_file"
	"sitax/repository/r_file"
)

type GetFileService interface {
	GetAllFile() ([]m_file.File, error)
	GetFileByID(fileID string) (*m_file.File, error)
}

type getFileService struct {
	getFileRepository r_file.GetFileRepository
}

func NewGetFileService(getFileRepository r_file.GetFileRepository) GetFileService {
	return &getFileService{
		getFileRepository: getFileRepository,
	}
}

func (s *getFileService) GetAllFile() ([]m_file.File, error) {
	return s.getFileRepository.GetAllFile()
}

func (s *getFileService) GetFileByID(fileID string) (*m_file.File, error) {
	return s.getFileRepository.GetFileByID(fileID)
}
