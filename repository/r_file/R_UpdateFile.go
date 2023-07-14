package r_file

import (
	"sitax/config/db"
	"sitax/model/m_file"
)

type UpdateFileRepository interface {
	UpdateFile(file *m_file.File) (*m_file.File, error)
}

type updateFileRepository struct{}

func NewUpdateFileRepository() UpdateFileRepository {
	return &updateFileRepository{}
}

func (r *updateFileRepository) UpdateFile(file *m_file.File) (*m_file.File, error) {
	// Check if the user exists
	var existingFile m_file.File
	if err := db.Server().Where("file_id = ?", file.FileID).First(&existingFile).Error; err != nil {
		return nil, err
	}

	// Update user data in the database
	if err := db.Server().Model(&m_file.File{}).
		Where("file_id = ?", file.FileID).
		Updates(map[string]interface{}{
			"file_judul": file.FileJudul,
			"file_path":  file.FilePath,
			"file_date":  file.FileDate,
			"file_jenis": file.FileJenis,
		}).Error; err != nil {
		return nil, err
	}
	return file, nil
}
