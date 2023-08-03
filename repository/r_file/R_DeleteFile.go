package r_file

import (
	"errors"
	"os"
	"path/filepath"
	"sitax/config/db"
	"sitax/model/m_file"
	"sitax/model/m_pajak"

	"gorm.io/gorm"
)

type DeleteFileRepository interface {
	DeleteFileByID(id string) error
}

type fileDeleteRepository struct{}

func NewDeleteFileRepository() DeleteFileRepository {
	return &fileDeleteRepository{}
}

func (*fileDeleteRepository) DeleteFileByID(id string) (err error) {
	// Mengambil data file dari database
	file := &m_file.File{}
	if err = db.Server().First(file, "file_id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// Jika data tidak ditemukan, kembalikan nil tanpa error
			return nil
		}
		// Tangani error lainnya seperti biasa
		return err
	}
	// Membangun path absolut dari file gambar
	basePath := "../../../../../../../laragon/www/Sitax/file/"
	gambarPath := filepath.Join(basePath, file.FileJudul)

	// Hapus file gambar
	err = os.Remove(gambarPath)
	if err != nil {
		// Tangani kesalahan jika terjadi kesalahan dalam menghapus gambar
		return err
	}

	// Mengubah nilai file_id pada entitas pengguna yang memiliki FileID yang sesuai menjadi null
	if err = db.Server().Model(&m_pajak.Pajak{}).Where("file_id = ?", id).Update("file_id", nil).Error; err != nil {
		return err
	}

	// Menghapus data file dari tabel File
	if err = db.Server().Unscoped().Where("file_id = ?", id).Delete(&m_file.File{}).Error; err != nil {
		return err
	}

	return nil
}
