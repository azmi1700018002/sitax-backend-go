package r_file

import (
	"sitax/config/db"
	"sitax/model/m_file"
)

type DeleteFileRepository interface {
	DeleteFileByID(id string) error
}

type fileDeleteRepository struct{}

func NewDeleteFileRepository() DeleteFileRepository {
	return &fileDeleteRepository{}
}

func (*fileDeleteRepository) DeleteFileByID(id string) (err error) {
	// Menghapus data file dari tabel File
	if err = db.Server().Unscoped().Where("file_id = ?", id).Delete(&m_file.File{}).Error; err != nil {
		return err
	}

	return nil
}
