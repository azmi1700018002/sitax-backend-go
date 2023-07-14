package r_file

import (
	"sitax/config/db"
	"sitax/model/m_file"
)

type AddFileRepository interface {
	AddFile(file *m_file.File, fileID string) (*m_file.File, error)
}

type addFileRepository struct{}

func NewAddFileRepository() AddFileRepository {
	return &addFileRepository{}
}

func (r *addFileRepository) AddFile(file *m_file.File, fileID string) (*m_file.File, error) {
	err := db.Server().Create(&file).Error
	if err != nil {
		return nil, err
	}
	return file, nil
}
