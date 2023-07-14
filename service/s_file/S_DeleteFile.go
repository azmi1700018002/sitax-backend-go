package s_file

import "sitax/repository/r_file"

type DeleteFileService interface {
	DeleteFileByID(id string) error
}

type deleteFileService struct {
	deleteFileRepo r_file.DeleteFileRepository
}

func NewDeleteFileService(deleteFileRepo r_file.DeleteFileRepository) DeleteFileService {
	return &deleteFileService{deleteFileRepo}
}

func (s *deleteFileService) DeleteFileByID(id string) error {
	err := s.deleteFileRepo.DeleteFileByID(id)
	if err != nil {
		return err
	}
	return nil
}
