package r_file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"sitax/config/db"
	"sitax/model/m_file"
)

type AddFileRepository interface {
	AddFile(file *m_file.File, fileID string, dataFile *multipart.FileHeader) (*m_file.File, error)
}

type addFileRepository struct{}

func NewAddFileRepository() AddFileRepository {
	return &addFileRepository{}
}

func (r *addFileRepository) AddFile(file *m_file.File, fileID string, dataFile *multipart.FileHeader) (*m_file.File, error) {
	// Generate filename
	// fileExt := filepath.Ext(dataFile.Filename)
	newFileName := filepath.Base(filepath.Clean(dataFile.Filename))

	// Open the source file
	src, err := dataFile.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	// Determine the destination path
	basePath := "../../../../../../../laragon/www/SitaxUpdate/file/"
	savePath := filepath.Join(basePath, newFileName)
	// savePath := "./uploads/" + newFileName
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

	// Create a FileInfo object to store in the database
	fileInfo := &m_file.File{
		FileID:    file.FileID,
		FileJudul: newFileName,
		FilePath:  file.FilePath,
		FileDate:  file.FileDate,
		FileJenis: file.FileJenis,
		// Add other fields as needed
	}

	// Store the fileInfo in the database
	err = db.Server().Create(fileInfo).Error
	if err != nil {
		return nil, err
	}

	return fileInfo, nil
}
