package r_file

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sitax/config/db"
	"sitax/model/m_file"
)

type UpdateFileRepository interface {
	UpdateFile(file *m_file.File, dataFile *multipart.FileHeader) (*m_file.File, error)
}

type updateFileRepository struct{}

func NewUpdateFileRepository() UpdateFileRepository {
	return &updateFileRepository{}
}

func (r *updateFileRepository) UpdateFile(file *m_file.File, dataFile *multipart.FileHeader) (*m_file.File, error) {
	// Check if the file exists
	var existingFile m_file.File
	if err := db.Server().Where("file_id = ?", file.FileID).First(&existingFile).Error; err != nil {
		return nil, err
	}

	// Path for update profile
	basePath := "../../../../../../../laragon/www/SitaxUpdate/file/"

	fmt.Println("Deleting old image file:", filepath.Join(basePath, existingFile.FileJudul))

	// If there is a new image file uploaded, delete the old file from local storage
	if dataFile != nil {
		if existingFile.FileJudul != "" {
			err := os.Remove(filepath.Join(basePath, existingFile.FileJudul))
			if err != nil {
				return nil, err
			}
		}

		newFileName := filepath.Base(filepath.Clean(dataFile.Filename))

		// Open the source file
		src, err := dataFile.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		// Create the destination file
		savePath := filepath.Join(basePath, newFileName)

		dst, err := os.Create(savePath)
		if err != nil {
			return nil, err
		}
		defer dst.Close()

		// Copy the contents from source to destination
		_, err = io.Copy(dst, src)
		if err != nil {
			return nil, err
		}

		// Update the product's image path
		file.FileJudul = newFileName
	} else {
		// If no new image file is uploaded, retain the previous file name
		file.FileJudul = existingFile.FileJudul
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
