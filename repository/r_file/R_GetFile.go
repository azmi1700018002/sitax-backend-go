package r_file

import (
	"sitax/config/db"
	"sitax/model/m_file"
)

type GetFileRepository interface {
	GetAllFile() ([]m_file.File, error)
	GetFileByID(fileID string) (*m_file.File, error)
}

type getFileRepository struct{}

func NewGetFileRepository() GetFileRepository {
	return &getFileRepository{}
}

func (r *getFileRepository) GetAllFile() ([]m_file.File, error) {
	var files []m_file.File
	result := db.Server().Preload("FileIDfk.PajakIDfk").Preload("FileIDpanduanpajakfk").Find(&files)
	// result := db.Server().Find(&files)
	if result.Error != nil {
		return nil, result.Error
	}
	return files, nil
}

func (r *getFileRepository) GetFileByID(fileID string) (*m_file.File, error) {
	var file m_file.File
	result := db.Server().Where("file_id = ?", fileID).First(&file)
	if result.Error != nil {
		return nil, result.Error
	}
	return &file, nil
}
